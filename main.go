package main

import (
	"context"
	"fmt"
	"net/http"

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
	c := cron.New(ctx)
	c.Start()

	server := routes.NewServer(ctx)

	fmt.Println("API is online and waiting for requests.")
	fmt.Printf("API: http://%v:8081/api\r\n", address)
	fmt.Printf("Swagger: http://%v:8081/swagger/index.html\r\n", address)

	err := http.ListenAndServe(":8081", server.Router)
	if err != nil {
		panic(err)
	}
}
