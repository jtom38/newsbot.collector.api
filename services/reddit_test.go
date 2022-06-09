package services_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services"
)

var RedditRecord database.Source = database.Source{
	ID: uuid.New(),
	Name: "dadjokes",
	Source: "reddit",
	Site: "reddit",
	Url: "https://reddit.com/r/dadjokes",
	Tags: "reddit, dadjokes",
}

func TestGetContent(t *testing.T) {
	//This test is flaky right now due to the http changes in 1.17
	rc := services.NewRedditClient(RedditRecord)
	raw, err := rc.GetContent()
	if err != nil {
		t.Error(err)
	}
	redditArticles := rc.ConvertToArticles(raw)
	for _, posts := range redditArticles {
		if posts.Title == "" {
			t.Error("Title is missing")
		}
	}
}