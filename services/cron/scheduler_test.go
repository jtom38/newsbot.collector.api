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
	c := cron.New(ctx)
	c.CheckReddit()
}

func TestCheckYouTube(t *testing.T) {
	ctx := context.Background()
	c := cron.New(ctx)
	c.CheckYoutube()
}

func TestCheckTwitch(t *testing.T) {
	ctx := context.Background()
	c := cron.New(ctx)
	err := c.CheckTwitch()
	if err != nil {
		t.Error(err)
	}
}
