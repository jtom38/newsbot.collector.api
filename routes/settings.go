package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// GetSettings
// @Summary  Returns a object based on the Key that was given/
// @Param    key  path  string  true  "Settings Key value"
// @Produce  application/json
// @Tags     settings
// @Router   /settings/{key} [get]
func (s *Server) getSettings(w http.ResponseWriter, r *http.Request) {
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