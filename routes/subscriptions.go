package routes

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
)

// GetSubscriptions
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Tags     Config, Subscription
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
// @Tags     Config, Subscription
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
// @Tags     Config, Subscription
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

// NewDiscordWebHookSubscription
// @Summary  Creates a new subscription to link a post from a Source to a DiscordWebHook.
// @Param    discordWebHookId  query  string  true  "discordWebHookId"
// @Param    sourceId          query  string  true  "sourceId"
// @Tags     Config, Source, Discord, Subscription
// @Router   /subscriptions/new/discordwebhook [post]
func (s *Server) newDiscordWebHookSubscription(w http.ResponseWriter, r *http.Request) {
	// Extract the values given
	query := r.URL.Query()
	discordWebHookId := query["discordWebHookId"][0]
	sourceId := query["sourceId"][0]

	// Check to make we didnt get a null
	if discordWebHookId == "" {
		http.Error(w, "invalid discordWebHooksId given", http.StatusBadRequest )
		return
	}
	if sourceId == "" {
		http.Error(w, "invalid sourceID given", http.StatusBadRequest )
		return
	}

	// Valide they are UUID values
	uHook, err := uuid.Parse(discordWebHookId)
	if err != nil {
		http.Error(w, "DiscordWebHooksID was not a uuid value.", http.StatusBadRequest)
		return
	}
	uSource, err := uuid.Parse(sourceId)
	if err != nil {
		http.Error(w, "SourceId was not a uuid value", http.StatusBadRequest)
		return
	}
	
	params := database.CreateSubscriptionParams{
		ID: uuid.New(),
		Discordwebhookid: uHook,
		Sourceid: uSource,
	}
	s.Db.CreateSubscription(*s.ctx, params)
	
	bJson, err := json.Marshal(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bJson)
}