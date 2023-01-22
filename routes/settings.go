package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s *Server) getSettings(w http.ResponseWriter, r *http.Request) {
	// GetSettings
	// @Summary  Returns a object based on the Key that was given.
	// @Param    key  path  string  true  "Settings Key value"
	// @Produce  application/json
	// @Tags     Settings
	// @Router   /settings/{key} [get]

	
	w.Header().Set("Content-Type", "application/json")

	//var item model.Sources
	id := chi.URLParam(r, "ID")

	uuid, err := uuid.Parse(id)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.Db.GetSourceByID(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusNotFound)
		return
	}

	bResult, err := json.Marshal(res)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bResult)
}
