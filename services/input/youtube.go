package input

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	//"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/mmcdole/gofeed"

	"github.com/jtom38/newsbot/collector/database"
)

type YoutubeClient struct {
	record database.Source

	// internal variables at time of collection
	channelID string
	avatarUri string

	// config
	//debug bool

	// cache config
	cacheGroup string
}

var (
	// This is a local slice to store what URI's have been seen to remove extra calls to the DB
	YoutubeUriCache            []*string
	ErrYoutubeChannelIdMissing = errors.New("unable to find the channelId on the requested page")
)

const YOUTUBE_FEED_URL string = "https://www.youtube.com/feeds/videos.xml?channel_id="

func NewYoutubeClient(Record database.Source) YoutubeClient {
	yc := YoutubeClient{
		record:     Record,
		cacheGroup: "youtube",
	}
	return yc
}

// CheckSource will go and run all the commands needed to process a source.
func (yc *YoutubeClient) GetContent() ([]database.Article, error) {
	var items []database.Article
	docParser, err := yc.GetParser(yc.record.Url)
	if err != nil {
		return items, err
	}

	// Check cache/db for existing value
	// If we have the value, skip
	//channelId, err := yc.extractChannelId()
	channelId, err := yc.GetChannelId(docParser)
	if err != nil {
		return items, err
	}
	if channelId == "" {
		return items, ErrYoutubeChannelIdMissing
	}
	yc.channelID = channelId

	// Check the cache/db forthe value.
	// if we have the value, skip
	avatar, err := yc.GetAvatarUri()
	if err != nil {
		return items, err
	}
	if avatar == "" {
		return items, ErrMissingAuthorImage
	}
	yc.avatarUri = avatar

	feed, err := yc.PullFeed()
	if err != nil {
		return items, err
	}

	newPosts, err := yc.CheckForNewPosts(feed)
	if err != nil {
		return items, err
	}

	for _, item := range newPosts {

		article := yc.ConvertToArticle(item)
		items = append(items, article)

		YoutubeUriCache = append(YoutubeUriCache, &item.Link)

		// Add the post to local cache
		//log.Println(article)
	}

	return items, nil
}

func (yc *YoutubeClient) GetBrowser() (*rod.Browser, error) {
	//browser := rod.New().MustConnect()
	var browser *rod.Browser
	if path, exists := launcher.LookPath(); exists {
		u, err := launcher.New().Bin(path).Launch()
		if err != nil {
			return browser, err
		}
		
		browser = rod.New().ControlURL(u).MustConnect()
	}
	return browser, nil
}

func (yc *YoutubeClient) GetPage(parser *rod.Browser, url string) *rod.Page {
	page := parser.MustPage(url)
	return page
}

func (yc *YoutubeClient) GetParser(uri string) (*goquery.Document, error) {
	html, err := http.Get(uri)
	if err != nil {
		log.Printf("[YouTube] %v", err)
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
			yc.channelID = item.Attr[1].Val
			return yc.channelID, nil
		}
	}
	return "", ErrYoutubeChannelIdMissing
}

// This pulls the youtube page and finds the ChannelID.
// This value is required to generate the RSS feed URI
//func (yc *YoutubeClient) extractChannelId(page *rod.Page) (string, error) {

//}

// This will parse the page to find the current Avatar of the channel.
func (yc *YoutubeClient) GetAvatarUri() (string, error) {
	var AvatarUri string

	//browser := rod.New().MustConnect()
	browser, err := yc.GetBrowser()
	if err != nil {
		return "", err
	}

	page := browser.MustPage(yc.record.Url)

	res := page.MustElement("#channel-header-container > yt-img-shadow:nth-child(1) > img:nth-child(1)").MustAttribute("src")

	if *res == "" || res == nil {
		return AvatarUri, ErrMissingAuthorImage
	}

	AvatarUri = *res

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
	return "", ErrMissingTags
}

func (yc *YoutubeClient) GetVideoThumbnail(parser *goquery.Document) (string, error) {
	meta := parser.Find("meta")

	for _, item := range meta.Nodes {
		if item.Attr[0].Val == "og:image" {
			res := item.Attr[1].Val
			return res, nil
		}
	}
	return "", ErrMissingThumbnail
}

// This will pull the RSS feed items and return the results
func (yc *YoutubeClient) PullFeed() (*gofeed.Feed, error) {
	feedUri := fmt.Sprintf("%v%v", YOUTUBE_FEED_URL, yc.channelID)
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
		uriExists := yc.CheckUriCache(&item.Link)
		if uriExists {
			continue
		}

		//TODO Check the DB if the cache is not aware
		//TODO If the db knew about it, append it to the local cache

		// if its new, append it.
		newPosts = append(newPosts, item)
	}

	return newPosts, nil
}

func (yc *YoutubeClient) CheckUriCache(uri *string) bool {
	for _, item := range YoutubeUriCache {
		if item == uri {
			return true
		}
	}

	return false
}

func (yc *YoutubeClient) ConvertToArticle(item *gofeed.Item) database.Article {
	parser, err := yc.GetParser(item.Link)
	if err != nil {
		log.Printf("[YouTube] Unable to process %v, submit this link as an issue.\n", item.Link)
	}

	tags, err := yc.GetTags(parser)
	if err != nil {
		msg := fmt.Sprintf("%v. %v", err, item.Link)
		log.Printf("[YouTube] %v", msg)
	}

	thumb, err := yc.GetVideoThumbnail(parser)
	if err != nil {
		msg := fmt.Sprintf("%v. %v", err, item.Link)
		log.Printf("[YouTube] %v", msg)
	}

	var article = database.Article{
		Sourceid:    yc.record.ID,
		Tags:        tags,
		Title:       item.Title,
		Url:         item.Link,
		Pubdate:     *item.PublishedParsed,
		Thumbnail:   thumb,
		Description: item.Description,
		Authorname:  sql.NullString{String: item.Author.Name},
		Authorimage: sql.NullString{String: yc.avatarUri},
	}
	return article
}
