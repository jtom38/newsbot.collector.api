package input

import (
	"fmt"
	"log"

	"github.com/jtom38/newsbot/collector/domain/model"
	"github.com/jtom38/newsbot/collector/services/cache"
	"github.com/mmcdole/gofeed"
)

type rssClient struct {
	SourceRecord model.Sources
}

func NewRssClient(sourceRecord model.Sources) rssClient {
	client := rssClient{
		SourceRecord: sourceRecord,
	}

	return client
}

//func (rc rssClient) ReplaceSourceRecord(source model.Sources) {
//rc.SourceRecord = source
//}

func (rc rssClient) getCacheGroup() string {
	return fmt.Sprintf("rss-%v", rc.SourceRecord.Name)
}

func (rc rssClient) GetContent() error {
	feed, err := rc.PullFeed()
	if err != nil {
		return err
	}

	cacheClient := cache.NewCacheClient(rc.getCacheGroup())

	for _, item := range feed.Items {
		log.Println(item)

		cacheClient.FindByValue(item.Link)

	}

	return nil
}

func (rc rssClient) PullFeed() (*gofeed.Feed, error) {
	feedUri := fmt.Sprintf("%v", rc.SourceRecord.Url)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedUri)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
