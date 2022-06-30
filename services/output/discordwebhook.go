package output

import (
	"strings"
	"time"

	"github.com/jtom38/newsbot/collector/database"
)

type discordField struct {
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Inline bool `json:"inline,omitempty"`
}

type discordAuthor struct {
	Name string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type discordImage struct {
	Url string `json:"url,omitempty"`
}

type discordEmbed struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Url string `json:"url,omitempty"`
	Color int32 `json:"color,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Fields []discordField `json:"fields,omitempty"`
	Author discordAuthor `json:"author,omitempty"`
	Image discordImage `json:"image,omitempty"`
	Thumbnail discordImage `json:"thumbnail,omitempty"`
}

// Root object for Discord Webhook messages
type discordMessage struct {
	Content string `json:"content,omitempty"`
	Embeds []discordEmbed `json:"embeds,omitempty"`
}

type Discord struct {
	Subscriptions []string
	article database.Article
	Message discordMessage
}

func NewDiscordWebHookMessage(Subscriptions []string, Article database.Article) Discord {
	return Discord{
		Subscriptions: Subscriptions,
		article: Article,
		Message: discordMessage{
			Embeds: []discordEmbed{},
		},
	}
}

func (dwh Discord) GeneratePayload() error {
	// Convert the message 
	embed := discordEmbed {
		Title: dwh.article.Title,
		Description: dwh.convertFromHtml(dwh.article.Description),
		Url: dwh.article.Url,
		Thumbnail: discordImage{
			Url: dwh.article.Thumbnail,
		},
	}
	var arr []discordEmbed

	arr = append(arr, embed)
	dwh.Message.Embeds = arr
	return nil
}

func (dwh Discord) SendPayload() error {
	return nil
}

func (dwh Discord) convertFromHtml(body string) string {
	clean := body
	clean = strings.ReplaceAll(clean, "<h2>", "**")
	clean = strings.ReplaceAll(clean, "</h2>", "**")
	clean = strings.ReplaceAll(clean, "<h3>", "**")
	clean = strings.ReplaceAll(clean, "</h3>", "**\r\n")
	clean = strings.ReplaceAll(clean, "<strong>", "**")
	clean = strings.ReplaceAll(clean, "</strong>", "**\r\n")
	clean = strings.ReplaceAll(clean, "<ul>", "\r\n")
	clean = strings.ReplaceAll(clean, "</ul>", "")
	clean = strings.ReplaceAll(clean, "</li>", "\r\n")
	clean = strings.ReplaceAll(clean, "<li>", "> ")
	clean = strings.ReplaceAll(clean, "&#8220;", "\"")
	clean = strings.ReplaceAll(clean, "&#8221;", "\"")
	clean = strings.ReplaceAll(clean, "&#8230;", "...")
	clean = strings.ReplaceAll(clean, "<b>", "**")
	clean = strings.ReplaceAll(clean, "</b>", "**")
	clean = strings.ReplaceAll(clean, "<br>", "\r\n")
	clean = strings.ReplaceAll(clean, "<br/>", "\r\n")
	clean = strings.ReplaceAll(clean, "\xe2\x96\xa0", "*")
	clean = strings.ReplaceAll(clean, "\xa0", "\r\n")
	clean = strings.ReplaceAll(clean, "<p>", "")
	clean = strings.ReplaceAll(clean, "</p>", "\r\n")
	return clean
}

func (dwh Discord) convertLinks(body string) string {
	//items := regexp.MustCompile("<a(.*?)a>")
	return ""
} 
