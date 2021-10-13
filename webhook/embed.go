package webhook

// Embed represents discord embed. https://discord.com/developers/docs/resources/channel#embed-object
type Embed struct {
	// Title of embed
	Title string `json:"title"`
	// Type of embed (always "rich" for webhook embeds)
	Type string `json:"type"`
	// Description of embed
	Description string `json:"description"`
	// URL of embed
	URL string `json:"url"`
	// Timestamp of embed content. Should be in ISO8601 format.
	Timestamp string `json:"timestamp"`
	// Color code of the embed
	Color     int             `json:"color"`
	Footer    *EmbedFooter    `json:"footer"`
	Image     *EmbedImage     `json:"image"`
	Thumbnail *EmbedThumbnail `json:"thumbnail"`
	Video     *EmbedVideo     `json:"video"`
	Provider  *EmbedProvider  `json:"provider"`
	Author    *EmbedAuthor    `json:"author"`
	Fields    []*EmbedField   `json:"fields"`
}

// EmbedFooter represents embed's footer
type EmbedFooter struct {
	// Footer text. Required.
	Text string `json:"text"`
	// URL of footer icon (only supports http(s) and attachments)
	IconURL string `json:"icon_url"`
	// Proxied URL of footer icon
	ProxyIconURL string `json:"proxy_icon_url"`
}

// EmbedImage represents Image that would be inside of the embed
type EmbedImage struct {
	// URL source of image (only supports http(s) and attachments)
	URL string `json:"url"`
	// A proxied url of the image
	ProxyURL string `json:"proxy_url"`
	// Height of image
	Height int `json:"height"`
	// Width of image
	Width int `json:"width"`
}

type EmbedThumbnail struct {
	// URL source of thumbnail (only supports http(s) and attachments)
	URL string `json:"url"`
	// A proxied url of the thumbnail
	ProxyURL string `json:"proxy_url"`
	// Height of thumbnail
	Height int `json:"height"`
	// Width of thumbnail
	Width int `json:"width"`
}

// EmbedVideo represents Video that would be inside of the embed
type EmbedVideo struct {
	// URL source of video
	URL string `json:"url"`
	// Height of video
	Height int `json:"height"`
	// Width of video
	Width int `json:"width"`
}

// EmbedAuthor represents author of the embed
type EmbedAuthor struct {
	// Name of author
	Name string `json:"name"`
	// URL of author
	URL string `json:"url"`
	// URL of author icon (only supports http(s) and attachments)
	IconURL string `json:"icon_url"`
	// Proxied URL of author icon
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedProvider struct {
	// Name of provider
	Name string `json:"name"`
	// URL of provider
	URL string `json:"url"`
}

// EmbedField represents embed field
type EmbedField struct {
	// Name of the field. Required.
	Name string `json:"name"`
	// Value of the field. Required.
	Value string `json:"value"`
	// Whether or not this field should display inline.
	Inline bool `json:"inline"`
}