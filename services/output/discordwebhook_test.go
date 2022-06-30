package output_test

import (
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services/output"
)

var article database.Article = database.Article{
	ID: uuid.New(),
	Sourceid: uuid.New(),
	Tags: "unit, testing",
	Title: "Demo",
	Url: "https://github.com/jtom38/newsbot.collector.api",
	Pubdate: time.Now(),
	Videoheight: 0,
	Videowidth: 0,
	Description: "Hello World",
}

func getWebhook() ([]string, error){
	var endpoints []string

	_, err := os.Open(".env")
	if err != nil {
		return endpoints, err
	}
		
	err = godotenv.Load()
	if err != nil {
		return endpoints, err
	}

	res := os.Getenv("TESTS_DISCORD_WEBHOOK")
	if res == "" {
		return endpoints, errors.New("TESTS_DISCORD_WEBHOOK is missing")
	}
	endpoints = strings.Split(res, "")
	return endpoints, nil
}

func TestNewDiscordWebHookContainsSubscriptions(t *testing.T) {
	hook, err := getWebhook()
	if err != nil {
		t.Error(err)
	}
	d := output.NewDiscordWebHookMessage(hook, article)
	if len(d.Subscriptions) == 0 {
		t.Error("no subscriptions found")
	}
}

func TestDiscordMessageContainsTitle(t *testing.T) {
	hook, err := getWebhook()
	if err != nil {
		t.Error(err)
	}
	d := output.NewDiscordWebHookMessage(hook, article)
	err = d.GeneratePayload()
	if err != nil {
		t.Error(err)
	}
	if d.Message.Embeds[0].Title == "" {
		t.Error("no title was found ")
	}
}