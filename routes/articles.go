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
// @Tags     Articles
// @Router   /articles [get]
func (s *Server) listArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ApplicationJson)

	res, err := s.Db.ListArticlesByDate(*s.ctx, 50)
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
// @Tags     Articles
// @Router   /articles/{ID} [get]
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
// @Param    id  query  string  true  "Source ID UUID"
// @Produce  application/json
// @Tags     Articles
// @Router   /articles/by/sourceid [get]
func (s *Server) GetArticlesBySourceId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.URL.Query()
	query := r.URL.Query()
	_id := query["id"][0]

	uuid, err := uuid.Parse(_id)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	res, err := s.Db.GetNewArticlesBySourceId(*s.ctx, uuid)
	//res, err := s.Db.GetArticlesBySourceId(*s.ctx, uuid)
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
// GetArticlesByTag
// @Summary  Finds the articles based on the SourceID provided.  Returns the top 50.
// @Param    tag  query  string  true  "Tag name"
// @Produce  application/json
// @Tags     Articles
// @Router   /articles/by/tag [get]
func (s *Server) GetArticlesByTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.URL.Query()
	query := r.URL.Query()
	_id := query["tag"][0]

	uuid, err := uuid.Parse(_id)
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
