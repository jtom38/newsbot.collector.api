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
	Key 	TEXT NOT NULL,
	Value 	TEXT NOT NULL,
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

