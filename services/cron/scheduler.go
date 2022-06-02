package cron

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"

	//"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/model"
	"github.com/jtom38/newsbot/collector/services"
	"github.com/jtom38/newsbot/collector/services/config"
	//"github.com/jtom38/newsbot/collector/services/cache"
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

func CheckReddit(ctx context.Context) {
	sources, err := _queries.GetAllSources(ctx)
	//sources, err := _queries.GetSourcesBySource(ctx, sql.NullString{String: "reddit"})
	if err != nil {
		panic(err)
	}

	for _, source := range sources {
		rc := services.NewRedditClient(source.Name.String, source.ID)

		raw, err := rc.GetContent()
		if err != nil {
			log.Println(err)
		}

		redditArticles := rc.ConvertToArticles(raw)

		for _, item := range redditArticles {
			_, err := _queries.GetArticleByUrl(ctx, item.Url)
			if err != nil {
				err = postArticle(ctx, item)
				if err != nil {
					log.Println("Failed to post article.")
				}
			}
		}
		time.Sleep(30 * time.Second)
	}
}

func CheckYoutube() {
	// Add call to the db to request youtube sources.

	// Loop though the services, and generate the clients.
	yt := services.NewYoutubeClient(0, "https://www.youtube.com/user/GameGrumps")
	yt.CheckSource()
}

func CheckFfxiv() {
	fc := services.NewFFXIVClient("na")
	_, err := fc.CheckSource()

	// This isnt in a thread yet, so just output to stdout
	if err != nil {
		log.Println(err)
	}

	/*
		dc := database.NewDatabaseClient()
		for _, item := range articles {
			_, err = dc.Articles.FindByUrl(item.Url)
			if err != nil {
				err = dc.Articles.Add(item)
				if err != nil { log.Println("Failed to post article.")}
			}
		}
	*/
}

func CheckTwitch() error {
	// TODO Wire this for the DB
	// just a mock object for now
	//dc := database.NewDatabaseClient()

	//sources, err := dc.Sources.FindBySource("Twitch")
	//if err != nil { return err }

	source := model.Sources{
		ID:   1,
		Name: "Nintendo",
	}
	client, err := services.NewTwitchClient(source)
	if err != nil {
		log.Println(err)
	}

	err = client.Login()
	if err != nil {
		return err
	}

	//for _, source := range sources {
	client.ReplaceSourceRecord(source)

	_, err = client.GetContent()
	if err != nil {
		return err
	}
	/*
		for _, item := range posts {
			_, err = dc.Articles.FindByUrl(item.Url)
			if err != nil {
				err = dc.Articles.Add(item)
				if err != nil { log.Println("Failed to post article.")}
			}
		}
	*/
	//}

	return nil
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
