package services_test

import (
	"testing"

	"github.com/jtom38/newsbot/collector/domain/model"
	"github.com/jtom38/newsbot/collector/services"
)

var rssRecord = model.Sources {
	ID: 1,
	Name: "ArsTechnica",
	Url: "https://feeds.arstechnica.com/arstechnica/index",
}

func TestRssClientConstructor(t *testing.T) {
	services.NewRssClient(rssRecord)
}

func TestRssGetFeed(t *testing.T) {
	client := services.NewRssClient(rssRecord)
	feed, err := client.PullFeed()
	if err != nil { t.Error(err) }
	if len(feed.Items) >= 0 { t.Error("failed to collect items from the fees")}

}