package webhook

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// modified from
// https://github.com/itsTurnip/dishooks

// from https://github.com/kyb3r/dhooks/blob/cdd5f3f3bc109cbbc06f16a1cd9d39ed9d75a94e/dhooks/client.py#L118
// discordHostRegex matches "discord.com", "discordapp.com", "canary.discord.com", "ptb.discord.com"
var (
	discordHostRegex  = regexp.MustCompile("^((canary|ptb)\\.)?discord(?:app)?\\.com$")
	discordIDRegex    = regexp.MustCompile("^[0-9]{0,20}$")
	discordTokenRegex = regexp.MustCompile(`^[A-Za-z0-9\.\-\_]+$`)
)

// ErrParseWebhook represents error when parsing url, id, or token
var ErrParseWebhook = errors.New("failed to parse discord webhook url")

// Webhook is a representation of discord webhooks
type Webhook struct {
	ID        string `json:"id"`
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Token     string `json:"token"`
}

// Message represents a message to send using webhook.
// When sending a message there must be one of content or embeds.
type Message struct {
	// Content is a message body up to 2000 characters.
	Content   string   `json:"content"`
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	TTS       bool     `json:"tts"`
	Embeds    []*Embed `json:"embeds"`

	// Two more fields named `file` and `payload_json` are not presented here,
	// they are used for sending files and actually not used for regular messages.
	// See `SendFile`.
}

// Update is used to update webhooks using tokens. Only name and avatar can be updated.
type Update struct {
	Name string `json:"name"`
	// Avatar is a base64 encoded image. https://discord.com/developers/docs/reference#image-data
	Avatar string `json:"avatar"`
}

// FromURL parses URL and returns Webhook struct.
func FromURL(webhookURL string) (webhook *Webhook, err error) {
	urls, err := url.Parse(webhookURL)
	if err != nil {
		return
	}
	ok := discordHostRegex.MatchString(urls.Hostname())
	path := strings.Split(urls.Path, "/")
	// len is 5 normally, could be 6 with trailing /
	if !ok || len(path) < 5 || !strings.HasPrefix(urls.Path, "/api/webhook") {
		err = ErrParseWebhook
		return
	}
	id, token := path[3], path[4]
	return FromIDAndToken(id, token)
}

// FromIDAndToken  parses id and token and returns Webhook struct.
func FromIDAndToken(id string, token string) (webhook *Webhook, err error) {
	if !(discordIDRegex.MatchString(id) || discordTokenRegex.MatchString(token)) {
		err = ErrParseWebhook
		return
	}
	webhook = &Webhook{
		ID:    id,
		Token: token,
	}
	return
}

// URL returns discord url of the webhook
func (w *Webhook) URL() string {
	return "https://discord.com/api/webhooks/" + w.ID + "/" + w.Token
}

// Get gets information about webhook.
// Useful to get name, guild and channel IDs etc.
func (w *Webhook) Get() error {
	resp, err := http.Get(w.URL())
	if err != nil {
		return err
	}
	content, err := checkResponse(resp)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, w)
}

// SendMessage sends message to webhook
func (w *Webhook) SendMessage(message *Message) error {
	content, err := json.Marshal(message)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(content)
	resp, err := http.Post(w.URL(), "application/json", buf)
	if err != nil {
		return err
	}
	return checkError(resp)
}

// SendFile sends file with message to webhook
func (w *Webhook) SendFile(file []byte, filename string, message *Message) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	f, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}
	_, err = f.Write(file)
	if err != nil {
		return err
	}
	payload, err := writer.CreateFormField("payload_json")
	if err != nil {
		return err
	}
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	encode := make([]byte, base64.URLEncoding.EncodedLen(len(msg)))
	base64.URLEncoding.Encode(encode, msg)
	_, err = payload.Write(msg)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	resp, err := http.Post(w.URL(), writer.FormDataContentType(), body)
	if err != nil {
		return err
	}
	return checkError(resp)
}

// Modify modifies webhook at the client and discord side.
func (w *Webhook) Modify(update *Update) error {
	body, err := json.Marshal(update)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPatch, w.URL(), buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	response, err := checkResponse(resp)
	if err != nil {
		return err
	}
	return json.Unmarshal(response, w)
}

// Delete deletes webhook from discord.
func (w *Webhook) Delete() error {
	req, err := http.NewRequest(http.MethodDelete, w.URL(), nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return checkError(resp)
}
