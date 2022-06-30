package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	DB_URI string	= "DB_URI"
	
	Sql_Connection_String string = "SQL_CONNECTION_STRING"

	FEATURE_ENABLE_REDDIT_BACKEND = "FEATURE_ENABLE_REDDIT_BACKEND"
	REDDIT_PULL_TOP = "REDDIT_PULL_TOP"
	REDDIT_PULL_HOT = "REDDIT_PULL_HOT"
	REDDIT_PULL_NSFW = "REDDIT_PULL_NSFW"

	FEATURE_ENABLE_YOUTUBE_BACKEND = "FEATURE_ENABLE_YOUTUBE_BACKEND"
	YOUTUBE_DEBUG = "YOUTUBE_DEBUG"

	FEATURE_ENABLE_TWITCH_BACKEND = "FEATURE_ENABLE_TWITCH_BACKEND"
	TWITCH_CLIENT_ID = "TWITCH_CLIENT_ID"
	TWITCH_CLIENT_SECRET = "TWITCH_CLIENT_SECRET"
	TWITCH_MONITOR_CLIPS = "TWITCH_MONITOR_CLIPS"
	TWITCH_MONITOR_VOD = "TWITCH_MONITOR_VOD"

	FEATURE_ENABLE_FFXIV_BACKEND = "FEATURE_ENABLE_FFXIV_BACKEND"

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

func (cc *ConfigClient) GetFeature(flag string) (bool, error) {
	cc.RefreshEnv()

	res, filled := os.LookupEnv(flag)
	if !filled {
		errorMessage := fmt.Sprintf("'%v' was not found", flag)
		return false, errors.New(errorMessage)
	}

	b, err := strconv.ParseBool(res)
	if err != nil {
		return false, err
	}
	return b, nil
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
