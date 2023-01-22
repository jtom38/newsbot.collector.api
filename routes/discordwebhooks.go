package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
	"github.com/jtom38/newsbot/collector/domain/models"
)

func (s Server) DiscordWebHookRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.ListDiscordWebHooks)
	r.Post("/new", s.NewDiscordWebHook)
	r.Get("/by/serverAndChannel", s.GetDiscordWebHooksByServerAndChannel)
	r.Route("/{ID}", func(r chi.Router) {
		r.Get("/", s.GetDiscordWebHooksById)
		r.Delete("/", s.deleteDiscordWebHook)
		r.Post("/disable", s.disableDiscordWebHook)
		r.Post("/enable", s.enableDiscordWebHook)
	})

	return r
}

type ListDiscordWebhooks struct {
	ApiStatusModel
	Payload []models.DiscordWebHooksDto `json:"payload"`
}

// ListDiscordWebhooks
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Tags     Discord, Webhook
// @Router   /discord/webhooks [get]
func (s *Server) ListDiscordWebHooks(w http.ResponseWriter, r *http.Request) {
	p := ListDiscordWebhooks{
		ApiStatusModel: ApiStatusModel{
			Message:    "OK",
			StatusCode: http.StatusOK,
		},
	}

	w.Header().Set(HeaderContentType, ApplicationJson)

	res, err := s.dto.ListDiscordWebHooks(r.Context(), 50)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.Payload = res

	bres, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bres)
}

type GetDiscordWebhook struct {
	ApiStatusModel
	Payload models.DiscordWebHooksDto `json:"payload"`
}

// GetDiscordWebHook
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Param    id  path  string  true  "id"
// @Tags     Discord, Webhook
// @Router   /discord/webhooks/{id} [get]
// @Success  200  {object}  GetDiscordWebhook  "OK"
func (s *Server) GetDiscordWebHooksById(w http.ResponseWriter, r *http.Request) {
	p := GetDiscordWebhook{
		ApiStatusModel: ApiStatusModel{
			Message:    "OK",
			StatusCode: http.StatusOK,
		},
	}

	w.Header().Set(HeaderContentType, ApplicationJson)

	_id := chi.URLParam(r, "ID")
	if _id == "" {
		s.WriteError(w, "id is missing", http.StatusBadRequest)
		return
	}

	uuid, err := uuid.Parse(_id)
	if err != nil {
		s.WriteError(w, "unable to parse id value", http.StatusBadRequest)
		return
	}

	res, err := s.dto.GetDiscordWebhook(r.Context(), uuid)
	if err != nil {
		s.WriteError(w, "no record found", http.StatusBadRequest)
		return
	}
	p.Payload = res

	bres, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, "unable to convert to json", http.StatusBadRequest)
		return
	}

	w.Write(bres)
}

// GetDiscordWebHookByServerAndChannel
// @Summary  Returns all the known web hooks based on the Server and Channel given.
// @Produce  application/json
// @Param    server   query  string  true  "Fancy Server"
// @Param    channel  query  string  true  "memes"
// @Tags     Discord, Webhook
// @Router   /discord/webhooks/by/serverAndChannel [get]
// @Success  200  {object}  ListDiscordWebhooks  "OK"
func (s *Server) GetDiscordWebHooksByServerAndChannel(w http.ResponseWriter, r *http.Request) {
	p := ListDiscordWebhooks{
		ApiStatusModel: ApiStatusModel{
			Message:    "OK",
			StatusCode: http.StatusOK,
		},
	}

	w.Header().Set(HeaderContentType, ApplicationJson)

	query := r.URL.Query()
	_server := query["server"][0]
	if _server == "" {
		s.WriteError(w, "ID is missing", http.StatusInternalServerError)
		return
	}

	_channel := query["channel"][0]
	if _channel == "" {
		s.WriteError(w, "Channel is missing", http.StatusInternalServerError)
		return
	}

	res, err := s.dto.GetDiscordWebHookByServerAndChannel(r.Context(), _server, _channel)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p.Payload = res

	bres, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, "unable to convert to json", http.StatusInternalServerError)
		return
	}

	w.Write(bres)
}

// NewDiscordWebHook
// @Summary  Creates a new record for a discord web hook to post data to.
// @Param    url      query  string  true  "url"
// @Param    server   query  string  true  "Server name"
// @Param    channel  query  string  true  "Channel name"
// @Tags     Discord, Webhook
// @Router   /discord/webhooks/new [post]
func (s *Server) NewDiscordWebHook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_url := query["url"][0]
	_server := query["server"][0]
	_channel := query["channel"][0]

	if _url == "" {
		http.Error(w, "url is missing a value", http.StatusBadRequest)
		return
	}
	if !strings.Contains(_url, "discord.com/api/webhooks") {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}
	if _server == "" {
		http.Error(w, "server is missing", http.StatusBadRequest)
	}
	if _channel == "" {
		http.Error(w, "channel is missing", http.StatusBadRequest)
	}
	params := database.CreateDiscordWebHookParams{
		ID:      uuid.New(),
		Url:     _url,
		Server:  _server,
		Channel: _channel,
		Enabled: true,
	}
	s.Db.CreateDiscordWebHook(*s.ctx, params)

	bJson, err := json.Marshal(&params)
	if err != nil {
		log.Panicln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bJson)
}

// DisableDiscordWebHooks
// @Summary  Disables a Webhook from being used.
// @Param    id  path  string  true  "id"
// @Tags     Discord, Webhook
// @Router   /discord/webhooks/{ID}/disable [post]
func (s *Server) disableDiscordWebHook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check to make sure we can find the record
	_, err = s.Db.GetDiscordWebHooksByID(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Db.DisableDiscordWebHook(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
	}
}

// EnableDiscordWebHook
// @Summary  Enables a source to continue processing.
// @Param    id  path  string  true  "id"
// @Tags     Discord, Webhook
// @Router   /discord/webhooks/{ID}/enable [post]
func (s *Server) enableDiscordWebHook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
	}

	// Check to make sure we can find the record
	_, err = s.Db.GetDiscordWebHooksByID(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
	}

	err = s.Db.EnableDiscordWebHook(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteDiscordWebHook
// @Summary  Deletes a record by ID.
// @Param    id  path  string  true  "id"
// @Tags     Discord, Webhook
// @Router   /discord/webhooks/{ID} [delete]
func (s *Server) deleteDiscordWebHook(w http.ResponseWriter, r *http.Request) {
	//var item model.Sources = model.Sources{}

	id := chi.URLParam(r, "ID")
	uuid, err := uuid.Parse(id)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
	}

	// Check to make sure we can find the record
	_, err = s.Db.GetDiscordQueueByID(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusBadRequest)
	}

	// Delete the record
	err = s.Db.DeleteDiscordWebHooks(*s.ctx, uuid)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateDiscordWebHook
// @Summary  Updates a valid discord webhook ID based on the body given.
// @Param    id  path  string  true  "id"
// @Tags     Discord, Webhook
// @Router   /discord/webhooks/{id} [patch]
func (s *Server) UpdateDiscordWebHook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "ID")

	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Panicln(err)
	}

	// Check to make sure we can find the record
	_, err = s.Db.GetDiscordQueueByID(*s.ctx, uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
