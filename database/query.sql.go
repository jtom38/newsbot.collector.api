// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createArticle = `-- name: CreateArticle :exec
INSERT INTO Articles 
(ID, SourceId, Tags, Title, Url, PubDate, Video, VideoHeight, VideoWidth, Thumbnail, Description, AuthorName, AuthorImage)
Values
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
`

type CreateArticleParams struct {
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

func (q *Queries) CreateArticle(ctx context.Context, arg CreateArticleParams) error {
	_, err := q.db.ExecContext(ctx, createArticle,
		arg.ID,
		arg.Sourceid,
		arg.Tags,
		arg.Title,
		arg.Url,
		arg.Pubdate,
		arg.Video,
		arg.Videoheight,
		arg.Videowidth,
		arg.Thumbnail,
		arg.Description,
		arg.Authorname,
		arg.Authorimage,
	)
	return err
}

const createDiscordQueue = `-- name: CreateDiscordQueue :exec
Insert into DiscordQueue
(ID, ArticleId)
Values
($1, $2)
`

type CreateDiscordQueueParams struct {
	ID        uuid.UUID
	Articleid uuid.UUID
}

// DiscordQueue
func (q *Queries) CreateDiscordQueue(ctx context.Context, arg CreateDiscordQueueParams) error {
	_, err := q.db.ExecContext(ctx, createDiscordQueue, arg.ID, arg.Articleid)
	return err
}

const createDiscordWebHook = `-- name: CreateDiscordWebHook :exec
Insert Into DiscordWebHooks
(ID, Name, Key, Url, Server, Channel, Enabled)
Values
($1, $2, $3, $4, $5, $6, $7)
`

type CreateDiscordWebHookParams struct {
	ID      uuid.UUID
	Name    sql.NullString
	Key     sql.NullString
	Url     sql.NullString
	Server  sql.NullString
	Channel sql.NullString
	Enabled sql.NullBool
}

// DiscordWebHooks
func (q *Queries) CreateDiscordWebHook(ctx context.Context, arg CreateDiscordWebHookParams) error {
	_, err := q.db.ExecContext(ctx, createDiscordWebHook,
		arg.ID,
		arg.Name,
		arg.Key,
		arg.Url,
		arg.Server,
		arg.Channel,
		arg.Enabled,
	)
	return err
}

const createIcon = `-- name: CreateIcon :exec

INSERT INTO Icons
(ID, FileName, Site)
VALUES
($1,$2,$3)
`

type CreateIconParams struct {
	ID       uuid.UUID
	Filename sql.NullString
	Site     sql.NullString
}

// Icons
func (q *Queries) CreateIcon(ctx context.Context, arg CreateIconParams) error {
	_, err := q.db.ExecContext(ctx, createIcon, arg.ID, arg.Filename, arg.Site)
	return err
}

const createSettings = `-- name: CreateSettings :one

Insert Into settings
(ID, Key, Value, OPTIONS)
Values
($1,$2,$3,$4)
RETURNING id, key, value, options
`

type CreateSettingsParams struct {
	ID      uuid.UUID
	Key     string
	Value   string
	Options sql.NullString
}

// Settings
func (q *Queries) CreateSettings(ctx context.Context, arg CreateSettingsParams) (Setting, error) {
	row := q.db.QueryRowContext(ctx, createSettings,
		arg.ID,
		arg.Key,
		arg.Value,
		arg.Options,
	)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Value,
		&i.Options,
	)
	return i, err
}

const createSource = `-- name: CreateSource :exec

