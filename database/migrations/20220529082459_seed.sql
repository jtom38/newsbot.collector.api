-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- Enable UUID's
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Reddit Entries
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'reddit', 'dadjokes', 'reddit', 'feed', 'a', TRUE, 'https://reddit.com/r/dadjokes', 'a');

-- RSS Entries
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'steampowered', 'steam deck', 'rss', 'feed', 'a', TRUE, 'https://store.steampowered.com/feeds/news/app/1675200/?cc=US&l=english&snr=1_2108_9__2107', 'rss, steampowered, steam deck');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DELETE FROM sources where url = 'https://reddit.com/r/dadjokes';

DELETE FROM sources where url = 'https://store.steampowered.com/feeds/news/app/1675200/?cc=US&l=english&snr=1_2108_9__2107';
-- +goose StatementEnd
