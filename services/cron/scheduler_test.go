package cron_test

import (
	"context"
	"testing"

	"github.com/jtom38/newsbot/collector/services/cron"
)

func TestInvokeTwitch(t *testing.T) {

}

// TODO add database mocks but not sure how to do that yet.
func TestCheckReddit(t *testing.T) {
	ctx := context.Background()
	c := cron.Cron{}
	c.CheckReddit(ctx)
}

func TestCheckYouTube(t *testing.T) {
	ctx := context.Background()
	c := cron.Cron{}
	c.CheckYoutube(ctx)
}

func TestCheckTwitch(t *testing.T) {
	ctx := context.Background()
	c := cron.Cron{}
	err := c.CheckTwitch(ctx)
	if err != nil {
		t.Error(err)
	}
}
