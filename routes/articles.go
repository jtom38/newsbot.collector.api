package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/domain/models"
)

func (s *Server) GetArticleRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.listArticles)
	r.Route("/{ID}", func(r chi.Router) {
		r.Get("/", s.getArticle)
		r.Get("/details", s.getArticleDetails)
	})
	r.Get("/by/sourceid", s.GetArticlesBySourceId)

	return r
}

type ArticlesListResults struct {
	ApiStatusModel
	Payload []models.ArticleDto `json:"payload"`
}

// ListArticles
// @Summary  Lists the top 50 records
// @Produce  application/json
// @Tags     Articles
// @Router   /articles [get]
// @Success  200  {object}  ArticlesListResults  "OK"
func (s *Server) listArticles(w http.ResponseWriter, r *http.Request) {
	p := ArticlesListResults{
		ApiStatusModel: ApiStatusModel{
			Message:    "OK",
			StatusCode: http.StatusOK,
		},
	}

	w.Header().Set(HeaderContentType, ApplicationJson)

	res, err := s.dto.ListArticles(r.Context(), 50)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p.Payload = res

	bres, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bres)
}

type ArticleGetResults struct {
	ApiStatusModel
	Payload models.ArticleDto `json:"payload"`
}

// GetArticle
// @Summary  Returns an article based on defined ID.
// @Param    ID  path  string  true  "uuid"
// @Produce  application/json
// @Tags     Articles
// @Router   /articles/{ID} [get]
// @Success  200  {object}  ArticleGetResults  "OK"
func (s *Server) getArticle(w http.ResponseWriter, r *http.Request) {
	p := ArticleGetResults {
		ApiStatusModel: ApiStatusModel{
			Message: "OK",
			StatusCode: http.StatusOK,
		},
	}

	w.Header().Set(HeaderContentType, ApplicationJson)

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.dto.GetArticle(r.Context(), uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p.Payload = res

	bres, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bres)
}

type ArticleDetailsResult struct {
	ApiStatusModel
	Payload models.ArticleDetailsDto `json:"payload"`
}

// GetArticleDetails
// @Summary  Returns an article and source based on defined ID.
// @Param    ID  path  string  true  "uuid"
// @Produce  application/json
// @Tags     Articles
// @Router   /articles/{ID}/details [get]
// @Success  200  {object}  ArticleDetailsResult  "OK"
func (s *Server) getArticleDetails(w http.ResponseWriter, r *http.Request) {
	p := ArticleDetailsResult {
		ApiStatusModel: ApiStatusModel{
			Message:    "OK",
			StatusCode: http.StatusOK,
		},
	}

	w.Header().Set(HeaderContentType, ApplicationJson)

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.dto.GetArticleDetails(r.Context(), uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p.Payload = res

	bres, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bres)
}

type ArticlesBySourceIDResults struct {
	ApiStatusModel
	Payload []models.ArticleDto `json:"payload"`
}

// TODO add page support
// GetArticlesBySourceID
// @Summary  Finds the articles based on the SourceID provided.  Returns the top 50.
// @Param    id  query  string  true  "Source ID UUID"
// @Produce  application/json
// @Tags     Articles
// @Router   /articles/by/sourceid [get]
// @Success  200  {object}  ArticlesListResults  "OK"
func (s *Server) GetArticlesBySourceId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.URL.Query()
	query := r.URL.Query()
	_id := query["id"][0]

	uuid, err := uuid.Parse(_id)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.dto.GetArticlesBySourceId(r.Context(), uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bres, err := json.Marshal(res)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bres)
}
