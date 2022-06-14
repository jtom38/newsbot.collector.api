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

type Cron struct {
	Db *database.Queries
	ctx *context.Context
	timer *cron.Cron
}

func openDatabase() (*database.Queries, error) {
	_env := config.New()
	connString := _env.GetConfig(config.Sql_Connection_String)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	queries := database.New(db)
	return queries, err
}

func New(ctx context.Context) *Cron {
	c := &Cron{
		ctx:  &ctx,
	}

	timer := cron.New()
	queries, err := openDatabase()
	if err != nil {
		panic(err)
	}
	c.Db = queries

	//c.AddFunc("*/5 * * * *", func()  { go CheckCache() })
	//c.AddFunc("* */1 * * *", func() { go CheckReddit(ctx) })
	//c.AddFunc("* */1 * * *", func() { go CheckYoutube() })
	//c.AddFunc("* */1 * * *", func() { go CheckFfxiv() })
	//c.AddFunc("* */1 * * *", func() { go CheckTwitch() })
	c.timer = timer
	return c
}

func (c *Cron) Start() {
	c.timer.Start()
}

func (c *Cron) Stop() {
	c.timer.Stop()
}

// This is the main entry point to query all the reddit services
func (c *Cron) CheckReddit(ctx context.Context) {
	sources, err := c.Db.ListSourcesBySource(*c.ctx, "reddit")
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
		c.checkPosts(*c.ctx, redditArticles)
	}
}

func (c *Cron) CheckYoutube(ctx context.Context) {
	// Add call to the db to request youtube sources.
	sources, err := c.Db.ListSourcesBySource(*c.ctx, "youtube")
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
		c.checkPosts(*c.ctx, raw)
	}
}

func (c *Cron) CheckFfxiv(ctx context.Context) {
	sources, err := c.Db.ListSourcesBySource(*c.ctx, "ffxiv")
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
		c.checkPosts(*c.ctx, items)
	}
}

func (c *Cron) CheckTwitch(ctx context.Context) error {
	sources, err := c.Db.ListSourcesBySource(*c.ctx, "twitch")
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
		c.checkPosts(*c.ctx, items)
	}

	return nil
}

func (c *Cron) checkPosts(ctx context.Context, posts []database.Article) {
	for _, item := range posts {
		_, err := c.Db.GetArticleByUrl(*c.ctx, item.Url)
		if err != nil {
			err = c.postArticle(ctx, item)
			if err != nil {
				log.Printf("Reddit - Failed to post article - %v - %v.\r", item.Url, err)
			} else {
				log.Printf("Reddit - Posted article - %v\r", item.Url)
			}
		}
	}
	time.Sleep(30 * time.Second)
}

func (c *Cron) postArticle(ctx context.Context, item database.Article) error {
	err := c.Db.CreateArticle(*c.ctx, database.CreateArticleParams{
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
