package cron

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"

	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services/config"
	"github.com/jtom38/newsbot/collector/services/input"
	"github.com/jtom38/newsbot/collector/services/output"
)

type Cron struct {
	Db    *database.Queries
	ctx   *context.Context
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
		ctx: &ctx,
	}

	timer := cron.New()
	queries, err := openDatabase()
	if err != nil {
		panic(err)
	}
	c.Db = queries

	//timer.AddFunc("*/5 * * * *", func()  { go CheckCache() })
	features := config.New()

	res, _ := features.GetFeature(config.FEATURE_ENABLE_REDDIT_BACKEND)
	if res {
		timer.AddFunc("*/5 * * * *", func() { go c.CheckReddit() })
		log.Print("[Input] Reddit backend was enabled")
		//go c.CheckReddit()
	}

	res, _ = features.GetFeature(config.FEATURE_ENABLE_YOUTUBE_BACKEND)
	if res {
		timer.AddFunc("*/5 * * * *", func() { go c.CheckYoutube() })
		log.Print("[Input] YouTube backend was enabled")
	}

	res, _ = features.GetFeature(config.FEATURE_ENABLE_FFXIV_BACKEND)
	if res {
		timer.AddFunc("* */1 * * *", func() { go c.CheckFfxiv() })
		log.Print("[Input] FFXIV backend was enabled")
	}

	res, _ = features.GetFeature(config.FEATURE_ENABLE_TWITCH_BACKEND)
	if res {
		timer.AddFunc("* */1 * * *", func() { go c.CheckTwitch() })
		log.Print("[Input] Twitch backend was enabled")
	}

	timer.AddFunc("*/5 * * * *", func() { go c.CheckDiscordQueue() })
	log.Print("[Output] Discord Output was enabled")

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
func (c *Cron) CheckReddit() {
	sources, err := c.Db.ListSourcesBySource(*c.ctx, "reddit")
	if err != nil {
		log.Printf("[Reddit] No sources found to query - %v\r", err)
	}

	for _, source := range sources {
		if !source.Enabled {
			continue
		}
		log.Printf("[Reddit] Checking '%v'...", source.Name)
		rc := input.NewRedditClient(source)
		raw, err := rc.GetContent()
		if err != nil {
			log.Println(err)
		}
		redditArticles := rc.ConvertToArticles(raw)
		c.checkPosts(redditArticles, "Reddit")
	}
	log.Print("[Reddit] Done!")
}

func (c *Cron) CheckYoutube() {
	// Add call to the db to request youtube sources.
	sources, err := c.Db.ListSourcesBySource(*c.ctx, "youtube")
	if err != nil {
		log.Printf("[Youtube] No sources found to query - %v\r", err)
	}

	for _, source := range sources {
		if !source.Enabled {
			continue
		}
		log.Printf("[YouTube] Checking '%v'...", source.Name)
		yc := input.NewYoutubeClient(source)
		raw, err := yc.GetContent()
		if err != nil {
			log.Println(err)
		}
		c.checkPosts(raw, "YouTube")
	}
	log.Print("[YouTube] Done!")
}

func (c *Cron) CheckFfxiv() {
	sources, err := c.Db.ListSourcesBySource(*c.ctx, "ffxiv")
	if err != nil {
		log.Printf("[FFXIV] No sources found to query - %v\r", err)
	}

	for _, source := range sources {
		if !source.Enabled {
			continue
		}
		fc := input.NewFFXIVClient(source)
		items, err := fc.CheckSource()
		if err != nil {
			log.Println(err)
		}
		c.checkPosts(items, "FFXIV")
	}
	log.Printf("[FFXIV Done!]")
}

func (c *Cron) CheckTwitch() error {
	sources, err := c.Db.ListSourcesBySource(*c.ctx, "twitch")
	if err != nil {
		log.Printf("[Twitch] No sources found to query - %v\r", err)
	}

	tc, err := input.NewTwitchClient()
	if err != nil {
		return err
	}

	for _, source := range sources {
		if !source.Enabled {
			continue
		}
		log.Printf("[Twitch] Checking '%v'...", source.Name)
		tc.ReplaceSourceRecord(source)
		items, err := tc.GetContent()
		if err != nil {
			log.Println(err)
		}
		c.checkPosts(items, "Twitch")
	}

	log.Print("[Twitch] Done!")
	return nil
}

func (c *Cron) CheckDiscordQueue() error {
	// Get items from the table
	queueItems, err := c.Db.ListDiscordQueueItems(*c.ctx, 50)
	if err != nil {
		return err
	}

	for _, queue := range queueItems {
		// Get the articleByID
		article, err := c.Db.GetArticleByID(*c.ctx, queue.Articleid)
		if err != nil {
			return err
		}

		var endpoints []string
		// List Subscription by SourceID
		subs, err := c.Db.ListSubscriptionsBySourceId(*c.ctx, article.Sourceid)
		if err != nil {
			return err
		}

		// if no one is subscribed to it, remove it from the index.
		if len(subs) == 0 {
			log.Printf("No subscriptions found bound to '%v' so it was removed.", article.Sourceid)
			err = c.Db.DeleteDiscordQueueItem(*c.ctx, queue.ID)
			if err != nil {
				return err
			}
			continue
		}

		// Get the webhhooks to send to
		for _, sub := range subs {
			webhook, err := c.Db.GetDiscordWebHooksByID(*c.ctx, sub.Discordwebhookid)
			if err != nil {
				return err
			}

			// store them in an array
			endpoints = append(endpoints, webhook.Url)
		}

		// Create Discord Message
		dwh := output.NewDiscordWebHookMessage(article)
		msg, err := dwh.GeneratePayload()
		if err != nil {
			return err
		}

		// Send Message(s)
		for _, i := range endpoints {
			err = dwh.SendPayload(msg, i)

			if err != nil {
				return err
			}
		}

		// Remove the item from the queue, given we sent our notification.
		err = c.Db.DeleteDiscordQueueItem(*c.ctx, queue.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Cron) checkPosts(posts []database.Article, sourceName string) error {
	for _, item := range posts {
		_, err := c.Db.GetArticleByUrl(*c.ctx, item.Url)
		if err != nil {
			id := uuid.New()

			err := c.postArticle(id, item)
			if err != nil {
				return fmt.Errorf("[%v] Failed to post article - %v - %v.\r", sourceName, item.Url, err)
			}

			err = c.addToDiscordQueue(id)
			if err != nil {
				return err
			}

		}
	}
	time.Sleep(30 * time.Second)
	return nil
}

func (c *Cron) postArticle(id uuid.UUID,item database.Article) error {
	err := c.Db.CreateArticle(*c.ctx, database.CreateArticleParams{
		ID:          id,
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

func (c *Cron) addToDiscordQueue(Id uuid.UUID) error {
	err := c.Db.CreateDiscordQueue(*c.ctx, database.CreateDiscordQueueParams{
		ID: uuid.New(),
		Articleid: Id,
	})
	if err != nil {
		return err
	}
	return nil
}
