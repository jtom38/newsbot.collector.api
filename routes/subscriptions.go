package routes

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// GetSubscriptions
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Tags     config, Subscriptions
// @Router   /subscriptions [get]
func (s *Server) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := s.Db.ListSubscriptions(*s.ctx, 100)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	bres, err := json.Marshal(res)
	if err != nil {
		http.Error(w, ErrUnableToConvertToJson, http.StatusBadRequest)
		panic(err)
	}

	w.Write(bres)
}

// GetSubscriptionsByDiscordId
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Param    id  query  string  true  "id"
// @Tags     config, Subscriptions
// @Router   /subscriptions/byDiscordId [get]
func (s *Server) GetSubscriptionsByDiscordId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	_id := query["id"][0]
	if _id == "" {
		http.Error(w, ErrIdValueMissing, http.StatusBadRequest)
		return
	}

	uuid, err := uuid.Parse(_id)
	if err != nil {
		http.Error(w, ErrValueNotUuid, http.StatusBadRequest)
		return
	}

	res, err := s.Db.GetSubscriptionsByDiscordWebHookId(*s.ctx, uuid)
	if err != nil {
		http.Error(w, ErrNoRecordFound, http.StatusBadRequest)
		return
	}

	bres, err := json.Marshal(res)
	if err != nil {
		http.Error(w, ErrUnableToConvertToJson, http.StatusBadRequest)
		return
	}

	w.Write(bres)
}

// GetSubscriptionsBySourceId
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Param    id  query  string  true  "id"
// @Tags     config, Subscriptions
// @Router   /subscriptions/bySourceId [get]
func (s *Server) GetSubscriptionsBySourceId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	_id := query["id"][0]
	if _id == "" {
		http.Error(w, ErrIdValueMissing, http.StatusBadRequest)
		return
	}

	uuid, err := uuid.Parse(_id)
	if err != nil {
		http.Error(w, ErrValueNotUuid, http.StatusBadRequest)
		return
	}

	res, err := s.Db.GetSubscriptionsByDiscordWebHookId(*s.ctx, uuid)
	if err != nil {
		http.Error(w, ErrNoRecordFound, http.StatusBadRequest)
		return
	}

	bres, err := json.Marshal(res)
	if err != nil {
		http.Error(w, ErrUnableToConvertToJson, http.StatusBadRequest)
		return
	}

	w.Write(bres)
}