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
	_, err := yc.GetParser(yc.Url)
	if err != nil { panic(err) }
}

func TestGetChannelId(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)
	parser, err := yc.GetParser(yc.Url)
	if err != nil { panic(err) }

	_, err = yc.GetChannelId(parser)
	if err != nil { panic(err) }
}

func TestPullFeed(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)
	parser, err := yc.GetParser(yc.Url)
	if err != nil { panic(err) }

	_, err = yc.GetChannelId(parser)
	if err != nil { panic(err) }

	_, err = yc.PullFeed()
	if err != nil { panic(err) }
}

func TestGetAvatarUri(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)
	res, err := yc.GetAvatarUri()
	if err != nil { panic(err) }
	if res == "" { panic(services.ErrAvatarMissing)}
}

func TestGetVideoTags(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)

	var videoUri = "https://www.youtube.com/watch?v=k_sQEXOBe68"

	parser, err := yc.GetParser(videoUri)
	if err != nil { panic(err) }

	tags, err := yc.GetTags(parser)
	if err == nil && tags == "" { panic("err was empty but value was missing.")}
	if err != nil { panic(err) }
}

func TestGetChannelTags(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)

	parser, err := yc.GetParser(yc.Url)
	if err != nil { panic(err) }

	tags, err := yc.GetTags(parser)
	if err == nil && tags == "" { panic("no err but expected value was missing.")}
	if err != nil { panic(err) }
}

func TestGetVideoThumbnail(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)
	parser, err := yc.GetParser("https://www.youtube.com/watch?v=k_sQEXOBe68")
	if err != nil {panic(err) }

	thumb, err := yc.GetVideoThumbnail(parser)
	if err == nil && thumb == "" { panic("no err but expected result was missing")}
	if err != nil { panic(err) }

}

func TestCheckSource(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)
	err := yc.CheckSource()
	if err != nil { panic(err) }

}

func TestCheckUriCache(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)
	item := "demo"

	services.YoutubeUriCache = append(services.YoutubeUriCache, &item)
	res := yc.CheckUriCache(&item)
	if res == false { panic("expected a value to come back")}
}

func TestCheckUriCacheFails(t *testing.T) {
	yc := services.NewYoutubeClient(
		0,
		"https://youtube.com/gamegrumps",
	)
	item := "demo1"
	
	res := yc.CheckUriCache(&item)
	if res == true { panic("expected no value to come back")}

}