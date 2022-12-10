package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
)

func (s *Server) GetSourcesRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.listSources)
	r.Get("/by/source", s.listSourcesBySource)
	r.Get("/by/sourceAndName", s.GetSourceBySourceAndName)

	r.Post("/new/reddit", s.newRedditSource)
	r.Post("/new/youtube", s.newYoutubeSource)
	r.Post("/new/twitch", s.newTwitchSource)

	r.Route("/{ID}", func(p chi.Router) {
		p.Get("/", s.getSources)
		p.Delete("/", s.deleteSources)
		p.Post("/disable", s.disableSource)
		p.Post("/enable", s.enableSource)
	})

	return r
}

type ListSourcesResults struct {
	StatusCode int                  `json:"status"`
	Message    string               `json:"message"`
	Payload    []database.SourceDto `json:"payload"`
}

type GetSourceResult struct {
	StatusCode int                `json:"status"`
	Message    string             `json:"message"`
	Payload    database.SourceDto `json:"payload"`
}

// ListSources
// @Summary  Lists the top 50 records
// @Produce  application/json
// @Tags     Source
// @Router   /sources [get]
// @Success  200  {object}  ListSourcesResults  "ok"
// @Failure  400  {object}  models.ApiError     "Unable to reach SQL or Data problems"
func (s *Server) listSources(w http.ResponseWriter, r *http.Request) {
	//TODO Add top?
	/*
		top := chi.URLParam(r, "top")
		topInt, err := strconv.ParseInt(top, 0, 32)
		if err != nil {
			panic(err)
		}
		res, err := s.Db.ListSources(*s.ctx, int32(topInt))
	*/

	w.Header().Set("Content-Type", "application/json")
	result := ListSourcesResults{
		StatusCode: http.StatusOK,
		Message:    "OK",
	}

	// Default way of showing all sources
	res, err := s.Db.ListSources(*s.ctx, 50)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	var dto []database.SourceDto
	for _, item := range res {
		dto = append(dto, database.ConvertToSourceDto(item))
	}

	result.Payload = dto

	bResult, err := json.Marshal(result)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	w.Write(bResult)
}

// ListSourcesBySource
// @Summary  Lists the top 50 records based on the name given. Example: reddit
// @Param    source  query  string  true  "Source Name"
// @Produce  application/json
// @Tags     Source
// @Router   /sources/by/source [get]
// @Success  200  {object}  ListSourcesResults  "ok"
// @Failure  400  {object}  models.ApiError  "Unable to query SQL."
// @Failure  500  {object}  models.ApiError     "Problems with data."
func (s *Server) listSourcesBySource(w http.ResponseWriter, r *http.Request) {
	//TODO Add top?
	/*
		top := chi.URLParam(r, "top")
		topInt, err := strconv.ParseInt(top, 0, 32)
		if err != nil {
			panic(err)
		}
		res, err := s.Db.ListSources(*s.ctx, int32(topInt))
	*/
	w.Header().Set("Content-Type", "application/json")

	result := ListSourcesResults{
		StatusCode: http.StatusOK,
		Message:    "OK",
	}

	query := r.URL.Query()
	_source := query["source"][0]

	// Shows the list by Sources.source
	res, err := s.Db.ListSourcesBySource(*s.ctx, strings.ToLower(_source))
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	for _, item := range res {
		result.Payload = append(result.Payload, database.ConvertToSourceDto(item))
	}

	bResult, err := json.Marshal(result)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	w.Write(bResult)
}

// GetSource
// @Summary  Returns a single entity by ID
// @Param    id  path  string  true  "uuid"
// @Produce  application/json
// @Tags     Source
// @Router   /sources/{id} [get]
// @Success  200  {object}  GetSourceResult  "ok"
// @Failure  204  {object}  models.ApiError  "No record found."
// @Failure  400  {object}  models.ApiError     "Unable to query SQL."
// @Failure  500  {object}  models.ApiError  "Failed to process data from SQL."
func (s *Server) getSources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	res, err := s.Db.GetSourceByID(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusNoContent, nil)
		return
	}

	dto := database.ConvertToSourceDto(res)

	payload := GetSourceResult{
		Message:    "OK",
		StatusCode: http.StatusOK,
		Payload:    dto,
	}

	bResult, err := json.Marshal(payload)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	w.Write(bResult)
}

// GetSourceByNameAndSource
// @Summary  Returns a single entity by ID
// @Param    name    query  string  true  "dadjokes"
// @Param    source  query  string  true  "reddit"
// @Produce  application/json
// @Tags     Source
// @Router   /sources/by/sourceAndName [get]
func (s *Server) GetSourceBySourceAndName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	name := query["name"][0]
	if name == "" {
		http.Error(w, "Parameter 'name' was missing in the query.", http.StatusInternalServerError)
		return
	}

	source := query["source"][0]
	if source == "" {
		http.Error(w, "The parameter 'source' was missing in the query.", http.StatusInternalServerError)
		return
	}

	item, err := s.Db.GetSourceByNameAndSource(context.Background(), database.GetSourceByNameAndSourceParams{
		Name:   name,
		Source: source,
	})
	if err != nil {
		http.Error(w, "Unable to find the requested record.", http.StatusInternalServerError)
	}

	bResult, err := json.Marshal(item)
	if err != nil {
		log.Panicln(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bResult)
}

