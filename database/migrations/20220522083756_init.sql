-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Articles (
	ID 			uuid PRIMARY KEY,
	SourceId 	uuid not null,
	Tags 		TEXT,
	Title 		TEXT,
	Url 		TEXT,
	PubDate 	timestamp,
	Video 		TEXT,
	VideoHeight int,
	VideoWidth 	int,
	Thumbnail 	TEXT,
	Description TEXT,
	AuthorName 	TEXT,
	AuthorImage TEXT
);

CREATE Table DiscordQueue ( 
    ID 			uuid PRIMARY KEY,
    ArticleId 	uuid NOT NULL
);

CREATE Table DiscordWebHooks (
	ID 		uuid PRIMARY KEY,
	Name 	TEXT,
	Key 	TEXT,
	Url 	TEXT,
	Server 	TEXT,
	Channel TEXT,
	Enabled BIT
);

CREATE Table Icons (
	ID 			uuid PRIMARY Key,
	FileName 	TEXT,
	Site 		TEXT
);

Create Table Settings (
	ID 		uuid PRIMARY Key,
	Key 	TEXT,
	Value 	TEXT,
	Options TEXT
);

Create Table Sources (
	ID 		uuid PRIMARY Key,
	Site 	TEXT,
	Name 	TEXT,
	Source 	TEXT,
	Type 	TEXT,
	Value 	TEXT,
	Enabled BIT,
	Url 	TEXT,
	Tags 	TEXT
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop Table Articles;
Drop Table DiscordQueue;
Drop Table DiscordWebHooks;
Drop Table Icons;
Drop Table Settings;
Drop Table Sources;
-- +goose StatementEnd
