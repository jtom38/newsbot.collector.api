package routes

import (
	"context"
	"database/sql"
	//"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/lib/pq"

	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services/config"
)

type Server struct {
	Router *chi.Mux
	Db *database.Queries
	ctx *context.Context
}

func NewServer(ctx context.Context) *Server {
	s := &Server{
		ctx: &ctx,
	}

	db, err := openDatabase(ctx)
	if err != nil {
		panic(err)
	}
	s.Db = db

	s.Router = chi.NewRouter()
	s.MountMiddleware()
	s.MountRoutes()
	return s
}

func openDatabase(ctx context.Context) (*database.Queries, error) {
	_env := config.New()
	connString := _env.GetConfig(config.Sql_Connection_String)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	queries := database.New(db)
	return queries, err
}

func (s *Server) MountMiddleware() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
}

func (s *Server) MountRoutes() {
	s.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"), //The url pointing to API definition
	))
	
	/* Root Routes */
	s.Router.Get("/api/helloworld", helloWorld)
	s.Router.Get("/api/hello/{who}", helloWho)
	s.Router.Get("/api/ping", ping)

	/* Article Routes */
	s.Router.Get("/api/config/articles", s.listArticles)
	s.Router.Route("/api/config/articles/{ID}", func(r chi.Router) {
		r.Get("/", s.getArticleById)
		
	})

	/* Source Routes */
	s.Router.Get("/api/config/sources", s.listSources)
	s.Router.Post("/api/config/sources/new/reddit", s.newRedditSource)
	s.Router.Post("/api/config/sources/new/youtube", s.newYoutubeSource)
	s.Router.Post("/api/config/sources/new/twitch", s.newTwitchSource)
	s.Router.Route("/api/config/sources/{ID}", func(r chi.Router) {
		r.Get("/", s.getSources)
		r.Delete("/", s.deleteSources)
		r.Post("/disable", s.disableSource)
		r.Post("/enable", s.enableSource)
	})
}
