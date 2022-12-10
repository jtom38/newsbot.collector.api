package input

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/models"
	"github.com/jtom38/newsbot/collector/services/config"
)

type RedditClient struct {
	config RedditConfig
	record database.Source
}

type RedditConfig struct {
	PullTop  string
	PullHot  string
	PullNSFW string
}

func NewRedditClient(Record database.Source) *RedditClient {
	rc := RedditClient{
		record: Record,
	}
	cc := config.New()
	rc.config.PullHot = cc.GetConfig(config.REDDIT_PULL_HOT)
	rc.config.PullNSFW = cc.GetConfig(config.REDDIT_PULL_NSFW)
	rc.config.PullTop = cc.GetConfig(config.REDDIT_PULL_TOP)

	//rc.disableHttp2Client()

	return &rc
}

// This is needed for to get modern go to talk to the endpoint.
// https://www.reddit.com/r/redditdev/comments/t8e8hc/getting_nothing_but_429_responses_when_using_go/
//func (rc *RedditClient) disableHttp2Client() {
//	os.Setenv("GODEBUG", "http2client=0")
//}

func (rc *RedditClient) GetBrowser() *rod.Browser {
	var browser *rod.Browser
	if path, exists := launcher.LookPath(); exists {
		u := launcher.New().Bin(path).MustLaunch()
		browser = rod.New().ControlURL(u).MustConnect()
	}
	return browser
}

func (rc *RedditClient) GetPage(parser *rod.Browser, url string) *rod.Page {
	page := parser.MustPage(url)
	return page
}

//func (rc RedditClient)

// GetContent() reaches out to Reddit and pulls the Json data.
// It will then convert the data to a struct and return the struct.
func (rc *RedditClient) GetContent() (models.RedditJsonContent, error) {
	var items models.RedditJsonContent = models.RedditJsonContent{}

	// TODO Wire this to support the config options
	Url := fmt.Sprintf("%v.json", rc.record.Url)

	log.Printf("[Reddit] Collecting results on '%v'", rc.record.Name)

	content, err := getHttpContent(Url)
	if err != nil {
		return items, err
	}
	if strings.Contains("<h1>whoa there, pardner!</h1>", string(content)) {
		return items, errors.New("did not get json data from the server")
	}

	json.Unmarshal(content, &items)
	if len(items.Data.Children) == 0 {
		return items, errors.New("failed to unmarshal the data")
	}
	return items, nil
}

func (rc *RedditClient) ConvertToArticles(items models.RedditJsonContent) []database.Article {
	var redditArticles []database.Article
	for _, item := range items.Data.Children {
		var article database.Article
		article, err := rc.convertToArticle(item.Data)
		if err != nil {
			log.Printf("[Reddit] %v", err)
			continue
		}
		redditArticles = append(redditArticles, article)
	}
	return redditArticles
}

// ConvertToArticle() will take the reddit model struct and convert them over to Article structs.
// This data can be passed to the database.
func (rc *RedditClient) convertToArticle(source models.RedditPost) (database.Article, error) {
	var item database.Article

	if source.Content == "" && source.Url != "" {
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

	if item.Description == "" && item.Title == "" {
		var err = errors.New("post failed to parse correctly")
		return item, err
	}

	return item, nil
}

func (rc *RedditClient) convertPicturePost(source models.RedditPost) database.Article {
	var item = database.Article{
		Sourceid:    rc.record.ID,
		Title:       source.Title,
		Tags:        fmt.Sprintf("%v", rc.record.Tags),
		Url:         fmt.Sprintf("https://www.reddit.com%v", source.Permalink),
		Pubdate:     time.Now(),
		Video:       sql.NullString{String: "null"},
		Videoheight: 0,
		Videowidth:  0,
		Thumbnail:   source.Thumbnail,
		Description: source.Content,
		Authorname:  sql.NullString{String: source.Author},
		Authorimage: sql.NullString{String: "null"},
	}
	return item
}

func (rc *RedditClient) convertTextPost(source models.RedditPost) database.Article {
	var item = database.Article{
		Sourceid:    rc.record.ID,
		Tags:        "a",
		Title:       source.Title,
		Pubdate:     time.Now(),
		Videoheight: 0,
		Videowidth:  0,
		Url:         fmt.Sprintf("https://www.reddit.com%v", source.Permalink),
		Authorname:  sql.NullString{String: source.Author},
		Description: source.Content,
	}
	return item
}

func (rc *RedditClient) convertVideoPost(source models.RedditPost) database.Article {
	var item = database.Article{
		Sourceid:    rc.record.ID,
		Tags:        "a",
		Title:       source.Title,
		Pubdate:     time.Now(),
		Url:         fmt.Sprintf("https://www.reddit.com%v", source.Permalink),
		Videoheight: 0,
		Videowidth:  0,
		Authorname:  sql.NullString{String: source.Author},
		Description: source.Media.RedditVideo.FallBackUrl,
	}
	return item
}

// This post is nothing more then a redirect to another location.
func (rc *RedditClient) convertRedirectPost(source models.RedditPost) database.Article {
	var item = database.Article{
		Sourceid:    rc.record.ID,
		Tags:        "a",
		Title:       source.Title,
		Pubdate:     time.Now(),
		Url:         fmt.Sprintf("https://www.reddit.com%v", source.Permalink),
		Videoheight: 0,
		Videowidth:  0,
		Authorname:  sql.NullString{String: source.Author},
		Description: source.UrlOverriddenByDest,
	}
	return item
}
