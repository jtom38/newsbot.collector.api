package input_test

import (
	"testing"

	"github.com/jtom38/newsbot/collector/domain/models"
	"github.com/jtom38/newsbot/collector/services/input"
)

var rssRecord = models.Sources{
	ID:   1,
	Name: "ArsTechnica",
	Url:  "https://feeds.arstechnica.com/arstechnica/index",
}

func TestRssClientConstructor(t *testing.T) {
	input.NewRssClient(rssRecord)
}

func TestRssGetFeed(t *testing.T) {
	client := input.NewRssClient(rssRecord)
	feed, err := client.PullFeed()
	if err != nil {
		t.Error(err)
	}
	if len(feed.Items) >= 0 {
		t.Error("failed to collect items from the fees")
	}

}
