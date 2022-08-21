package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
)

// ListSources
// @Summary  Lists the top 50 records
// @Produce  application/json
// @Tags     Config, Source
// @Router   /config/sources [get]
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

	// Default way of showing all sources
	res, err := s.Db.ListSources(*s.ctx, 50)
	if err != nil {
		http.Error(w, "url is missing a value", http.StatusBadRequest)
		return
	}
	bResult, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "unable to convert to json", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bResult)

}

// ListSourcesBySource
// @Summary  Lists the top 50 records based on the name given. Example: reddit
// @Param    source  query  string  true  "Source Name"
// @Produce  application/json
// @Tags     Config, Source
// @Router   /config/sources/by/source [get]
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

	query := r.URL.Query()
	_source := query["source"][0]

	// Shows the list by Sources.source
	res, err := s.Db.ListSourcesBySource(*s.ctx, strings.ToLower(_source))
	if err != nil {
		http.Error(w, "invalid source is missing a value", http.StatusBadRequest)
		return
	}
	bResult, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "unable to convert to json", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bResult)
}

// GetSource
// @Summary  Returns a single entity by ID
// @Param    id  path  string  true  "uuid"
// @Produce  application/json
// @Tags     Config, Source
// @Router   /config/sources/{id} [get]
func (s *Server) getSources(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")

	uuid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "id is not a uuid", http.StatusBadRequest)
		return
	}

	res, err := s.Db.GetSourceByID(*s.ctx, uuid)
	if err != nil {
		http.Error(w, "invalid id was given", http.StatusBadRequest)
		panic(err)
	}

	bResult, err := json.Marshal(res)
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
// @Tags     Config, Source, Reddit
// @Router   /config/sources/new/reddit [post]
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
// @Tags     Config, Source, YouTube
// @Router   /config/sources/new/youtube [post]
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
// @Tags     Config, Source, Twitch
// @Router   /config/sources/new/twitch [post]
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
// @Summary  Deletes a record by ID.
// @Param    id  path  string  true  "id"
// @Tags     Config, Source
// @Router   /config/sources/{id} [delete]
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
// @Tags     Config, Source
// @Router   /config/sources/{id}/disable [post]
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
// @Tags     Config, Source
// @Router   /config/sources/{id}/enable [post]
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
