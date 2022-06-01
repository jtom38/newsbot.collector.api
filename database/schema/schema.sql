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
	Name 	TEXT,
	Key 	TEXT,
	Url 	TEXT,
	Server 	TEXT,
	Channel TEXT,
	Enabled BOOLEAN
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
	Enabled BOOLEAN,
	Url 	TEXT,
	Tags 	TEXT
);

