package routes

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi/v5"
)

func RootRoutes() chi.Router {
	app := chi.NewRouter()
	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	app.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		msg := "pong"
		w.Write([]byte(msg))
	})

	app.Get("/hello/{world}", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Hello %v", chi.URLParam(r, "world"))
		w.Write([]byte(msg))
	})
	return app
}