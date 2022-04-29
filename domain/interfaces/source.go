package interfaces

import (
	"github.com/go-rod/rod"
	"github.com/mmcdole/gofeed"
)

type Sources interface {
	CheckSource() error
	PullFeed() (*gofeed.Feed, error)

	GetBrowser() *rod.Browser
	GetPage(parser *rod.Browser, url string) *rod.Page

	ExtractThumbnail(page *rod.Page) (string, error)
	ExtractPubDate(page *rod.Page) (string, error)
	ExtractDescription(page *rod.Page) (string, error)
	ExtractAuthor(page *rod.Page) (string, error)
	ExtractAuthorImage(page *rod.Page) (string, error)
	ExtractTags(page *rod.Page) (string, error)
	ExtractTitle(page *rod.Page) (string, error)
}

