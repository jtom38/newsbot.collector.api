package services_test

import (
	"testing"

	ffxiv "github.com/jtom38/newsbot/collector/services"
)

func TestFfxivGetParser(t *testing.T) {
	fc := ffxiv.NewFFXIVClient("na")
	_, err := fc.GetParser()
	if err != nil { panic(err) }
}

func TestFfxivPullFeed(t *testing.T) {
	fc := ffxiv.NewFFXIVClient("na")

	parser := fc.GetBrowser()
	defer parser.Close()

	links, err := fc.PullFeed(parser)
	if err != nil { panic(err) }
	if len(links) == 0 { panic("expected links to come back but got 0") }

}

func TestFfxivExtractThumbnail(t *testing.T) {
	fc := ffxiv.NewFFXIVClient("na")

	parser := fc.GetBrowser()
	defer parser.Close()

	links, err := fc.PullFeed(parser)
	if err != nil { panic(err) }

	page := fc.GetPage(parser, links[0])
	defer page.Close()

	thumb, err := fc.ExtractThumbnail(page)
	if err != nil { panic(err) }
	if thumb == "" { panic("expected a link but got nothing.")}
}

func TestFfxivExtractPubDate(t *testing.T) {
	fc := ffxiv.NewFFXIVClient("na")

	parser := fc.GetBrowser()
	defer parser.Close()

	links, err := fc.PullFeed(parser)
	if err != nil { panic(err) }

	page := fc.GetPage(parser, links[0])
	defer page.Close()

	_, err = fc.ExtractPubDate(page)
	if err != nil { panic(err) }
}

func TestFfxivExtractDescription(t *testing.T) {
	fc := ffxiv.NewFFXIVClient("na")

	parser := fc.GetBrowser()
	defer parser.Close()

	links, err := fc.PullFeed(parser)
	if err != nil { panic(err) }

	page := fc.GetPage(parser, links[0])
	defer page.Close()

	_, err = fc.ExtractDescription(page)
	if err != nil { panic(err) }
}

func TestFfxivExtractAuthor(t *testing.T) {
	fc := ffxiv.NewFFXIVClient("na")

	parser := fc.GetBrowser()
	defer parser.Close()

	links, err := fc.PullFeed(parser)
	if err != nil { panic(err) }

	page := fc.GetPage(parser, links[0])
	defer page.Close()

	author, err := fc.ExtractAuthor(page)
	if err != nil { panic(err) }
	if author == "" { panic("failed to locate the author name") }
}

func TestFfxivExtractTags(t *testing.T) {
	fc := ffxiv.NewFFXIVClient("na")

	parser := fc.GetBrowser()
	defer parser.Close()

	links, err := fc.PullFeed(parser)
	if err != nil { panic(err) }

	page := fc.GetPage(parser, links[0])
	defer page.Close()

	res, err := fc.ExtractTags(page)
	if err != nil { panic(err) }
	if res == "" {panic("failed to locate the tags")}
}

func TestFfxivExtractTitle(t *testing.T) {
	fc := ffxiv.NewFFXIVClient("na")

	parser := fc.GetBrowser()
	defer parser.Close()

	links, err := fc.PullFeed(parser)
	if err != nil { panic(err) }

	page := fc.GetPage(parser, links[0])
	defer page.Close()

	res, err := fc.ExtractTitle(page)
	if err != nil { panic(err) }
	if res == "" { panic("failed to locate the tags") }
}

func TestFFxivExtractAuthorIamge(t *testing.T) {
	fc := ffxiv.NewFFXIVClient("na")

	parser := fc.GetBrowser()
	defer parser.Close()

	links, err := fc.PullFeed(parser)
	if err != nil { panic(err) }

	page := fc.GetPage(parser, links[0])
	defer page.Close()

	res, err := fc.ExtractAuthorImage(page)
	if err != nil { panic(err) }
	if res == "" { panic("failed to locate the tags") }
}

func TestFfxivCheckSource(t *testing.T) {
	fc := ffxiv.NewFFXIVClient("na")
	fc.CheckSource()
	
}