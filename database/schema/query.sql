/* Articles */
-- name: GetArticleByID :one
Select * from Articles
WHERE ID = $1 LIMIT 1;

-- name: GetArticleByUrl :one
Select * from Articles
Where Url = $1 LIMIT 1;

-- name: ListArticles :many
Select * From articles 
Order By PubDate DESC
offset $2
fetch next $1 rows only;

-- name: ListArticlesByDate :many
Select * From articles 
ORDER BY pubdate desc 
Limit $1;

-- name: GetArticlesBySource :many
select * from articles
INNER join sources on articles.sourceid=Sources.ID
where site = $1;

-- name: ListNewArticlesBySourceId :many
SELECT * FROM articles
Where sourceid = $1
ORDER BY pubdate desc
offset $3
fetch next $2 rows only;

-- name: ListOldestArticlesBySourceId :many
SELECT * FROM articles
Where sourceid = $1
ORDER BY pubdate asc
offset $3
fetch next $2 rows only;


-- name: ListArticlesBySourceId :many
Select * From articles
Where sourceid = $1 
Limit 50;

-- name: GetArticlesBySourceName :many
select 
articles.ID, articles.SourceId, articles.Tags, articles.Title, articles.Url, articles.PubDate, articles.Video, articles.VideoHeight, articles.VideoWidth, articles.Thumbnail, articles.Description, articles.AuthorName, articles.AuthorImage, sources.source, sources.name
From articles
Left Join sources
On articles.sourceid = sources.id
Where name = $1;

-- name: ListArticlesByPage :many
select * from articles
order by pubdate desc
offset $2
fetch next $1 rows only;

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

-- name: ListDiscordQueueItems :many
Select * from DiscordQueue LIMIT $1;

/* DiscordWebHooks */
-- name: CreateDiscordWebHook :exec
Insert Into DiscordWebHooks
(ID, Url, Server, Channel, Enabled)
Values
($1, $2, $3, $4, $5);

-- name: GetDiscordWebHooksByID :one
Select * from DiscordWebHooks
Where ID = $1 LIMIT 1;

-- name: ListDiscordWebHooksByServer :many
Select * From DiscordWebHooks
Where Server = $1;

-- name: GetDiscordWebHooksByServerAndChannel :many
SELECT * FROM DiscordWebHooks
WHERE Server = $1 and Channel = $2;

-- name: GetDiscordWebHookByUrl :one
Select * From DiscordWebHooks Where url = $1;

-- name: ListDiscordWebhooks :many
Select * From discordwebhooks LIMIT $1;

-- name: DeleteDiscordWebHooks :exec
Delete From discordwebhooks Where ID = $1;

-- name: DisableDiscordWebHook :exec
Update discordwebhooks Set Enabled = FALSE where ID = $1;

-- name: EnableDiscordWebHook :exec
Update discordwebhooks Set Enabled = TRUE where ID = $1;

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

-- name: GetSourceByName :one
Select * from Sources where name = $1 Limit 1;

-- name: GetSourceByNameAndSource :one
Select * from Sources WHERE name = $1 and source = $2;

-- name: ListSources :many
Select * From Sources Limit $1;

-- name: ListSourcesBySource :many
Select * From Sources where Source = $1;

-- name: DeleteSource :exec
UPDATE Sources Set Disabled = TRUE where id = $1;

-- name: DisableSource :exec
Update Sources Set Enabled = FALSE where ID = $1;

-- name: EnableSource :exec
Update Sources Set Enabled = TRUE where ID = $1;


/* Subscriptions */

-- name: CreateSubscription :exec
Insert Into subscriptions (ID, DiscordWebHookId, SourceId) Values ($1, $2, $3);

-- name: ListSubscriptions :many
Select * From subscriptions Limit $1;

-- name: ListSubscriptionsBySourceId :many
Select * From subscriptions where sourceid = $1;

-- name: QuerySubscriptions :one
Select * From subscriptions Where discordwebhookid = $1 and sourceid = $2 Limit 1;

-- name: GetSubscriptionsBySourceID :many
Select * From subscriptions Where sourceid = $1;

-- name: GetSubscriptionsByDiscordWebHookId :many
Select * from subscriptions Where discordwebhookid = $1;

-- name: DeleteSubscription :exec
Delete From subscriptions Where id = $1;