package main

import (
	//"fmt"
	"log"
	"net/http"


	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jtom38/newsbot/collector/routes"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/model"
	"github.com/jtom38/newsbot/collector/services"
)

func main() {
	//EnableScheduler()
	dc := database.NewDatabaseClient()
	err := dc.Diagnosis.Ping()
	if err != nil { log.Fatalln(err) }

	CheckReddit()

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

	rc := services.NewReddit(sources[0].Name, sources[0].ID)
	raw, err := rc.GetContent()
	if err != nil { log.Println(err) }
	
	var redditArticles []model.Articles
	for _, item := range raw.Data.Children {
		var article model.Articles
		article, err = rc.ConvertToArticle(item.Data)
		if err != nil { log.Println(err); continue }
		redditArticles = append(redditArticles, article)
	}

	for _, item := range redditArticles {
		dc.Articles.FindByUrl(item.Url)
	}
	dc.Articles.Add()

}