package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/models"
)

func (s *Server) GetSubscriptionsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.ListSubscriptions)
	r.Get("/by/discordId", s.GetSubscriptionsByDiscordId)
	r.Get("/by/sourceId", s.GetSubscriptionsBySourceId)
	r.Post("/discord/webhook/new", s.newDiscordWebHookSubscription)
	r.Delete("/discord/webhook/delete", s.DeleteDiscordWebHookSubscription)

	return r
}

type ListSubscriptionResults struct {
	ApiStatusModel
	Payload    []models.SubscriptionDto `json:"payload"`
}

// GetSubscriptions
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Tags     Subscription
// @Router   /subscriptions [get]
// @Success  200  {object}  ListSubscriptionResults  "ok"
// @Failure  400  {object}  ApiError                 "Unable to reach SQL."
// @Failure  500  {object}  ApiError                 "Failed to process data from SQL."
func (s *Server) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	payload := ListSubscriptionResults{
		ApiStatusModel: ApiStatusModel{
			StatusCode: http.StatusOK,
			Message:    "OK",
		},
	}

	res, err := s.Db.ListSubscriptions(*s.ctx, 100)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, item := range res {
		payload.Payload = append(payload.Payload, models.ConvertToSubscriptionDto(item))
	}

	bres, err := json.Marshal(payload)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bres)
}

type GetSubscriptionResults struct {
	ApiStatusModel
	Payload    models.SubscriptionDto `json:"payload"`
}

// GetSubscriptionsByDiscordId
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Param    id  query  string  true  "id"
// @Tags     Subscription
// @Router   /subscriptions/by/discordId [get]
// @Success  200  {object}  ListSubscriptionResults  "ok"
// @Failure  400  {object}  ApiError                 "Unable to reach SQL or Data problems"
// @Failure  500  {object}  ApiError                 "Data problems"
func (s *Server) GetSubscriptionsByDiscordId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := ListSubscriptionResults {
		ApiStatusModel: ApiStatusModel{
			StatusCode: http.StatusOK,
			Message: "OK",
		},
	}

	query := r.URL.Query()
	_id := query["id"][0]
	if _id == "" {
		s.WriteError(w, ErrIdValueMissing, http.StatusBadRequest)
		return
	}

	uuid, err := uuid.Parse(_id)
	if err != nil {
		s.WriteError(w, ErrValueNotUuid, http.StatusBadRequest)
		return
	}

	res, err := s.Db.GetSubscriptionsByDiscordWebHookId(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusNoContent)
		return
	}

	for _, item := range res {
		p.Payload = append(p.Payload, models.ConvertToSubscriptionDto(item))
	}

	bres, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bres)
}

// GetSubscriptionsBySourceId
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Param    id  query  string  true  "id"
// @Tags     Subscription
// @Router   /subscriptions/by/SourceId [get]
// @Success  200  {object}  ListSubscriptionResults  "ok"
func (s *Server) GetSubscriptionsBySourceId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := ListSubscriptionResults {
		ApiStatusModel: ApiStatusModel{
			StatusCode: http.StatusOK,
			Message: "OK",
		},
	}

	query := r.URL.Query()
	_id := query["id"][0]
	if _id == "" {
		s.WriteError(w, ErrIdValueMissing, http.StatusBadRequest)
		return
	}

	uuid, err := uuid.Parse(_id)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := s.Db.GetSubscriptionsByDiscordWebHookId(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusNoContent)
		return
	}

	for _, item := range res {
		p.Payload = append(p.Payload, models.ConvertToSubscriptionDto(item))
	}

	bres, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bres)
}

// NewDiscordWebHookSubscription
// @Summary  Creates a new subscription to link a post from a Source to a DiscordWebHook.
// @Param    discordWebHookId  query  string  true  "discordWebHookId"
// @Param    sourceId          query  string  true  "sourceId"
// @Tags     Subscription
// @Router   /subscriptions/new/discord/webhook [post]
func (s *Server) newDiscordWebHookSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the values given
	query := r.URL.Query()
	discordWebHookId := query["discordWebHookId"][0]
	sourceId := query["sourceId"][0]

	// Check to make we didn't get a null
	if discordWebHookId == "" {
		s.WriteError(w, "invalid discordWebHooksId given", http.StatusBadRequest)
		return
	}
	if sourceId == "" {
		s.WriteError(w, "invalid sourceID given", http.StatusBadRequest)
		return
	}

	// Validate they are UUID values
	uHook, err := uuid.Parse(discordWebHookId)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
	uSource, err := uuid.Parse(sourceId)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the sub already exists
	_, err = s.Db.QuerySubscriptions(*s.ctx, database.QuerySubscriptionsParams{
		Discordwebhookid: uHook,
		Sourceid:         uSource,
	})
	if err == nil {
		s.WriteError(w, "a subscription already exists between these two entities", http.StatusBadRequest)
		return
	}

	// Does not exist, so make it.
	params := database.CreateSubscriptionParams{
		ID:               uuid.New(),
		Discordwebhookid: uHook,
		Sourceid:         uSource,
	}
	err = s.Db.CreateSubscription(*s.ctx, params)
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

// DeleteDiscordWebHookSubscription
// @Summary  Removes a Discord WebHook Subscription based on the Subscription ID.
// @Param    Id  query  string  true  "Id"
// @Tags     Subscription
// @Router   /subscriptions/discord/webhook/delete [delete]
func (s *Server) DeleteDiscordWebHookSubscription(w http.ResponseWriter, r *http.Request) {
	var ErrMissingSubscriptionID string = "the request was missing a 'Id'"
	query := r.URL.Query()

	id := query["Id"][0]
	if id == "" {
		s.WriteError(w, ErrMissingSubscriptionID, http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(query["Id"][0])
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Db.DeleteSubscription(context.Background(), uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
