package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	//"time"

	"github.com/jtom38/newsbot/collector/database"
)

type discordField struct {
	Name   *string `json:"name,omitempty"`
	Value  *string `json:"value,omitempty"`
	Inline *bool   `json:"inline,omitempty"`
}

type discordFooter struct {
	Value   *string `json:"text,omitempty"`
	IconUrl *string `json:"icon_url,omitempty"`
}

type discordAuthor struct {
	Name    *string `json:"name,omitempty"`
	Url     *string `json:"url,omitempty"`
	IconUrl *string `json:"icon_url,omitempty"`
}

type discordImage struct {
	Url *string `json:"url,omitempty"`
}

type DiscordEmbed struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Url         *string `json:"url,omitempty"`
	Color       *int32  `json:"color,omitempty"`
	//Timestamp   time.Time      `json:"timestamp,omitempty"`
	Fields    []*discordField `json:"fields,omitempty"`
	Author    discordAuthor   `json:"author,omitempty"`
	Image     discordImage    `json:"image,omitempty"`
	Thumbnail discordImage    `json:"thumbnail,omitempty"`
	Footer    *discordFooter  `json:"footer,omitempty"`
}

// Root object for Discord Webhook messages
type DiscordMessage struct {
	Username *string         `json:"username,omitempty"`
	Content  *string         `json:"content,omitempty"`
	Embeds   *[]DiscordEmbed `json:"embeds,omitempty"`
}

type Discord struct {
	Subscriptions []string
	article       database.Article
	Message       *DiscordMessage
}

func NewDiscordWebHookMessage(Article database.Article) Discord {
	return Discord{
		article: Article,
	}
}

// Generates the link field to expose in the message
func (dwh Discord) getFields() []*discordField {
	var fields []*discordField

	key := "Link"
	linkField := discordField{
		Name:  &key,
		Value: &dwh.article.Url,
	}

	fields = append(fields, &linkField)

	return fields
}

// This will create the message that will be sent out.
func (dwh Discord) GeneratePayload() (*DiscordMessage, error) {

	// Create the embed
	footerMessage := "Brought to you by Newsbot"
	footerUrl := ""
	description := dwh.convertFromHtml(dwh.article.Description)

	embed := DiscordEmbed{
		Title:       &dwh.article.Title,
		Description: &description,
		Thumbnail: discordImage{
			Url: &dwh.article.Thumbnail,
		},
		Fields: dwh.getFields(),
		Footer: &discordFooter{
			Value:   &footerMessage,
			IconUrl: &footerUrl,
		},
	}

	// attach the embed to an array
	var embedArray []DiscordEmbed
	embedArray = append(embedArray, embed)

	// create the base message
	msg := DiscordMessage{
		Embeds: &embedArray,
	}

	return &msg, nil
}

func (dwh Discord) SendPayload(Message *DiscordMessage, Url string) error {
	// Convert the message to a io.reader object
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(Message)

	// Send the message
	resp, err := http.Post(Url, "application/json", buffer)
	if err != nil {
		return err
	}

	// Check for 204
	if resp.StatusCode != 204 {
		defer resp.Body.Close()

		errMsg, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf(string(errMsg))
	}

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
