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
	r.Get("/details", s.ListSubscriptionDetails)
	r.Get("/by/discordId", s.GetSubscriptionsByDiscordId)
	r.Get("/by/sourceId", s.GetSubscriptionsBySourceId)
	r.Post("/discord/webhook/new", s.newDiscordWebHookSubscription)
	r.Delete("/discord/webhook/delete", s.DeleteDiscordWebHookSubscription)

	return r
}

type ListSubscriptions struct {
	ApiStatusModel
	Payload []models.SubscriptionDto `json:"payload"`
}

type GetSubscription struct {
	ApiStatusModel
	Payload models.SubscriptionDto `json:"payload"`
}

type ListSubscriptionDetails struct {
	ApiStatusModel
	Payload []models.SubscriptionDetailsDto `json:"payload"`
}

// GetSubscriptions
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Tags     Subscription
// @Router   /subscriptions [get]
// @Success  200  {object}  ListSubscriptions  "ok"
// @Failure  400  {object}  ApiError           "Unable to reach SQL."
// @Failure  500  {object}  ApiError           "Failed to process data from SQL."
func (s *Server) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	payload := ListSubscriptions{
		ApiStatusModel: ApiStatusModel{
			StatusCode: http.StatusOK,
			Message:    "OK",
		},
	}

	res, err := s.dto.ListSubscriptions(r.Context(), 50)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload.Payload = res
	s.WriteJson(w, payload)
}

// ListSubscriptionDetails
// @Summary  Returns the top 50 entries with full deatils on the source and output.
// @Produce  application/json
// @Tags     Subscription
// @Router   /subscriptions/details [get]
// @Success  200  {object}  ListSubscriptionDetails  "ok"
func (s *Server) ListSubscriptionDetails(w http.ResponseWriter, r *http.Request) {
	payload := ListSubscriptionDetails{
		ApiStatusModel: ApiStatusModel{
			StatusCode: http.StatusOK,
			Message:    "OK",
		},
	}

	res, err := s.dto.ListSubscriptionDetails(r.Context(), 50)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload.Payload = res
	s.WriteJson(w, payload)
}

// GetSubscriptionsByDiscordId
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Param    id  query  string  true  "id"
// @Tags     Subscription
// @Router   /subscriptions/by/discordId [get]
// @Success  200  {object}  ListSubscriptions  "ok"
// @Failure  400  {object}  ApiError           "Unable to reach SQL or Data problems"
// @Failure  500  {object}  ApiError           "Data problems"
func (s *Server) GetSubscriptionsByDiscordId(w http.ResponseWriter, r *http.Request) {
	p := ListSubscriptions{
		ApiStatusModel: ApiStatusModel{
			StatusCode: http.StatusOK,
			Message:    "OK",
		},
	}

	query := r.URL.Query()
	if query["id"][0] == "" {
		s.WriteError(w, ErrIdValueMissing, http.StatusBadRequest)
		return
	}

	uuid, err := uuid.Parse(query["id"][0])
	if err != nil {
		s.WriteError(w, ErrValueNotUuid, http.StatusBadRequest)
		return
	}

	res, err := s.dto.ListSubscriptionsByDiscordWebhookId(r.Context(), uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusNoContent)
		return
	}

	p.Payload = res
	s.WriteJson(w, p)
}

// GetSubscriptionsBySourceId
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Param    id  query  string  true  "id"
// @Tags     Subscription
// @Router   /subscriptions/by/SourceId [get]
// @Success  200  {object}  ListSubscriptions  "ok"
func (s *Server) GetSubscriptionsBySourceId(w http.ResponseWriter, r *http.Request) {
	p := ListSubscriptions{
		ApiStatusModel: ApiStatusModel{
			StatusCode: http.StatusOK,
			Message:    "OK",
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

	res, err := s.dto.ListSubscriptionsBySourceId(r.Context(), uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusNoContent)
		return
	}

	p.Payload = res
	s.WriteJson(w, p)
}

// NewDiscordWebHookSubscription
// @Summary  Creates a new subscription to link a post from a Source to a DiscordWebHook.
// @Param    discordWebHookId  query  string  true  "discordWebHookId"
// @Param    sourceId          query  string  true  "sourceId"
// @Tags     Subscription
// @Router   /subscriptions/discord/webhook/new [post]
func (s *Server) newDiscordWebHookSubscription(w http.ResponseWriter, r *http.Request) {
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
	_, err = s.Db.QuerySubscriptions(r.Context(), database.QuerySubscriptionsParams{
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
	err = s.Db.CreateSubscription(r.Context(), params)
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
// @Param    id  query  string  true  "id"
// @Tags     Subscription
// @Router   /subscriptions/discord/webhook/delete [delete]
func (s *Server) DeleteDiscordWebHookSubscription(w http.ResponseWriter, r *http.Request) {
	var ErrMissingSubscriptionID string = "the request was missing a 'Id'"
	query := r.URL.Query()

	id := query["id"][0]
	if id == "" {
		s.WriteError(w, ErrMissingSubscriptionID, http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(query["id"][0])
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
