package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/docs"
	"github.com/jtom38/newsbot/collector/routes"
	"github.com/jtom38/newsbot/collector/services/config"
	"github.com/jtom38/newsbot/collector/services/cron"
)

// @title     NewsBot collector
// @version   0.1
// @BasePath  /api
func main() {
	cfg := config.New()
	address := cfg.GetConfig(config.ServerAddress)
	docs.SwaggerInfo.Host = fmt.Sprintf("%v:8081", address)

	ctx := context.Background()
	db, err := sql.Open("postgres", cfg.GetConfig(config.Sql_Connection_String))
	if err != nil {
		panic(err)
	}

	queries := database.New(db)

	c := cron.New(ctx)
	c.Start()

	server := routes.NewServer(ctx, queries)

	fmt.Println("API is online and waiting for requests.")
	fmt.Printf("API: http://%v:8081/api\r\n", address)
	fmt.Printf("Swagger: http://%v:8081/swagger/index.html\r\n", address)

	err = http.ListenAndServe(":8081", server.Router)
	if err != nil {
		panic(err)
	}
}
