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

// WebhookType is a webhook type for discord webhooks
// https://discordapp.com/developers/docs/resources/webhook#webhook-object-webhook-types
type WebhookType int

const (
	// Incoming type Webhooks can post messages to channels with a generated token
	Incoming WebhookType = iota + 1
	// ChannelFollower type Webhooks are internal webhooks used with Channel Following
	// to post new messages into channels
	ChannelFollower
)

const discordHost string = "discordapp.com"
const discordPathPattern string = "/api/webhooks/[0-9]{0,20}/[a-zA-Z0-9_-]+"

// ErrParseWebhook represents error from `WebhookFromURL`
var ErrParseWebhook = errors.New("Provided URL is not a Discord webhook URL")

// Webhook is a representation of discord webhooks
type Webhook struct {
	// ID of the webhook
	ID string `json:"id"`
	// Type is WebhookType
	Type      WebhookType `json:"type"`
	GuildID   string      `json:"guild_id"`
	ChannelID string      `json:"channel_id"`
	Name      string      `json:"name"`
	Avatar    string      `json:"avatar"`
	Token     string      `json:"token"`

	// There may be a one more field named `user` but it is not used in case
	// of posting messages using webhook. It is used for bots and clients only.
	// You would probably consider using discordgo or disgord for this purposes.
	// github.com/bwmarrin/discordgo
	// github.com/andersfylling/disgord
}

// WebhookMessage represents a message to send using webhook.
// When sending a message there must be one of content or embeds.
type WebhookMessage struct {
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

// WebhookUpdate is used to update webhooks using tokens.
type WebhookUpdate struct {
	Name string `json:"name"`
	// Avatar is a base64 encoded image. https://discordapp.com/developers/docs/reference#image-data
	Avatar string `json:"avatar"`

	// ChannelID field is not presented here as this is not allowed to update
	// channel ID using token.
}

// WebhookFromURL parses URL and returns Webhook struct.
func WebhookFromURL(webhookURL string) (webhook *Webhook, err error) {
	urls, err := url.Parse(webhookURL)
	if err != nil {
		return
	}
	if urls.Hostname() != discordHost {
		err = ErrParseWebhook
		return
	}
	ok, err := regexp.MatchString(discordPathPattern, urls.Path)
	if err != nil {
		return
	}
	if !ok {
		err = ErrParseWebhook
		return
	}
	path := strings.Split(urls.Path, "/")
	id, token := path[3], path[4]
	webhook = &Webhook{
		ID:    id,
		Token: token,
		Type:  Incoming,
	}
	return
}

// URL returns discord url of the webhook
func (w *Webhook) URL() string {
	return "https://" + discordHost + "/api/webhooks/" + w.ID + "/" + w.Token
}

// Get gets information about webhook.
// Useful to get name, guild and channel IDs etc.
func (w *Webhook) Get() error {
	resp, err := http.Get(w.URL())
	if err != nil {
		return err
	}
	content, err := CheckResponse(resp)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, w)
	if err != nil {
		return err
	}
	return nil
}

// SendMessage sends message to webhook
func (w *Webhook) SendMessage(message *WebhookMessage) (response []byte, err error) {
	content, err := json.Marshal(message)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(content)
	resp, err := http.Post(w.URL(), "application/json", buf)
	if err != nil {
		return
	}
	response, err = CheckResponse(resp)
	if err != nil {
		return
	}
	return
}

// SendFile sends file with message to webhook
func (w *Webhook) SendFile(file []byte, filename string, message *WebhookMessage) (response []byte, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	f, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return
	}
	_, err = f.Write(file)
	if err != nil {
		return
	}
	payload, err := writer.CreateFormField("payload_json")
	if err != nil {
		return
	}
	msg, err := json.Marshal(message)
	if err != nil {
		return
	}
	encode := make([]byte, base64.URLEncoding.EncodedLen(len(msg)))
	base64.URLEncoding.Encode(encode, msg)
	_, err = payload.Write(msg)
	if err != nil {
		return
	}
	err = writer.Close()
	if err != nil {
		return
	}
	resp, err := http.Post(w.URL(), writer.FormDataContentType(), body)
	if err != nil {
		return
	}
	response, err = CheckResponse(resp)
	if err != nil {
		return
	}
	return
}

// Modify modifies webhook at the client and discord side.
func (w *Webhook) Modify(update *WebhookUpdate) (response []byte, err error) {
	body, err := json.Marshal(update)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPatch, w.URL(), buf)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	response, err = CheckResponse(resp)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, w)
	if err != nil {
		return
	}
	return
}

// Delete deletes webhook from discord.
func (w *Webhook) Delete() (response []byte, err error) {
	req, err := http.NewRequest(http.MethodDelete, w.URL(), nil)
	if err != nil {
		return
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	response, err = CheckResponse(resp)
	if err != nil {
		return
	}
	return
}
