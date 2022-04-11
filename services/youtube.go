package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/jtom38/newsbot/collector/domain/model"
	"github.com/mmcdole/gofeed"
)

type YoutubeClient struct {
	SourceID uint
	url string
	ChannelID string
}

const YOUTUBE_FEED_URL string = "https://www.youtube.com/feeds/videos.xml?channel_id="

func NewYoutubeClient(SourceID uint, Url string) YoutubeClient {
	return YoutubeClient{
		SourceID: SourceID,
		url: Url,
	}
}

// CheckSource will go and run all the commands needed to process a source.
func (yc *YoutubeClient) CheckSource() error {
	docParser, err := yc.GetPageParser()
	if err != nil { return err }

	_, err = yc.GetChannelId(docParser)
	if err != nil { return err }

	_, err = yc.GetAvatarUri(docParser)
	if err != nil { return err }

	feed, err := yc.PullFeed()
	if err != nil { return err }

	err = yc.CheckForNewPosts(feed)
	if err != nil { return err }

	return nil
}

func (yc *YoutubeClient) GetPageParser() (*goquery.Document, error) {
	html, err := http.Get(yc.url)
	if err != nil { log.Println(err) }
	defer html.Body.Close()

	doc, err := goquery.NewDocumentFromReader(html.Body)
	if err != nil { return nil, err }
	return doc, nil
}

// This pulls the youtube page and finds the ChannelID.
// This value is required to generate the RSS feed URI
func (yc *YoutubeClient) GetChannelId(doc *goquery.Document) (string, error) {
	meta := doc.Find("meta")
	for _, item := range meta.Nodes {

		if item.Attr[0].Val == "channelId" {
			yc.ChannelID = item.Attr[1].Val
			return yc.ChannelID, nil
		}
	}
	return "", errors.New("Unable to find the channelId on the requested page")
}

// This will parse the page to find the current Avatar of the channel.
func (yc *YoutubeClient) GetAvatarUri(doc *goquery.Document) (string, error) {
	return "", errors.New("Not Implemented")
}

// This will parse and look for the tags that has been defined by the user.
func (yc *YoutubeClient) GetTags() error {
	return errors.New("Not Implemented")
}

// This will pull the RSS feed items and return the results
func (yc *YoutubeClient) PullFeed() (*gofeed.Feed, error) {
	feedUri := fmt.Sprintf("%v%v", YOUTUBE_FEED_URL, yc.ChannelID)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedUri)
	if err != nil { return nil, err }

	return feed, nil
}

// CheckForNewPosts will talk to the Database to see if it has a record for the posts that have been extracted.
// If the post does not exist, it will be added.
func (yc *YoutubeClient) CheckForNewPosts(feed *gofeed.Feed) error {
	for _, item := range feed.Items {
		article := yc.convertToArticles(item)
		log.Println(article)
	}

	return nil
}


func (yc *YoutubeClient) convertToArticles(item *gofeed.Item) model.Articles {
	var article = model.Articles {
		SourceID: yc.SourceID,
		Tags: "",
		Title: item.Title,
		Url: item.Link,
		PubDate: *item.PublishedParsed,
		Thumbnail: item.Image.URL,
		Description: item.Description,
		AuthorName: item.Author.Name,
		AuthorImage: "",
	}
	return article
}