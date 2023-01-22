package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/models"
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

type ListSources struct {
	ApiStatusModel
	Payload []models.SourceDto `json:"payload"`
}

type GetSource struct {
	ApiStatusModel
	Payload models.SourceDto `json:"payload"`
}

// ListSources
// @Summary  Lists the top 50 records
// @Produce  application/json
// @Tags     Source
// @Router   /sources [get]
// @Success  200  {object}  ListSources  "ok"
// @Failure  400  {object}  ApiError     "Unable to reach SQL or Data problems"
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

	w.Header().Set(HeaderContentType, ApplicationJson)
	result := ListSources{
		ApiStatusModel: ApiStatusModel{
			StatusCode: http.StatusOK,
			Message:    "OK",
		},
	}

	// Default way of showing all sources
	items, err := s.dto.ListSources(r.Context(), 50)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result.Payload = items

	bResult, err := json.Marshal(result)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
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
// @Success  200  {object}  ListSources  "ok"
// @Failure  400  {object}  ApiError     "Unable to query SQL."
// @Failure  500  {object}  ApiError     "Problems with data."
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
	w.Header().Set(HeaderContentType, ApplicationJson)

	result := ListSources{
		ApiStatusModel: ApiStatusModel{
			StatusCode: http.StatusOK,
			Message:    "OK",
		},
	}

	query := r.URL.Query()
	_source := query["source"][0]

	// Shows the list by Sources.source
	res, err := s.dto.ListSourcesBySource(r.Context(), _source)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result.Payload = res

	bResult, err := json.Marshal(result)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
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
// @Success  200  {object}  GetSource  "ok"
// @Failure  204  {object}  ApiError   "No record found."
// @Failure  400  {object}  ApiError   "Unable to query SQL."
// @Failure  500  {object}  ApiError   "Failed to process data from SQL."
func (s *Server) getSources(w http.ResponseWriter, r *http.Request) {
	payload := GetSource{
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

	res, err := s.dto.GetSourceById(r.Context(), uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusNoContent)
		return
	}

	payload.Payload = res

	bResult, err := json.Marshal(payload)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
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
// @Success  200  {object}  GetSource  "ok"
// @Failure  204  {object}  ApiError   "No record found."
// @Failure  400  {object}  ApiError   "Unable to query SQL."
// @Failure  500  {object}  ApiError   "Failed to process data from SQL."
func (s *Server) GetSourceBySourceAndName(w http.ResponseWriter, r *http.Request) {
	p := GetSource{
		ApiStatusModel: ApiStatusModel{
			Message:    "OK",
			StatusCode: http.StatusOK,
		},
	}

	query := r.URL.Query()
	name := query["name"][0]
	if name == "" {
		s.WriteError(w, "Parameter 'name' was missing in the query.", http.StatusInternalServerError)
		return
	}

	source := query["source"][0]
	if source == "" {
		s.WriteError(w, "The parameter 'source' was missing in the query.", http.StatusInternalServerError)
		return
	}

	item, err := s.dto.GetSourceByNameAndSource(r.Context(), name, source)
	if err != nil {
		s.WriteError(w, "Unable to find the requested record.", http.StatusInternalServerError)
		return
	}

	p.Payload = item

	bResult, err := json.Marshal(item)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(HeaderContentType, ApplicationJson)
	w.Write(bResult)
}

// NewRedditSource
// @Summary  Creates a new reddit source to monitor.
// @Param    name  query  string  true  "name"
// @Param    url   query  string  true  "url"
// @Tags     Source
// @Router   /sources/new/reddit [post]
func (s *Server) newRedditSource(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_name := query["name"][0]
	_url := query["url"][0]
	//_tags := query["tags"][0]

	w.Header().Set("Content-Type", "application/json")

	if _url == "" {
		s.WriteError(w, "url is missing a value", http.StatusBadRequest)
		return
	}
	if !strings.Contains(_url, "reddit.com") {
		s.WriteError(w, "invalid url", http.StatusBadRequest)
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
	err := s.Db.CreateSource(*s.ctx, params)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bJson, err := json.Marshal(&params)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bJson)
}

// NewYoutubeSource
// @Summary  Creates a new youtube source to monitor.
// @Param    name  query  string  true  "name"
// @Param    url   query  string  true  "url"
// @Tags     Source
// @Router   /sources/new/youtube [post]
func (s *Server) newYoutubeSource(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_name := query["name"][0]
	_url := query["url"][0]
	//_tags := query["tags"][0]

	w.Header().Set("Content-Type", "application/json")
	if _url == "" {
		s.WriteError(w, "url is missing a value", http.StatusBadRequest)
		return
	}
	if !strings.Contains(_url, "youtube.com") {
		s.WriteError(w, "invalid url", http.StatusBadRequest)
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
	err := s.Db.CreateSource(*s.ctx, params)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bJson, err := json.Marshal(&params)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bJson)
}

// NewTwitchSource
// @Summary  Creates a new twitch source to monitor.
// @Param    name  query  string  true  "name"
// @Tags     Source
// @Router   /sources/new/twitch [post]
func (s *Server) newTwitchSource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
	err := s.Db.CreateSource(*s.ctx, params)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bJson, err := json.Marshal(&params)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bJson)
}

// DeleteSource
// @Summary  Marks a source as deleted based on its ID value.
// @Summary  Marks a source as deleted based on its ID value.
// @Param    id  path  string  true  "id"
// @Tags     Source
// @Router   /sources/{id} [POST]
func (s *Server) deleteSources(w http.ResponseWriter, r *http.Request) {
	//var item model.Sources = model.Sources{}

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check to make sure we can find the record
	_, err = s.Db.GetSourceByID(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete the record
	err = s.Db.DeleteSource(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := ApiStatusModel{
		Message:    "OK",
		StatusCode: http.StatusOK,
	}

	b, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
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
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
	}

	// Check to make sure we can find the record
	_, err = s.Db.GetSourceByID(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
	}

	err = s.Db.DisableSource(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
	}

	p := ApiStatusModel{
		Message:    "OK",
		StatusCode: http.StatusOK,
	}

	b, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
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
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
	}

	// Check to make sure we can find the record
	_, err = s.Db.GetSourceByID(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
	}

	err = s.Db.EnableSource(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
	}

	p := ApiStatusModel{
		Message:    "OK",
		StatusCode: http.StatusOK,
	}

	b, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
