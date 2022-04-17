package interfaces

import (
	"github.com/mmcdole/gofeed"
)

type Sources interface {
	CheckSource() error
	PullFeed() (*gofeed.Feed, error)
}