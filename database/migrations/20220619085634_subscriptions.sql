-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
Create TABLE Subscriptions (
	ID 			      uuid Primary Key,
	DiscordWebHookID  uuid Not Null,
	SourceID         uuid Not Null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
Drop Table Subscriptions;
-- +goose StatementEnd
