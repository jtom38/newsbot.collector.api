package models

import (
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/jtom38/newsbot/collector/database"
)

type ArticleDto struct {
	ID          uuid.UUID `json:"id"`
	Source      uuid.UUID `json:"sourceid"`
	Tags        []string  `json:"tags"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Pubdate     time.Time `json:"pubdate"`
	Video       string    `json:"video"`
	Videoheight int32     `json:"videoHeight"`
	Videowidth  int32     `json:"videoWidth"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	Authorname  string    `json:"authorName"`
	Authorimage string    `json:"authorImage"`
}

type ArticleDetailsDto struct {
	ID          uuid.UUID `json:"id"`
	Source      SourceDto `json:"source"`
	Tags        []string  `json:"tags"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Pubdate     time.Time `json:"pubdate"`
	Video       string    `json:"video"`
	Videoheight int32     `json:"videoHeight"`
	Videowidth  int32     `json:"videoWidth"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	Authorname  string    `json:"authorName"`
	Authorimage string    `json:"authorImage"`
}

type DiscordWebHooksDto struct {
	ID      uuid.UUID `json:"ID"`
	Url     string    `json:"url"`
	Server  string    `json:"server"`
	Channel string    `json:"channel"`
	Enabled bool      `json:"enabled"`
}

func ConvertToDiscordWebhookDto(i database.Discordwebhook) DiscordWebHooksDto {
	return DiscordWebHooksDto{
		ID:      i.ID,
		Url:     i.Url,
		Server:  i.Server,
		Channel: i.Channel,
		Enabled: i.Enabled,
	}
}

type SourceDto struct {
	ID      uuid.UUID `json:"id"`
	Site    string    `json:"site"`
	Name    string    `json:"name"`
	Source  string    `json:"source"`
	Type    string    `json:"type"`
	Value   string    `json:"value"`
	Enabled bool      `json:"enabled"`
	Url     string    `json:"url"`
	Tags    []string  `json:"tags"`
	Deleted bool      `json:"deleted"`
}

func ConvertToSourceDto(i database.Source) SourceDto {
	var deleted bool
	if !i.Deleted.Valid {
		deleted = true
	}

	return SourceDto{
		ID:      i.ID,
		Site:    i.Site,
		Name:    i.Name,
		Source:  i.Source,
		Type:    i.Type,
		Value:   i.Value.String,
		Enabled: i.Enabled,
		Url:     i.Url,
		Tags:    splitTags(i.Tags),
		Deleted: deleted,
	}
}

type DiscordQueueDto struct {
	ID        uuid.UUID `json:"id"`
	Articleid uuid.UUID `json:"articleId"`
}

type DiscordQueueDetailsDto struct {
	ID      uuid.UUID         `json:"id"`
	Article ArticleDetailsDto `json:"article"`
}

type SubscriptionDto struct {
	ID               uuid.UUID `json:"id"`
	DiscordWebhookId uuid.UUID `json:"discordwebhookid"`
	SourceId         uuid.UUID `json:"sourceid"`
}

func ConvertToSubscriptionDto(i database.Subscription) SubscriptionDto {
	c := SubscriptionDto{
		ID:               i.ID,
		DiscordWebhookId: i.Discordwebhookid,
		SourceId:         i.Sourceid,
	}
	return c
}

type SubscriptionDetailsDto struct {
	ID             uuid.UUID          `json:"id"`
	Source         SourceDto          `json:"source"`
	DiscordWebHook DiscordWebHooksDto `json:"discordwebhook"`
}

func splitTags(t string) []string {
	items := strings.Split(t, ", ")
	return items
}
