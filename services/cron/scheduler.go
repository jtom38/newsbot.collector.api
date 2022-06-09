package cron

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"

	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services"
	"github.com/jtom38/newsbot/collector/services/config"
)

var _env config.ConfigClient
var _connString string
var _queries *database.Queries

func EnableScheduler(ctx context.Context) {
	c := cron.New()
	OpenDatabase(ctx)

	//c.AddFunc("*/5 * * * *", func()  { go CheckCache() })
	c.AddFunc("* */1 * * *", func() { go CheckReddit(ctx) })
	//c.AddFunc("* */1 * * *", func() { go CheckYoutube() })
	//c.AddFunc("* */1 * * *", func() { go CheckFfxiv() })
	//c.AddFunc("* */1 * * *", func() { go CheckTwitch() })

	c.Start()
}

// Open the connection to the database and share it with the package so all of them are able to share.
func OpenDatabase(ctx context.Context) error {
	_env = config.New()
	_connString = _env.GetConfig(config.Sql_Connection_String)
	db, err := sql.Open("postgres", _connString)
	if err != nil {
		panic(err)
	}

	queries := database.New(db)
	_queries = queries
	return err
}

// This is the main entry point to query all the reddit services
func CheckReddit(ctx context.Context) {
	sources, err := _queries.ListSourcesBySource(ctx, "reddit")
	if err != nil {
		log.Printf("No defines sources for reddit to query - %v\r", err)
	}

	for _, source := range sources {
		if !source.Enabled {
			continue
		}
		rc := services.NewRedditClient(source)
		raw, err := rc.GetContent()
		if err != nil {
			log.Println(err)
		}
		redditArticles := rc.ConvertToArticles(raw)
		checkPosts(ctx, redditArticles)
	}
}

func CheckYoutube(ctx context.Context) {
	// Add call to the db to request youtube sources.
	sources, err := _queries.ListSourcesBySource(ctx, "youtube")
	if err != nil {
		log.Printf("Youtube - No sources found to query - %v\r", err)
	}

	for _, source := range sources {
		if !source.Enabled {
			continue
		}
		yc := services.NewYoutubeClient(source)
		raw, err := yc.GetContent()
		if err != nil {
			log.Println(err)
		}
		checkPosts(ctx, raw)
	}
}

func CheckFfxiv(ctx context.Context) {
	sources, err := _queries.ListSourcesBySource(ctx, "ffxiv")
	if err != nil {
		log.Printf("Final Fantasy XIV - No sources found to query - %v\r", err)
	}

	for _, source := range sources {
		if !source.Enabled {
			continue
		}
		fc := services.NewFFXIVClient(source)
		items, err := fc.CheckSource()
		if err != nil {
			log.Println(err)
		}
		checkPosts(ctx, items)
	}
}

func CheckTwitch(ctx context.Context) error {
	sources, err := _queries.ListSourcesBySource(ctx, "twitch")
	if err != nil {
		log.Printf("Twitch - No sources found to query - %v\r", err)
	}
	
	tc, err := services.NewTwitchClient()
	if err != nil {
		return err
	}

	for _, source := range sources {
		if !source.Enabled {
			continue
		}
		tc.ReplaceSourceRecord(source)
		items, err := tc.GetContent()
		if err != nil {
			log.Println(err)
		}
		checkPosts(ctx, items)
	}

	return nil
}

func checkPosts(ctx context.Context, posts []database.Article) {
	for _, item := range posts {
		_, err := _queries.GetArticleByUrl(ctx, item.Url)
		if err != nil {
			err = postArticle(ctx, item)
			if err != nil {
				log.Printf("Reddit - Failed to post article - %v - %v.\r", item.Url, err)
			} else {
				log.Printf("Reddit - Posted article - %v\r", item.Url)
			}
		}
	}
	time.Sleep(30 * time.Second)
}

func postArticle(ctx context.Context, item database.Article) error {
	err := _queries.CreateArticle(ctx, database.CreateArticleParams{
		ID:          uuid.New(),
		Sourceid:    item.Sourceid,
		Tags:        item.Tags,
		Title:       item.Title,
		Url:         item.Url,
		Pubdate:     item.Pubdate,
		Video:       item.Video,
		Videoheight: item.Videoheight,
		Videowidth:  item.Videowidth,
		Thumbnail:   item.Thumbnail,
		Description: item.Description,
		Authorname:  item.Authorname,
		Authorimage: item.Authorimage,
	})
	return err
}
