package services_test

import (
	"testing"

	"github.com/jtom38/newsbot/collector/services"
)

func TestGetPageParser(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)
	_, err := yc.GetPageParser()
	if err != nil { panic(err) }
}

func TestGetChannelId(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)
	parser, err := yc.GetPageParser()
	if err != nil { panic(err) }

	_, err = yc.GetChannelId(parser)
	if err != nil { panic(err) }
}

func TestPullFeed(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)
	parser, err := yc.GetPageParser()
	if err != nil { panic(err) }

	_, err = yc.GetChannelId(parser)
	if err != nil { panic(err) }

	_, err = yc.PullFeed()
	if err != nil { panic(err) }
}