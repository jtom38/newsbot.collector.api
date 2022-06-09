package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/jtom38/newsbot/collector/routes"
	"github.com/jtom38/newsbot/collector/services/cron"
)

func main() {
	ctx := context.Background()

	cron.EnableScheduler(ctx)

	app := chi.NewRouter()
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	//app.Mount("/swagger", httpSwagger.WrapHandler)
	app.Mount("/api", routes.RootRoutes())
	
	log.Println("API is online and waiting for requests.")
	log.Println("API: http://localhost:8081/api")
	//log.Println("Swagger: http://localhost:8080/swagger/index.html")
	err := http.ListenAndServe(":8081", app)
	if err != nil { log.Fatalln(err) }
}