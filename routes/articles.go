package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// ListArticles
// @Summary  Lists the top 50 records
// @Produce  application/json
// @Tags     articles
// @Router   /articles [get]
func (s *Server) listArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := s.Db.ListArticles(*s.ctx, 50)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	bres, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	w.Write(bres)
}

// GetArticleById
// @Summary  Returns an article based on defined ID.
// @Param    id  path  string  true  "uuid"
// @Produce  application/json
// @Tags     articles
// @Router   /articles/{id} [get]
func (s *Server) getArticleById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	res, err := s.Db.GetArticleByID(*s.ctx, uuid)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	bres, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	w.Write(bres)
}

// TODO add page support
// GetArticlesBySourceID
// @Summary  Finds the articles based on the SourceID provided.  Returns the top 50.
// @Param    id  path  string  true  "Source ID UUID"
// @Produce  application/json
// @Tags     articles
// @Router   /articles/by/sourceid/{id} [get]
func (s *Server) GetArticlesBySourceId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	res, err := s.Db.GetArticlesBySourceId(*s.ctx, uuid)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	bres, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	w.Write(bres)
}