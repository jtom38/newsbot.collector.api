package main

import (
	//"fmt"
	"log"
	"net/http"


	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jtom38/newsbot/collector/routes"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services"
)

func main() {
	var err error
	//EnableScheduler()
	//dc := database.NewDatabaseClient()
	//err := dc.Diagnosis.Ping()
	//if err != nil { log.Fatalln(err) }

	//CheckReddit()
	CheckYoutube()

	app := chi.NewRouter()
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	//app.Mount("/swagger", httpSwagger.WrapHandler)
	app.Mount("/api", routes.RootRoutes())
	
	log.Println("API is online and waiting for requests.")
	log.Println("API: http://localhost:8081/api")
	//log.Println("Swagger: http://localhost:8080/swagger/index.html")
	err = http.ListenAndServe(":8081", app)
	if err != nil { log.Fatalln(err) }
}

func CheckReddit() {
	dc := database.NewDatabaseClient()
	sources, err := dc.Sources.FindBySource("reddit")
	if err != nil { log.Println(err) }

	rc := services.NewRedditClient(sources[0].Name, sources[0].ID)
	raw, err := rc.GetContent()
	if err != nil { log.Println(err) }
	
	redditArticles := rc.ConvertToArticles(raw)
	
	for _, item := range redditArticles {		
		_, err = dc.Articles.FindByUrl(item.Url)
		if err != nil {
			err = dc.Articles.Add(item)
			if err != nil { log.Println("Failed to post article.")}
		}
	}
}

func CheckYoutube() {
	// Add call to the db to request youtube sources.

	// Loop though the services, and generate the clients.
	yt := services.NewYoutubeClient(0, "https://www.youtube.com/user/GameGrumps")
	yt.CheckSource()
}