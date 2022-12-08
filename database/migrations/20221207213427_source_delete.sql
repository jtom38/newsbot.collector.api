-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE sources Add COLUMN Deleted BOOLEAN;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE sources Drop Deleted Deleted BOOLEAN;
-- +goose StatementEnd
