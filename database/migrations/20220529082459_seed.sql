-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- Enable UUID's
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Final Fantasy XIV Entries
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'ffxiv', 'Final Fantasy XIV - NA', 'ffxiv', 'scrape', 'a', TRUE, 'https://na.finalfantasyxiv.com/lodestone/', 'ffxiv, final, fantasy, xiv, na, lodestone');
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'ffxiv', 'Final Fantasy XIV - JP', 'ffxiv', 'scrape', 'a', FALSE, 'https://jp.finalfantasyxiv.com/lodestone/', 'ffxiv, final, fantasy, xiv, jp, lodestone');
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'ffxiv', 'Final Fantasy XIV - EU', 'ffxiv', 'scrape', 'a', FALSE, 'https://eu.finalfantasyxiv.com/lodestone/', 'ffxiv, final, fantasy, xiv, eu, lodestone');
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'ffxiv', 'Final Fantasy XIV - FR', 'ffxiv', 'scrape', 'a', FALSE, 'https://fr.finalfantasyxiv.com/lodestone/', 'ffxiv, final, fantasy, xiv, fr, lodestone');
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'ffxiv', 'Final Fantasy XIV - DE', 'ffxiv', 'scrape', 'a', FALSE, 'https://de.finalfantasyxiv.com/lodestone/', 'ffxiv, final, fantasy, xiv, de, lodestone');

-- Reddit Entries
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'reddit', 'dadjokes', 'reddit', 'feed', 'a', TRUE, 'https://reddit.com/r/dadjokes', 'reddit, dadjokes');
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'reddit', 'steamdeck', 'reddit', 'feed', 'a', TRUE, 'https://reddit.com/r/steamdeck', 'reddit, steam deck, steam, deck');

-- Youtube Entries
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'youtube', 'Game Grumps', 'youtube', 'feed', 'a', TRUE, 'https://www.youtube.com/user/GameGrumps', 'youtube, game grumps, game, grumps');

-- RSS Entries
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'steampowered', 'steam deck', 'rss', 'feed', 'a', TRUE, 'https://store.steampowered.com/feeds/news/app/1675200/?cc=US&l=english&snr=1_2108_9__2107', 'rss, steampowered, steam, deck, steam deck');

-- Twitch Entries 
INSERT INTO sources VALUES 
(uuid_generate_v4(), 'twitch', 'Nintendo', 'twitch', 'api', 'a', TRUE, 'https://twitch.tv/nintendo', 'twitch, nintendo');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
--SELECT 'down SQL query';

DELETE FROM sources where source = 'reddit' and name = 'dadjokes';
DELETE FROM sources where source = 'reddit' and name = 'steamdeck';
DELETE FROM sources where source = 'ffxiv';
DELETE FROM sources WHERE source = 'twitch' and name = 'Nintendo';
DELETE FROM sources WHERE source = 'youtube' and name = 'Game Grumps';
DELETE FROM SOURCES WHERE source = 'rss' and name = 'steam deck';
-- +goose StatementEnd
