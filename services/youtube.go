package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/mmcdole/gofeed"

	"github.com/jtom38/newsbot/collector/domain/model"
)

type YoutubeClient struct {
	SourceID  uint
	Url       string
	ChannelID string
	AvatarUri string
	//config YoutubeConfig
}

type YoutubeConfig struct {

}

const YOUTUBE_FEED_URL string = "https://www.youtube.com/feeds/videos.xml?channel_id="

func NewYoutubeClient(SourceID uint, Url string) YoutubeClient {
	return YoutubeClient{
		SourceID: SourceID,
		Url:      Url,
	}
}

// CheckSource will go and run all the commands needed to process a source.
func (yc *YoutubeClient) CheckSource() error {
	docParser, err := yc.GetParser(yc.Url)
	if err != nil { return err }

	// Check cache/db for existing value
	// If we have the value, skip
	channelId, err := yc.GetChannelId(docParser)
	if err != nil { return err }
	yc.ChannelID = channelId
	
	// Check the cache/db forthe value.
	// if we have the value, skip
	avatar, err := yc.GetAvatarUri()
	if err != nil { return err }
	yc.AvatarUri = avatar

	feed, err := yc.PullFeed()
	if err != nil { return err }
	
	_, err = yc.CheckForNewPosts(feed)
	if err != nil { return err }
	
	return nil
}

func (yc *YoutubeClient) GetParser(uri string) (*goquery.Document, error) {
	html, err := http.Get(uri)
	if err != nil {
		log.Println(err)
	}
	defer html.Body.Close()

	doc, err := goquery.NewDocumentFromReader(html.Body)
	if err != nil {
		return nil, err
	}
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
	return "", errors.New("unable to find the channelId on the requested page")
}

// This will parse the page to find the current Avatar of the channel.
func (yc *YoutubeClient) GetAvatarUri() (string, error) {
	var AvatarUri string

	browser := rod.New().MustConnect()
	page := browser.MustPage(yc.Url)
	
	res := page.MustElement("#channel-header-container > yt-img-shadow:nth-child(1) > img:nth-child(1)").MustAttribute("src")

	if *res == "" || res == nil {
		return AvatarUri, errors.New("unable to find the avatar on the page")
	}
		
	defer browser.Close()
	defer page.Close()
	return AvatarUri, nil
}

// This will parse and look for the tags that has been defined by the user.
func (yc *YoutubeClient) GetTags(parser *goquery.Document) (string, error) {
	meta := parser.Find("meta")

	for _, item := range meta.Nodes {
		if item.Attr[0].Val == "keywords" {
			res := item.Attr[1].Val
			return res, nil
		}
	}
	return "", errors.New("unable to find the tags on the video")
}

func (yc *YoutubeClient) GetVideoThumbnail(parser *goquery.Document) (string, error) {
	meta := parser.Find("meta")

	for _, item := range meta.Nodes {
		if item.Attr[0].Val == "og:image" {
			res := item.Attr[1].Val
			return res, nil
		}
	}
	return "", errors.New("unable to find the video thumbnail on a youtube video")
}

// This will pull the RSS feed items and return the results
func (yc *YoutubeClient) PullFeed() (*gofeed.Feed, error) {
	feedUri := fmt.Sprintf("%v%v", YOUTUBE_FEED_URL, yc.ChannelID)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedUri)
	if err != nil {
		return nil, err
	}

	return feed, nil
}

// CheckForNewPosts will talk to the Database to see if it has a record for the posts that have been extracted.
// If the post does not exist, it will be added.
func (yc *YoutubeClient) CheckForNewPosts(feed *gofeed.Feed) ([]*gofeed.Item, error) {
	var newPosts []*gofeed.Item
	for _, item := range feed.Items {
		// Check the cache/db to see if this URI has been seen already
		


		// if its new, process it.
		newPosts = append(newPosts, item)

		// Add the post to local cache

		article := yc.ConvertToArticle(item)
		log.Println(article)
	}

	return newPosts, nil
}

func (yc *YoutubeClient) ConvertToArticle(item *gofeed.Item) model.Articles {
	parser, err := yc.GetParser(item.Link)
	if err != nil { 
		log.Printf("Unable to process %v, submit this link as an issue.\n", item.Link)
		
	}

	tags, err := yc.GetTags(parser)
	if err != nil {
		log.Printf("Unable to find the tags on %v, submit this link as an issue.\n", item.Link)
	}

	thumb, err := yc.GetVideoThumbnail(parser)
	if err != nil { 
		log.Printf("Unable to find the thumbnail at %v, submit this link as an issue.\n", item.Link)
	}

	var article = model.Articles{
		SourceID:    yc.SourceID,
		Tags:        tags,
		Title:       item.Title,
		Url:         item.Link,
		PubDate:     *item.PublishedParsed,
		Thumbnail:   thumb,
		Description: item.Description,
		AuthorName:  item.Author.Name,
		AuthorImage: yc.AvatarUri,
	}
	return article
}