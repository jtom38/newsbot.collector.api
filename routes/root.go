package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RootRoutes() chi.Router {
	app := chi.NewRouter()
	app.Route("/", func(r chi.Router) {
		r.Get("/helloworld", helloWorld)
		r.Get("/ping", ping)
		r.Route("/hello/{who}", func(r chi.Router) {
			r.Get("/", helloWho)
		})
	})
	return app
}

// HelloWorld
// @Summary  Responds back with "Hello world!"
// @Produce  plain
// @Tags     debug
// @Router   /helloworld [get]
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// Ping
// @Summary  Sends back "pong".  Good to test with.
// @Produce  plain
// @Tags     debug
// @Router   /ping [get]
func ping(w http.ResponseWriter, r *http.Request) {
	msg := "pong"
	w.Write([]byte(msg))
}

// HelloWho
// @Summary  Responds back with "Hello x" depending on param passed in.
// @Param    who  path  string  true  "Who"
// @Produce  plain
// @Tags     debug
// @Router   /hello/{who} [get]
func helloWho(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Hello %v", chi.URLParam(r, "who"))
	w.Write([]byte(msg))
}
