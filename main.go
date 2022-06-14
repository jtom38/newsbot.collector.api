package main

import (
	"context"
	"fmt"
	"net/http"
	
	_ "github.com/jtom38/newsbot/collector/docs"
	"github.com/jtom38/newsbot/collector/routes"
	"github.com/jtom38/newsbot/collector/services/cron"
)



// @title     NewsBot collector
// @version   0.1
// @BasePath  /api
func main() {
	ctx := context.Background()
	c := cron.New(ctx)
	c.Start()

	server := routes.NewServer(ctx)

	fmt.Println("API is online and waiting for requests.")
	fmt.Println("API: http://localhost:8081/api")
	fmt.Println("Swagger: http://localhost:8081/swagger/index.html")
	err := http.ListenAndServe(":8081", server.Router)
	if err != nil { panic(err) }
}