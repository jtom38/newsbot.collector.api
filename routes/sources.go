package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
)

// ListSources
// @Summary  Lists the top 50 records
// @Param    top  query  int  true  "top"
// @Produce  application/json
// @Tags     config, source
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

	res, err := s.Db.ListSources(*s.ctx, 50)

	if err != nil {
		log.Panicln(err)
		return
	}

	//dtos := convertToSourcesDtoSlice(items)
	bResult, err := json.Marshal(res)
	if err != nil {
		log.Panicln(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bResult)
}

// GetSource
// @Summary  Returns a single entity by ID
// @Param    id  path  string  true  "uuid"
// @Produce  application/json
// @Tags     config, source
// @Router   /config/sources/{id} [get]
func (s *Server) getSources(w http.ResponseWriter, r *http.Request) {
	//var item model.Sources
	id := chi.URLParam(r, "ID")

	uuid, err := uuid.Parse(id)
	if err != nil {
		panic(err)
	}

	res, err := s.Db.GetSourceByID(*s.ctx, uuid)
	if err != nil {
		panic(err)
	}

	//itemId := fmt.Sprint(item.ID)
	//if id != itemId {
	//	log.Panicln("Unable to find the requested record.  Either unable to access SQL or the record does not exist.")
	//}

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
// @Tags     config, source, reddit
// @Router   /config/sources/ [post]
func (s *Server) postSources(w http.ResponseWriter, r *http.Request) {
	var item database.Source

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Panicln("Received a payload to but the body was incorrect")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tags := fmt.Sprintf("reddit, %v, %v", item.Name, item.Tags)
	params := database.CreateSourceParams{
		ID:      uuid.New(),
		Site:    "reddit",
		Name:    item.Name,
		Source:  "reddit",
		Type:    "feed",
		Enabled: true,
		Url:     item.Url,
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
// @Summary  Deletes a record by ID
// @Param    id  path  string  true  "id"
// @Tags     config, source
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
