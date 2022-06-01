package cron

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/robfig/cron/v3"

	//"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/model"
	"github.com/jtom38/newsbot/collector/services"
	"github.com/jtom38/newsbot/collector/services/config"
	//"github.com/jtom38/newsbot/collector/services/cache"
)

func EnableScheduler() {
	c := cron.New()

	//c.AddFunc("*/5 * * * *", func()  { go CheckCache() })	
	c.AddFunc("* */1 * * *", func() { go CheckReddit() })
	//c.AddFunc("* */1 * * *", func() { go CheckYoutube() })
	//c.AddFunc("* */1 * * *", func() { go CheckFfxiv() })
	//c.AddFunc("* */1 * * *", func() { go CheckTwitch() })

	c.Start()
}

var ctx context.Context = context.Background()

func CheckReddit() {
	env := config.New()
	connString := env.GetConfig(config.Sql_Connection_String)

	db, err := sql.Open("postgres", connString)
	if err != nil { panic(err) }

	queries := database.New(db)
	sources, err  := queries.GetSourcesBySource(sql.NullString{String: "reddit"})
	if err != nil { panic(err) }

	for _, source := range sources {
		rc := services.NewRedditClient(source.Name.String, source.ID )

		
			raw, err := rc.GetContent()
			if err != nil { log.Println(err) }
			
			redditArticles := rc.ConvertToArticles(raw)
			
			for _, item := range redditArticles {
				_, err := queries.GetArticleByUrl(ctx, item.Url)	
				if err != nil {
					queries.CreateArticle(ctx, database.CreateArticleParams{
		
					})
					//err = dc.Articles.Add(item)
					if err != nil { log.Println("Failed to post article.")}
				}
			}
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
	if err != nil { log.Println(err) }
	
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
		ID: 1,
		Name: "Nintendo",
	}
	client, err := services.NewTwitchClient(source)
	if err != nil { log.Println(err) }

	err = client.Login()
	if err != nil { return err }

	//for _, source := range sources {
		client.ReplaceSourceRecord(source)
	
		_, err = client.GetContent()
		if err != nil { return err }
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