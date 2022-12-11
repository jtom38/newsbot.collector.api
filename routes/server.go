package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	//"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/services/config"
)

type Server struct {
	Router *chi.Mux
	Db     *database.Queries
	ctx    *context.Context
}

var (
	ErrIdValueMissing        string = "id value is missing"
	ErrValueNotUuid          string = "a value given was expected to be a uuid but was not correct."
	ErrNoRecordFound         string = "no record was found."
	ErrUnableToConvertToJson string = "Unable to convert to json"
)

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
	//s.Router.Use(middleware.Heartbeat())
}

func (s *Server) MountRoutes() {
	s.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"), //The url pointing to API definition
	))

	/* Article Routes */
	s.Router.Get("/api/articles", s.listArticles)
	s.Router.Route("/api/articles/{ID}", func(r chi.Router) {
		r.Get("/", s.getArticleById)
	})
	s.Router.Get("/api/articles/by/sourceid", s.GetArticlesBySourceId)

	/* Queue */
	s.Router.Mount("/api/queue", s.GetQueueRouter())

	/* Discord WebHooks */
	s.Router.Get("/api/discord/webhooks", s.GetDiscordWebHooks)
	s.Router.Post("/api/discord/webhooks/new", s.NewDiscordWebHook)
	//s.Router.Get("/api/discord/webhooks/byId", s.GetDiscordWebHooksById)
	s.Router.Get("/api/discord/webhooks/by/serverAndChannel", s.GetDiscordWebHooksByServerAndChannel)

	s.Router.Route("/api/discord/webhooks/{ID}", func(r chi.Router) {
		r.Get("/", s.GetDiscordWebHooksById)
		r.Delete("/", s.deleteDiscordWebHook)
		r.Post("/disable", s.disableDiscordWebHook)
		r.Post("/enable", s.enableDiscordWebHook)
	})

	/* Settings */
	s.Router.Get("/api/settings", s.getSettings)

	s.Router.Mount("/api/sources", s.GetSourcesRouter())
	s.Router.Mount("/api/subscriptions", s.GetSubscriptionsRouter())
}

type ApiStatusModel struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type ApiError struct {
	*ApiStatusModel
}

func (s *Server) WriteError(w http.ResponseWriter, errMessage string, HttpStatusCode int) {
	e := ApiError{
		ApiStatusModel: &ApiStatusModel{
			StatusCode: http.StatusInternalServerError,
			Message:    errMessage,
		},
	}

	b, err := json.Marshal(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(b)
}
