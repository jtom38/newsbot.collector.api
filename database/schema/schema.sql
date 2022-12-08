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
	Key 	TEXT NOT NULL,
	Value 	TEXT NOT NULL,
	Options TEXT
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
	Tags 	TEXT NOT NULL,
	Deleted BOOLEAN
);

/* This table is used to track what the Web Hook wants to have sent by Source */;
Create TABLE Subscriptions (
	ID 					uuid Primary Key,
	DiscordWebHookID  	uuid Not Null,
	SourceID         	uuid Not Null
);
