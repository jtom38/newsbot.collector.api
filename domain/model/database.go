package model

import (
	"time"
)

// Articles represents the model for an Article
type Articles struct {
	ID int64 `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`

	SourceId int32 `json:"sourceId"`
	Tags string	`json:"tags"`
	Title string `json:"title"`
	Url string `json:"url"`
	PubDate time.Time `json:"pubdate"`
	Video string `json:"video"`
	VideoHeight int16 `json:"videoHeight"`
	VideoWidth int16 `json:"videoWidth"`
	Thumbnail string `json:"thumbnail"`
	Description string `json:"description"`
	AuthorName string `json:"authorName"`
	AuthorImage string `json:"authorImage"`
}

type DiscordQueue struct {
	ID int64 `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`
	ArticleId string `json:"articleId"`
}

type DiscordWebHooks struct {
	ID int32 `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`

	Name string 	`json:"name"`
	Key string 		`json:"key"`
	Url string 		`json:"url"`
	Server string 	`json:"server"`
	Channel string	`json:"channel"`
	Enabled bool	`json:"enabled"`
}

type Icons struct {
	ID int32 `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`

	FileName string	`json:"fileName"`
	Site string		`json:"site"`
}

type Settings struct {
	ID int16 `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`

	Key string		`json:"key"`
	Value string	`json:"value"`
	Options string	`json:"options"`
}

type Sources struct {
	ID int32 `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`

	Site string		`json:"site"`
	Name string		`json:"name"`
	Source string	`json:"source"`
	Type string		`json:"type"`
	Value string	`json:"value"`
	Enabled bool	`json:"enabled"`
	Url string		`json:"url"`
	Tags string		`json:"tags"`
}

type SourceLinks struct {
	ID int32 `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`
	
	SourceID string		`json:"sourceId"`
	SourceType string	`json:"sourceType"`
	SourceName string	`json:"sourceName"`
	DiscordID string	`json:"discordId"`
	DiscordName string	`json:"discordName"`
}
