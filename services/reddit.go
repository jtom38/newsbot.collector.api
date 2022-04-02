package services

import (
	"encoding/json"
	"fmt"
	"log"
	"errors"

	"github.com/jtom38/newsbot/collector/domain/model"
)

type RedditClient struct {
	subreddit string
	url string
	sourceId int32
}

var (
	PULLTOP string
	PULLHOT string
	PULLNSFW string
)

func init() {
	cc := NewConfigClient()
	PULLTOP = cc.GetConfig(REDDIT_PULL_TOP)
	PULLHOT = cc.GetConfig(REDDIT_PULL_HOT)
	PULLNSFW = cc.GetConfig(REDDIT_PULL_NSFW)
}

func NewReddit(subreddit string, sourceID int32) RedditClient {
	rc := RedditClient{
		subreddit: subreddit,
		url: fmt.Sprintf("https://www.reddit.com/r/%v.json", subreddit),
		sourceId:  sourceID,
	}
	return rc
}

// GetContent() reaches out to Reddit and pulls the Json data.
// It will then convert the data to a struct and return the struct.
func (rc RedditClient) GetContent() (model.RedditJsonContent, error ) {
	var items model.RedditJsonContent = model.RedditJsonContent{}

	log.Printf("Collecting results on '%v'", rc.subreddit)
	content, err := getHttpContent(rc.url)
	if err != nil { return items, err }

	json.Unmarshal(content, &items)
	
	return items, nil
}

// ConvertToArticle() will take the reddit model struct and convert them over to Article structs.
// This data can be passed to the database.
func (rc RedditClient) ConvertToArticle(source model.RedditPost) (model.Articles, error) {
	var item model.Articles


	if source.Content == "" && source.Url != ""{
		item = rc.convertPicturePost(source)
	}

	if item.Description == "" {
		var err = errors.New("reddit post failed to parse correctly")
		return item, err
	}

	return item, nil
}

func (rc RedditClient) convertPicturePost(source model.RedditPost) model.Articles {
	var item = model.Articles{
		SourceId: rc.sourceId,
		Url: fmt.Sprintf("https://www.reddit.com/%v", source.Permalink),
		Title: source.Title,
		AuthorName: source.Author,
		Description: source.Content,
		
	}
	return item
}

func (rc RedditClient) isTextPost(source model.RedditPost) {

}

func (rc RedditClient) isVideoPost(source model.RedditPost) {

}