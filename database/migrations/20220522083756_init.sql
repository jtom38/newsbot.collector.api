-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE Articles (
	ID 			uuid PRIMARY KEY,
	SourceId 	uuid NOT null,
	Tags 		TEXT NOT NULL,
	Title 		TEXT NOT NULL,
	Url 		TEXT NOT NULL,
	PubDate 	timestamp NOT NULL,
	Video 		TEXT,
	VideoHeight int NOT NULL,
	VideoWidth 	int NOT NULL,
	Thumbnail 	TEXT NOT NULL,
	Description TEXT NOT NULL,
	AuthorName 	TEXT,
	AuthorImage TEXT
);

CREATE Table DiscordQueue ( 
    ID 			uuid PRIMARY KEY,
    ArticleId 	uuid NOT NULL
);

CREATE Table DiscordWebHooks (
	ID 		uuid PRIMARY KEY,
	Name 	TEXT NOT NULL, -- Defines webhook purpose
	Key 	TEXT,
	Url 	TEXT NOT NULL, -- Webhook Url
	Server 	TEXT NOT NULL, -- Defines the server its bound it. Used for refrence
	Channel TEXT NOT NULL, -- Defines the channel its bound to.  Used for refrence
	Enabled BOOLEAN NOT NULL
);

CREATE Table Icons (
	ID 			uuid PRIMARY Key,
	FileName 	TEXT NOT NULL,
	Site 		TEXT NOT NULL
);

Create Table Settings (
	ID 		uuid PRIMARY Key,
	Key 	TEXT NOT NULL, -- How you search for a entry
	Value 	TEXT NOT NULL, -- The value for one
	Options TEXT -- any notes about the entry
);

Create Table Sources (
	ID 		uuid PRIMARY Key,
	Site 	TEXT NOT NULL, -- Vanity name
	Name 	TEXT NOT NULL, -- Defines the name of the source. IE: dadjokes
	Source 	TEXT NOT NULL, -- Defines the service that will use this reocrd. IE reddit or youtube
	Type 	TEXT NOT NULL, -- Defines what kind of feed this is.  feed, user, tag
	Value 	TEXT,
	Enabled BOOLEAN NOT NULL,
	Url 	TEXT NOT NULL,
	Tags 	TEXT NOT NULL
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
