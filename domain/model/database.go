package model

import (
	"time"
)

// Articles represents the model for an Article
type Articles struct {
	ID          uint      `json:"ID"`
	SourceID    uint      `json:"sourceId"`
	Tags        string    `json:"tags"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	PubDate     time.Time `json:"pubdate"`
	Video       string    `json:"video"`
	VideoHeight uint16    `json:"videoHeight"`
	VideoWidth  uint16    `json:"videoWidth"`
	Thumbnail   string    `json:"thumbnail"`
	Description string    `json:"description"`
	AuthorName  string    `json:"authorName"`
	AuthorImage string    `json:"authorImage"`
}

type DiscordQueue struct {
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`
	ArticleId string    `json:"articleId"`
}

type DiscordWebHooks struct {
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`

	Name    string `json:"name"`
	Key     string `json:"key"`
	Url     string `json:"url"`
	Server  string `json:"server"`
	Channel string `json:"channel"`
	Enabled bool   `json:"enabled"`
}

type Icons struct {
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`

	FileName string `json:"fileName"`
	Site     string `json:"site"`
}

type Settings struct {
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`

	Key     string `json:"key"`
	Value   string `json:"value"`
	Options string `json:"options"`
}

type Sources struct {
	ID      uint   `json:"ID"`
	Site    string `json:"site"`
	Name    string `json:"name"`
	Source  string `json:"source"`
	Type    string `json:"type"`
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
	Url     string `json:"url"`
	Tags    string `json:"tags"`
}

type SourceLinks struct {
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`

	SourceID    uint   `json:"sourceId"`
	SourceType  string `json:"sourceType"`
	SourceName  string `json:"sourceName"`
	DiscordID   uint   `json:"discordId"`
	DiscordName string `json:"discordName"`
}
