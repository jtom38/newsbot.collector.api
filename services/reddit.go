package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/jtom38/newsbot/collector/domain/model"
	"github.com/jtom38/newsbot/collector/services/config"
)

type RedditClient struct {
	subreddit string
	url string
	sourceId uint
	config RedditConfig
}

type RedditConfig struct {
	PullTop string
	PullHot string
	PullNSFW string
}

func NewRedditClient(subreddit string, sourceID uint) RedditClient {
	rc := RedditClient{
		subreddit: subreddit,
		url: fmt.Sprintf("https://www.reddit.com/r/%v.json", subreddit),
		sourceId:  sourceID,
	}
	cc := config.New()
	rc.config.PullHot = cc.GetConfig(config.REDDIT_PULL_HOT)
	rc.config.PullNSFW = cc.GetConfig(config.REDDIT_PULL_NSFW)
	rc.config.PullTop = cc.GetConfig(config.REDDIT_PULL_TOP)

	rc.disableHttp2Client()

	return rc
}

// This is needed for to get modern go to talk to the endpoint.
// https://www.reddit.com/r/redditdev/comments/t8e8hc/getting_nothing_but_429_responses_when_using_go/
func (rc RedditClient) disableHttp2Client() {
	os.Setenv("GODEBUG", "http2client=0")
}

func (rc RedditClient) GetBrowser() *rod.Browser {
	browser := rod.New().MustConnect()
	return browser
}

func (rc RedditClient) GetPage(parser *rod.Browser, url string) *rod.Page {
	page := parser.MustPage(url)
	return page
}

// GetContent() reaches out to Reddit and pulls the Json data.
// It will then convert the data to a struct and return the struct.
func (rc RedditClient) GetContent() (model.RedditJsonContent, error ) {
	var items model.RedditJsonContent = model.RedditJsonContent{}

	log.Printf("Collecting results on '%v'", rc.subreddit)
	content, err := getHttpContent(rc.url)
	if err != nil { return items, err }
	if strings.Contains("<h1>whoa there, pardner!</h1>", string(content) ) {
		return items, errors.New("did not get json data from the server")
	}

	json.Unmarshal(content, &items)
	if len(items.Data.Children) == 0 {
		return items, errors.New("failed to unmarshal the data")
	}
	return items, nil
}

func (rc RedditClient) ConvertToArticles(items model.RedditJsonContent) []model.Articles {
	var redditArticles []model.Articles
	for _, item := range items.Data.Children {
		var article model.Articles
		article, err := rc.convertToArticle(item.Data)
		if err != nil { log.Println(err); continue }
		redditArticles = append(redditArticles, article)
	}
	return redditArticles
}

// ConvertToArticle() will take the reddit model struct and convert them over to Article structs.
// This data can be passed to the database.
func (rc RedditClient) convertToArticle(source model.RedditPost) (model.Articles, error) {
	var item model.Articles

	
	if source.Content == "" && source.Url != ""{
		item = rc.convertPicturePost(source)
	}
	
	if source.Media.RedditVideo.FallBackUrl != "" {
		item = rc.convertVideoPost(source)
	}

	if source.Content != "" {
		item = rc.convertTextPost(source)
	}

	if source.UrlOverriddenByDest != "" {
		item = rc.convertRedirectPost(source)
	}

	if item.Description == "" {
		var err = errors.New("reddit post failed to parse correctly")
		return item, err
	}

	return item, nil
}

func (rc RedditClient) convertPicturePost(source model.RedditPost) model.Articles {
	var item = model.Articles{
		SourceID: rc.sourceId,
		Tags: "a",
		Title: source.Title,
		Url: fmt.Sprintf("https://www.reddit.com%v", source.Permalink),
		PubDate: time.Now(),
		Video: "null",
		VideoHeight: 0,
		VideoWidth: 0,
		Thumbnail: source.Thumbnail,
		Description: source.Content,
		AuthorName: source.Author,
		AuthorImage: "null",
	}
	return item
}

func (rc RedditClient) convertTextPost(source model.RedditPost) model.Articles {
	var item = model.Articles{
		SourceID: rc.sourceId,
		Tags: "a",
		Title: source.Title,
		Url: fmt.Sprintf("https://www.reddit.com%v", source.Permalink),
		AuthorName: source.Author,
		Description: source.Content,
		
	}
	return item
}

func (rc RedditClient) convertVideoPost(source model.RedditPost) model.Articles {
	var item = model.Articles{
		SourceID: rc.sourceId,
		Tags: "a",
		Title: source.Title,
		Url: fmt.Sprintf("https://www.reddit.com%v", source.Permalink),
		AuthorName: source.Author,
		Description: source.Media.RedditVideo.FallBackUrl,
	}
	return item
}

// This post is nothing more then a redirect to another location.
func (rc *RedditClient) convertRedirectPost(source model.RedditPost) model.Articles {
	var item = model.Articles{
		SourceID: rc.sourceId,
		Tags: "a",
		Title: source.Title,
		Url: fmt.Sprintf("https://www.reddit.com%v", source.Permalink),
		AuthorName: source.Author,
		Description: source.UrlOverriddenByDest,
	}
	return item
}