Insert Into Sources
(ID, Site, Name, Source, Type, Value, Enabled, Url, Tags)
Values
($1,$2,$3,$4,$5,$6,$7,$8,$9)
`

type CreateSourceParams struct {
	ID      uuid.UUID
	Site    sql.NullString
	Name    sql.NullString
	Source  sql.NullString
	Type    sql.NullString
	Value   sql.NullString
	Enabled sql.NullBool
	Url     sql.NullString
	Tags    sql.NullString
}

// Sources
func (q *Queries) CreateSource(ctx context.Context, arg CreateSourceParams) error {
	_, err := q.db.ExecContext(ctx, createSource,
		arg.ID,
		arg.Site,
		arg.Name,
		arg.Source,
		arg.Type,
		arg.Value,
		arg.Enabled,
		arg.Url,
		arg.Tags,
	)
	return err
}

const deleteDiscordQueueItem = `-- name: DeleteDiscordQueueItem :exec
Delete From DiscordQueue Where ID = $1
`

func (q *Queries) DeleteDiscordQueueItem(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDiscordQueueItem, id)
	return err
}

const deleteDiscordWebHooks = `-- name: DeleteDiscordWebHooks :exec
Delete From discordwebhooks Where ID = $1
`

func (q *Queries) DeleteDiscordWebHooks(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteDiscordWebHooks, id)
	return err
}

const deleteIcon = `-- name: DeleteIcon :exec
Delete From Icons where ID = $1
`

func (q *Queries) DeleteIcon(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteIcon, id)
	return err
}

const deleteSetting = `-- name: DeleteSetting :exec
Delete From settings Where ID = $1
`

func (q *Queries) DeleteSetting(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSetting, id)
	return err
}

const deleteSource = `-- name: DeleteSource :exec
DELETE From sources where id = $1
`

func (q *Queries) DeleteSource(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSource, id)
	return err
}

const getArticleByID = `-- name: GetArticleByID :one
Select id, sourceid, tags, title, url, pubdate, video, videoheight, videowidth, thumbnail, description, authorname, authorimage from Articles
WHERE ID = $1 LIMIT 1
`

// Articles
func (q *Queries) GetArticleByID(ctx context.Context, id uuid.UUID) (Article, error) {
	row := q.db.QueryRowContext(ctx, getArticleByID, id)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Sourceid,
		&i.Tags,
		&i.Title,
		&i.Url,
		&i.Pubdate,
		&i.Video,
		&i.Videoheight,
		&i.Videowidth,
		&i.Thumbnail,
		&i.Description,
		&i.Authorname,
		&i.Authorimage,
	)
	return i, err
}

const getArticleByUrl = `-- name: GetArticleByUrl :one
Select id, sourceid, tags, title, url, pubdate, video, videoheight, videowidth, thumbnail, description, authorname, authorimage from Articles
Where Url = $1 LIMIT 1
`

func (q *Queries) GetArticleByUrl(ctx context.Context, url string) (Article, error) {
	row := q.db.QueryRowContext(ctx, getArticleByUrl, url)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Sourceid,
		&i.Tags,
		&i.Title,
		&i.Url,
		&i.Pubdate,
		&i.Video,
		&i.Videoheight,
		&i.Videowidth,
		&i.Thumbnail,
		&i.Description,
		&i.Authorname,
		&i.Authorimage,
	)
	return i, err
}

const getDiscordQueueByID = `-- name: GetDiscordQueueByID :one
Select id, articleid from DiscordQueue
Where ID = $1 LIMIT 1
`

func (q *Queries) GetDiscordQueueByID(ctx context.Context, id uuid.UUID) (Discordqueue, error) {
	row := q.db.QueryRowContext(ctx, getDiscordQueueByID, id)
	var i Discordqueue
	err := row.Scan(&i.ID, &i.Articleid)
	return i, err
}

const getDiscordQueueItems = `-- name: GetDiscordQueueItems :many
Select id, articleid from DiscordQueue LIMIT $1
`

func (q *Queries) GetDiscordQueueItems(ctx context.Context, limit int32) ([]Discordqueue, error) {
	rows, err := q.db.QueryContext(ctx, getDiscordQueueItems, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Discordqueue
	for rows.Next() {
		var i Discordqueue
		if err := rows.Scan(&i.ID, &i.Articleid); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDiscordWebHooksByID = `-- name: GetDiscordWebHooksByID :one
