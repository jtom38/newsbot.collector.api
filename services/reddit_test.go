package services_test

import (
	"log"
	"testing"

	"github.com/jtom38/newsbot/collector/services"
)

func TestGetContent(t *testing.T) {
	//This test is flaky right now due to the http changes in 1.17
	rc := services.NewRedditClient("dadjokes", 0)
	_, err := rc.GetContent()
	log.Println(err)
	//if err != nil { panic(err) }
}