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
	cron.OpenDatabase(ctx)
	cron.CheckReddit(ctx)
}

func TestCheckYouTube(t *testing.T) {
	ctx := context.Background()
	cron.OpenDatabase(ctx)
	cron.CheckYoutube(ctx)
}

func TestCheckTwitch(t *testing.T) {
	ctx := context.Background()
	err := cron.OpenDatabase(ctx)
	if err != nil {
		t.Error(err)
	}
	err = cron.CheckTwitch(ctx)
	if err != nil {
		t.Error(err)
	}
}
