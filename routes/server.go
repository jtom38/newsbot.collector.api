package routes

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/dto"
	"github.com/jtom38/newsbot/collector/services/config"
)

type Server struct {
	Router *chi.Mux
	Db     *database.Queries
	dto    dto.DtoClient
	ctx    *context.Context
}

const (
	HeaderContentType = "Content-Type"

	ApplicationJson = "application/json"
)

var (
	ErrIdValueMissing        string = "id value is missing"
	ErrValueNotUuid          string = "a value given was expected to be a uuid but was not correct."
	ErrNoRecordFound         string = "no record was found."
	ErrUnableToConvertToJson string = "Unable to convert to json"
)

func NewServer(ctx context.Context, db *database.Queries) *Server {
	s := &Server{
		ctx: &ctx,
		Db:  db,
		dto: dto.NewDtoClient(db),
	}

	//db, err := openDatabase(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//s.Db = db

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

	s.Router.Mount("/api/articles", s.GetArticleRouter())
	s.Router.Mount("/api/queue", s.GetQueueRouter())
	s.Router.Mount("/api/discord/webhooks", s.DiscordWebHookRouter())
	
	//s.Router.Get("/api/settings", s.getSettings)

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
			StatusCode: HttpStatusCode,
			Message:    errMessage,
		},
	}

	b, err := json.Marshal(e)
	if err != nil {
		http.Error(w, err.Error(), HttpStatusCode)
	}

	w.Write(b)
}

func (s *Server) WriteJson(w http.ResponseWriter, model interface{}) {
	w.Header().Set(HeaderContentType, ApplicationJson)
	
	bres, err := json.Marshal(model)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bres)
}