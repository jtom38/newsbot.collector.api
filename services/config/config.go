package config

import (
	"os"
	"log"

	"github.com/joho/godotenv"
)

const (
	DB_URI string	= "DB_URI"
	
	Sql_Connection_String string = "SQL_CONNECTION_STRING"

	REDDIT_PULL_TOP = "REDDIT_PULL_TOP"
	REDDIT_PULL_HOT = "REDDIT_PULL_HOT"
	REDDIT_PULL_NSFW = "REDDIT_PULL_NSFW"

	YOUTUBE_DEBUG = "YOUTUBE_DEBUG"

	TWITCH_CLIENT_ID = "TWITCH_CLIENT_ID"
	TWITCH_CLIENT_SECRET = "TWITCH_CLIENT_SECRET"
	TWITCH_MONITOR_CLIPS = "TWITCH_MONITOR_CLIPS"
	TWITCH_MONITOR_VOD = "TWITCH_MONITOR_VOD"
)

type ConfigClient struct {}

func New() ConfigClient {
	_, err := os.Open(".env")
	if err == nil {
		loadEnvFile()
	}

	return ConfigClient{}
}

func (cc *ConfigClient) GetConfig(key string) string {
	res, filled := os.LookupEnv(key)
	if !filled {
		log.Printf("Missing the a value for '%v'.  Could generate errors.", key)
	}
	return res
}

// Use this when your ConfigClient has been opened for awhile and you want to ensure you have the most recent env changes.
func (cc *ConfigClient) RefreshEnv() {
	loadEnvFile()
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
}