// NewRedditSource
// @Summary  Creates a new reddit source to monitor.
// @Param    name  query  string  true  "name"
// @Param    url   query  string  true  "url"
// @Tags     Source, Reddit
// @Router   /sources/new/reddit [post]
func (s *Server) newRedditSource(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_name := query["name"][0]
	_url := query["url"][0]
	//_tags := query["tags"][0]

	if _url == "" {
		http.Error(w, "url is missing a value", http.StatusBadRequest)
		return
	}
	if !strings.Contains(_url, "reddit.com") {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	/*
		var tags string
		if _tags == "" {
			tags = fmt.Sprintf("twitch, %v", _name)
		} else {
		}
	*/
	tags := fmt.Sprintf("twitch, %v", _name)

	params := database.CreateSourceParams{
		ID:      uuid.New(),
		Site:    "reddit",
		Name:    _name,
		Source:  "reddit",
		Type:    "feed",
		Enabled: true,
		Url:     _url,
		Tags:    tags,
	}
	s.Db.CreateSource(*s.ctx, params)

	bJson, err := json.Marshal(&params)
	if err != nil {
		log.Panicln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bJson)
}

// NewYoutubeSource
// @Summary  Creates a new youtube source to monitor.
// @Param    name  query  string  true  "name"
// @Param    url   query  string  true  "url"
// @Tags     Source, YouTube
// @Router   /sources/new/youtube [post]
func (s *Server) newYoutubeSource(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_name := query["name"][0]
	_url := query["url"][0]
	//_tags := query["tags"][0]

	if _url == "" {
		http.Error(w, "url is missing a value", http.StatusBadRequest)
		return
	}
	if !strings.Contains(_url, "youtube.com") {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	/*
		if _tags == "" {
			tags = fmt.Sprintf("twitch, %v", _name)
			} else {
			}
	*/
	tags := fmt.Sprintf("twitch, %v", _name)

	params := database.CreateSourceParams{
		ID:      uuid.New(),
		Site:    "youtube",
		Name:    _name,
		Source:  "youtube",
		Type:    "feed",
		Enabled: true,
		Url:     _url,
		Tags:    tags,
	}
	s.Db.CreateSource(*s.ctx, params)

	bJson, err := json.Marshal(&params)
	if err != nil {
		log.Panicln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bJson)
}

// NewTwitchSource
// @Summary  Creates a new twitch source to monitor.
// @Param    name  query  string  true  "name"
// @Tags     Source, Twitch
// @Router   /sources/new/twitch [post]
func (s *Server) newTwitchSource(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_name := query["name"][0]

	tags := fmt.Sprintf("twitch, %v", _name)
	_url := fmt.Sprintf("https://twitch.tv/%v", _name)

	params := database.CreateSourceParams{
		ID:      uuid.New(),
		Site:    "twitch",
		Name:    _name,
		Source:  "twitch",
		Type:    "api",
		Enabled: true,
		Url:     _url,
		Tags:    tags,
	}
	s.Db.CreateSource(*s.ctx, params)

	bJson, err := json.Marshal(&params)
	if err != nil {
		log.Panicln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bJson)
}

// DeleteSource
// @Summary  Marks a source as deleted based on its ID value.
// @Param    id  path  string  true  "id"
// @Tags     Source
// @Router   /sources/{id} [POST]
func (s *Server) deleteSources(w http.ResponseWriter, r *http.Request) {
	//var item model.Sources = model.Sources{}

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Panicln(err)
	}

	// Check to make sure we can find the record
	_, err = s.Db.GetSourceByID(*s.ctx, uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Delete the record
	err = s.Db.DeleteSource(*s.ctx, uuid)
	if err != nil {
		log.Panic(err)
	}
}

// DisableSource
// @Summary  Disables a source from processing.
// @Param    id  path  string  true  "id"
// @Tags     Source
// @Router   /sources/{id}/disable [post]
func (s *Server) disableSource(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Panicln(err)
	}

	// Check to make sure we can find the record
	_, err = s.Db.GetSourceByID(*s.ctx, uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = s.Db.DisableSource(*s.ctx, uuid)
	if err != nil {
		log.Panic(err)
	}
}

// EnableSource
// @Summary  Enables a source to continue processing.
// @Param    id  path  string  true  "id"
// @Tags     Source
// @Router   /sources/{id}/enable [post]
func (s *Server) enableSource(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Panicln(err)
	}

	// Check to make sure we can find the record
	_, err = s.Db.GetSourceByID(*s.ctx, uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = s.Db.EnableSource(*s.ctx, uuid)
	if err != nil {
		log.Panic(err)
	}
}
