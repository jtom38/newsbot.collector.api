/* Articles */
-- name: GetArticleByID :one
Select * from Articles
WHERE ID = $1 LIMIT 1;

-- name: GetArticleByUrl :one
Select * from Articles
Where Url = $1 LIMIT 1;

-- name: CreateArticle :exec
INSERT INTO Articles 
(ID, SourceId, Tags, Title, Url, PubDate, Video, VideoHeight, VideoWidth, Thumbnail, Description, AuthorName, AuthorImage)
Values
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);


/* DiscordQueue */
-- name: CreateDiscordQueue :exec
Insert into DiscordQueue
(ID, ArticleId)
Values
($1, $2);

-- name: GetDiscordQueueByID :one
Select * from DiscordQueue
Where ID = $1 LIMIT 1;

-- name: DeleteDiscordQueueItem :exec
Delete From DiscordQueue Where ID = $1;

-- name: GetDiscordQueueItems :many
Select * from DiscordQueue LIMIT $1;


/* DiscordWebHooks */
-- name: CreateDiscordWebHook :exec
Insert Into DiscordWebHooks
(ID, Name, Key, Url, Server, Channel, Enabled)
Values
($1, $2, $3, $4, $5, $6, $7);

-- name: GetDiscordWebHooksByID :one
Select * from DiscordWebHooks
Where ID = $1 LIMIT 1;

-- name: ListDiscordWebHooksByServer :many
Select * From DiscordWebHooks
Where Server = $1;

-- name: DeleteDiscordWebHooks :exec
Delete From discordwebhooks Where ID = $1;


/* Icons */

-- name: CreateIcon :exec
INSERT INTO Icons
(ID, FileName, Site)
VALUES
($1,$2,$3);

-- name: GetIconByID :one
Select * FROM Icons
Where ID = $1 Limit 1;

-- name: GetIconBySite :one
Select * FROM Icons
Where Site = $1 Limit 1;

-- name: DeleteIcon :exec
Delete From Icons where ID = $1;

/* Settings */

-- name: CreateSettings :one
Insert Into settings
(ID, Key, Value, OPTIONS)
Values
($1,$2,$3,$4)
RETURNING *;

-- name: GetSettingByID :one
Select * From settings
Where ID = $1 Limit 1;

-- name: GetSettingByKey :one
Select * From settings Where 
Key = $1 Limit 1;

-- name: GetSettingByValue :one
Select * From settings Where 
Value = $1 Limit 1;

-- name: DeleteSetting :exec
Delete From settings Where ID = $1;

/* Sources */

-- name: CreateSource :exec
Insert Into Sources
(ID, Site, Name, Source, Type, Value, Enabled, Url, Tags)
Values
($1,$2,$3,$4,$5,$6,$7,$8,$9);

-- name: GetSourceByID :one
Select * From Sources where ID = $1 Limit 1;

-- name: ListSourcesBySource :many
Select * From Sources where Source = $1;

-- name: DeleteSource :exec
DELETE From sources where id = $1;
