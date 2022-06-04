package services_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services"
)

var record database.Source = database.Source{
	ID: uuid.New(),
	Name: "dadjokes",
	Source: "reddit",
	Site: "reddit",
}

func TestGetContent(t *testing.T) {
	//This test is flaky right now due to the http changes in 1.17
	rc := services.NewRedditClient(record)
	_, err := rc.GetContent()
	if err != nil {
		t.Error(err)
	}
}