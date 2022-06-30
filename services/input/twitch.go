package input

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services/config"
	"github.com/nicklaw5/helix/v2"
)

type TwitchClient struct {
	SourceRecord database.Source

	// config
	monitorClips string
	monitorVod   string

	// API Connection
	api *helix.Client
}

var (
	ErrTwitchClientIdMissing     = errors.New("unable to find the client id for auth")
	ErrTwitchClientSecretMissing = errors.New("unable to find the client secret for auth")

	twitchScopes = "user:read:email"
)

func NewTwitchClient() (TwitchClient, error) {
	c := config.New()

	id := c.GetConfig(config.TWITCH_CLIENT_ID)
	if id == "" {
		return TwitchClient{}, ErrTwitchClientIdMissing
	}

	secret := c.GetConfig(config.TWITCH_CLIENT_SECRET)
	if secret == "" {
		return TwitchClient{}, ErrTwitchClientSecretMissing
	}

	api, err := initTwitchApi(id, secret)
	if err != nil {
		return TwitchClient{}, nil
	}

	client := TwitchClient{
		//SourceRecord: &source,
		monitorClips: c.GetConfig(config.TWITCH_MONITOR_CLIPS),
		monitorVod:   c.GetConfig(config.TWITCH_MONITOR_VOD),
		api:          &api,
	}

	return client, nil
}

// Sets up the API connection to Twitch.
func initTwitchApi(ClientId string, ClientSecret string) (helix.Client, error) {
	api, err := helix.NewClient(&helix.Options{
		ClientID:     ClientId,
		ClientSecret: ClientSecret,
	})
	if err != nil {
		return helix.Client{}, err
	}

	return *api, nil
}

// This will let you replace the bound source record to keep the same session alive.
func (tc *TwitchClient) ReplaceSourceRecord(source database.Source) {
	tc.SourceRecord = source
}

// Invokes Logon request to the API
func (tc TwitchClient) Login() error {
	token, err := tc.api.RequestAppAccessToken([]string{twitchScopes})
	if err != nil {
		return err
	}

	tc.api.SetAppAccessToken(token.Data.AccessToken)
	return nil
}

func (tc TwitchClient) GetContent() ([]database.Article, error) {
	var items []database.Article

	user, err := tc.GetUserDetails()
	if err != nil {
		return items, err
	}

	posts, err := tc.GetPosts(user)
	if err != nil {
		return items, err
	}

	for _, video := range posts {
		var article database.Article

		AuthorName, err := tc.ExtractAuthor(video)
		if err != nil { return items, err }
		article.Authorname = sql.NullString{String: AuthorName}
		
		Authorimage, err := tc.ExtractAuthorImage(user)
		if err != nil { return items, err }
		article.Authorimage = sql.NullString{String: Authorimage}

		article.Description, err = tc.ExtractDescription(video)
		if err != nil {return items, err }

		article.Pubdate, err = tc.ExtractPubDate(video)
		if err != nil { return items, err }

		article.Sourceid = tc.SourceRecord.ID
		article.Tags, err = tc.ExtractTags(video, user)
		if err != nil { return items, err }

		article.Thumbnail, err = tc.ExtractThumbnail(video)
		if err != nil { return items, err }

		article.Title, err = tc.ExtractTitle(video)
		if err != nil { return items, err }
		
		article.Url, err = tc.ExtractUrl(video)
		if err != nil { return items, err }

		items = append(items, article)
	}

	return items, nil
}

func (tc TwitchClient) GetUserDetails() (helix.User, error) {
	var blank helix.User

	users, err := tc.api.GetUsers(&helix.UsersParams{
		Logins: []string{tc.SourceRecord.Name},
	})
	if err != nil {
		return blank, err
	}
	return users.Data.Users[0], nil
}

// This will reach out and collect the posts made by the user.
func (tc TwitchClient) GetPosts(user helix.User) ([]helix.Video, error) {
	var blank []helix.Video

	videos, err := tc.api.GetVideos(&helix.VideosParams{
		UserID: user.ID,
	})
	if err != nil {
		return blank, err
	}

	//log.Println(videos.Data.Videos)
	return videos.Data.Videos, nil
}

func (tc TwitchClient) ExtractAuthor(post helix.Video) (string, error) {
	if post.UserName == "" {
		return "", ErrMissingAuthorName
	}
	return post.UserName, nil
}

func (tc TwitchClient) ExtractThumbnail(post helix.Video) (string, error) {
	if post.ThumbnailURL == "" {
		return "", ErrMissingThumbnail
	}
	var thumb = post.ThumbnailURL
	thumb = strings.Replace(thumb, "%{width}", "600", -1)
	thumb = strings.Replace(thumb, "%{height}", "400", -1)
	return thumb, nil
}

func (tc TwitchClient) ExtractPubDate(post helix.Video) (time.Time, error) {
	if post.PublishedAt == "" {
		return time.Now(), ErrMissingPublishDate
	}
	pubDate, err := time.Parse("2006-01-02T15:04:05Z", post.PublishedAt)
	if err != nil {
		return time.Now(), err
	}
	return pubDate, nil
}

func (tc TwitchClient) ExtractDescription(post helix.Video) (string, error) {
	// Check if the description is null but we have a title.
	// The poster didnt add a description but this isnt an error.
	if post.Description == "" && post.Title == "" {
		return "", ErrMissingDescription
	}
	if post.Description == "" {
		return "No description was given", nil
	}
	return post.Description, nil
}

// Extracts the avatar of the author with some validation.
func (tc TwitchClient) ExtractAuthorImage(user helix.User) (string, error) {
	if user.ProfileImageURL == "" { return "", ErrMissingAuthorImage }
	if !strings.Contains(user.ProfileImageURL, "-profile_image-") { return "", ErrInvalidAuthorImage }
	return user.ProfileImageURL, nil
}

// Generate tags based on the video metadata.
// TODO Figure out how to query what game is played
func (tc TwitchClient) ExtractTags(post helix.Video, user helix.User) (string, error) {
	res := fmt.Sprintf("twitch,%v,%v", post.Title, user.DisplayName)
	return res, nil
}

// Extracts the title from a post with some validation.
func (tc TwitchClient) ExtractTitle(post helix.Video) (string, error) {
	if post.Title == "" {
		return "", errors.New("unable to find the title on the requested post")
	}
	return post.Title, nil
}

func (tc TwitchClient) ExtractUrl(post helix.Video) (string, error) {
	if post.URL == "" { return "", ErrMissingUrl }
	return post.URL, nil
}