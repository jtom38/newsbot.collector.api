package routes

import "github.com/go-chi/chi/v5"

func RedditRouter() chi.Router {
	app := chi.NewRouter()
	app.Route("/", func(r chi.Router) {
		//r.Get("/")
	})
	
	return app
}

//func listSources