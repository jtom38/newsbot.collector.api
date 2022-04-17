package services_test

import (
	"testing"

	"github.com/jtom38/newsbot/collector/services"
)

func TestGetContent(t *testing.T) {
	rc := services.NewRedditClient("dadjokes", 0)
	_, err := rc.GetContent()

	if err != nil { panic(err) }
}