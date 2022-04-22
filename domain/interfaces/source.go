package interfaces

import (
	"github.com/go-rod/rod"
	"github.com/mmcdole/gofeed"
)

type Sources interface {
	CheckSource() error
	PullFeed() (*gofeed.Feed, error)

	GetBrowser() *rod.Browser
	GetPage(url string) *rod.Page

	ExtractThumbnail(page *rod.Page) (string, error)
	ExtractPubDate(page *rod.Page) (string, error)
}