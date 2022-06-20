-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
Create TABLE Subscriptions (
	ID 			      uuid Primary Key,
	DiscordWebHookID  uuid Not Null,
	SourceID         uuid Not Null
);

ALTER TABLE discordwebhooks drop COLUMN Name;
ALTER TABLE discordwebhooks drop COLUMN Key;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
Drop Table Subscriptions;
ALTER TABLE discordwebhooks Add COLUMN Name TEXT;
--ALTER TABLE discordwebhooks Add COLUMN Key TEXT;
-- +goose StatementEnd
