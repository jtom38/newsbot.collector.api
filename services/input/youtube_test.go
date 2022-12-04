package input_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services/input"
)

var YouTubeRecord database.Source = database.Source{
	ID:     uuid.New(),
	Name:   "dadjokes",
	Source: "reddit",
	Site:   "reddit",
	Url:    "https://youtube.com/gamegrumps",
}

func TestGetPageParser(t *testing.T) {
	yc := input.NewYoutubeClient(YouTubeRecord)
	_, err := yc.GetParser(YouTubeRecord.Url)
	if err != nil {
		panic(err)
	}
}

func TestGetChannelId(t *testing.T) {
	yc := input.NewYoutubeClient(YouTubeRecord)
	parser, err := yc.GetParser(YouTubeRecord.Url)
	if err != nil {
		panic(err)
	}

	_, err = yc.GetChannelId(parser)
	if err != nil {
		panic(err)
	}
}

func TestPullFeed(t *testing.T) {
	yc := input.NewYoutubeClient(YouTubeRecord)
	parser, err := yc.GetParser(YouTubeRecord.Url)
	if err != nil {
		panic(err)
	}

	_, err = yc.GetChannelId(parser)
	if err != nil {
		panic(err)
	}

	_, err = yc.PullFeed()
	if err != nil {
		panic(err)
	}
}

func TestGetAvatarUri(t *testing.T) {
	yc := input.NewYoutubeClient(YouTubeRecord)
	res, err := yc.GetAvatarUri()
	if err != nil {
		panic(err)
	}
	if res == "" {
		panic(input.ErrMissingAuthorImage)
	}
}

func TestGetVideoTags(t *testing.T) {
	yc := input.NewYoutubeClient(YouTubeRecord)

	var videoUri = "https://www.youtube.com/watch?v=k_sQEXOBe68"

	parser, err := yc.GetParser(videoUri)
	if err != nil {
		panic(err)
	}

	tags, err := yc.GetTags(parser)
	if err == nil && tags == "" {
		panic("err was empty but value was missing.")
	}
	if err != nil {
		panic(err)
	}
}

func TestGetChannelTags(t *testing.T) {
	yc := input.NewYoutubeClient(YouTubeRecord)

	parser, err := yc.GetParser(YouTubeRecord.Url)
	if err != nil {
		panic(err)
	}

	tags, err := yc.GetTags(parser)
	if err == nil && tags == "" {
		panic("no err but expected value was missing.")
	}
	if err != nil {
		panic(err)
	}
}

func TestGetVideoThumbnail(t *testing.T) {
	yc := input.NewYoutubeClient(YouTubeRecord)
	parser, err := yc.GetParser("https://www.youtube.com/watch?v=k_sQEXOBe68")
	if err != nil {
		panic(err)
	}

	thumb, err := yc.GetVideoThumbnail(parser)
	if err == nil && thumb == "" {
		panic("no err but expected result was missing")
	}
	if err != nil {
		panic(err)
	}

}

func TestCheckSource(t *testing.T) {
	yc := input.NewYoutubeClient(YouTubeRecord)
	_, err := yc.GetContent()
	if err != nil {
		panic(err)
	}
}

func TestCheckUriCache(t *testing.T) {
	yc := input.NewYoutubeClient(YouTubeRecord)
	item := "demo"

	input.YoutubeUriCache = append(input.YoutubeUriCache, &item)
	res := yc.CheckUriCache(&item)
	if res == false {
		panic("expected a value to come back")
	}
}

func TestCheckUriCacheFails(t *testing.T) {
	yc := input.NewYoutubeClient(YouTubeRecord)
	item := "demo1"

	res := yc.CheckUriCache(&item)
	if res == true {
		panic("expected no value to come back")
	}

}