Select id, name, key, url, server, channel, enabled from DiscordWebHooks
Where ID = $1 LIMIT 1
`

func (q *Queries) GetDiscordWebHooksByID(ctx context.Context, id uuid.UUID) (Discordwebhook, error) {
	row := q.db.QueryRowContext(ctx, getDiscordWebHooksByID, id)
	var i Discordwebhook
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Key,
		&i.Url,
		&i.Server,
		&i.Channel,
		&i.Enabled,
	)
	return i, err
}

const getDiscordWebHooksByServer = `-- name: GetDiscordWebHooksByServer :many
Select id, name, key, url, server, channel, enabled From DiscordWebHooks
Where Server = $1
`

func (q *Queries) GetDiscordWebHooksByServer(ctx context.Context, server sql.NullString) ([]Discordwebhook, error) {
	rows, err := q.db.QueryContext(ctx, getDiscordWebHooksByServer, server)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Discordwebhook
	for rows.Next() {
		var i Discordwebhook
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Key,
			&i.Url,
			&i.Server,
			&i.Channel,
			&i.Enabled,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getIconByID = `-- name: GetIconByID :one
Select id, filename, site FROM Icons
Where ID = $1 Limit 1
`

func (q *Queries) GetIconByID(ctx context.Context, id uuid.UUID) (Icon, error) {
	row := q.db.QueryRowContext(ctx, getIconByID, id)
	var i Icon
	err := row.Scan(&i.ID, &i.Filename, &i.Site)
	return i, err
}

const getIconBySite = `-- name: GetIconBySite :one
Select id, filename, site FROM Icons
Where Site = $1 Limit 1
`

func (q *Queries) GetIconBySite(ctx context.Context, site sql.NullString) (Icon, error) {
	row := q.db.QueryRowContext(ctx, getIconBySite, site)
	var i Icon
	err := row.Scan(&i.ID, &i.Filename, &i.Site)
	return i, err
}

const getSettingByID = `-- name: GetSettingByID :one
Select id, key, value, options From settings
Where ID = $1 Limit 1
`

func (q *Queries) GetSettingByID(ctx context.Context, id uuid.UUID) (Setting, error) {
	row := q.db.QueryRowContext(ctx, getSettingByID, id)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Value,
		&i.Options,
	)
	return i, err
}

const getSettingByKey = `-- name: GetSettingByKey :one
Select id, key, value, options From settings Where 
Key = $1 Limit 1
`

func (q *Queries) GetSettingByKey(ctx context.Context, key string) (Setting, error) {
	row := q.db.QueryRowContext(ctx, getSettingByKey, key)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Value,
		&i.Options,
	)
	return i, err
}

const getSettingByValue = `-- name: GetSettingByValue :one
Select id, key, value, options From settings Where 
Value = $1 Limit 1
`

func (q *Queries) GetSettingByValue(ctx context.Context, value string) (Setting, error) {
	row := q.db.QueryRowContext(ctx, getSettingByValue, value)
	var i Setting
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Value,
		&i.Options,
	)
	return i, err
}

const getSourceByID = `-- name: GetSourceByID :one
Select id, site, name, source, type, value, enabled, url, tags From Sources where ID = $1 Limit 1
`

func (q *Queries) GetSourceByID(ctx context.Context, id uuid.UUID) (Source, error) {
	row := q.db.QueryRowContext(ctx, getSourceByID, id)
	var i Source
	err := row.Scan(
		&i.ID,
		&i.Site,
		&i.Name,
		&i.Source,
		&i.Type,
		&i.Value,
		&i.Enabled,
		&i.Url,
		&i.Tags,
	)
	return i, err
}

const getSourcesBySource = `-- name: GetSourcesBySource :many
Select id, site, name, source, type, value, enabled, url, tags From Sources where Source = $1
`

func (q *Queries) GetSourcesBySource(ctx context.Context, source sql.NullString) ([]Source, error) {
	rows, err := q.db.QueryContext(ctx, getSourcesBySource, source)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Source
	for rows.Next() {
		var i Source
		if err := rows.Scan(
			&i.ID,
			&i.Site,
			&i.Name,
			&i.Source,
			&i.Type,
			&i.Value,
			&i.Enabled,
			&i.Url,
			&i.Tags,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
