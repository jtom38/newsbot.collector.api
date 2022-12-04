package input_test

import (
	"testing"

	"github.com/jtom38/newsbot/collector/domain/model"
	"github.com/jtom38/newsbot/collector/services/input"
)

var rssRecord = model.Sources{
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
