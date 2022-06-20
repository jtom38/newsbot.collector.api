// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ID          uuid.UUID
	Sourceid    uuid.UUID
	Tags        string
	Title       string
	Url         string
	Pubdate     time.Time
	Video       sql.NullString
	Videoheight int32
	Videowidth  int32
	Thumbnail   string
	Description string
	Authorname  sql.NullString
	Authorimage sql.NullString
}

type Discordqueue struct {
	ID        uuid.UUID
	Articleid uuid.UUID
}

type Discordwebhook struct {
	ID      uuid.UUID
	Url     string
	Server  string
	Channel string
	Enabled bool
}

type Icon struct {
	ID       uuid.UUID
	Filename string
	Site     string
}

type Setting struct {
	ID      uuid.UUID
	Key     string
	Value   string
	Options sql.NullString
}

type Source struct {
	ID      uuid.UUID
	Site    string
	Name    string
	Source  string
	Type    string
	Value   sql.NullString
	Enabled bool
	Url     string
	Tags    string
}

type Subscription struct {
	ID               uuid.UUID
	Discordwebhookid uuid.UUID
	Sourceid         uuid.UUID
}
