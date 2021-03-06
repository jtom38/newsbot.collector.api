package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jtom38/newsbot/collector/database"
)

// GetDiscordWebHooks
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Tags     Config, Discord, Webhook
// @Router   /discord/webhooks [get]
func (s *Server) GetDiscordWebHooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := s.Db.ListDiscordWebhooks(*s.ctx, 100)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	bres, err := json.Marshal(res)
	if err != nil {
		w.Write([]byte(err.Error()))
		panic(err)
	}

	w.Write(bres)
}

// GetDiscorWebHooksById
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Param    id  query  string  true  "id"
// @Tags     Config, Discord, Webhook
// @Router   /discord/webhooks/byId [get]
func (s *Server) GetDiscordWebHooksById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	_id := query["id"][0]
	if _id == "" {
		http.Error(w, "id is missing", http.StatusBadRequest)
		return
	}

	uuid, err := uuid.Parse(_id)
	if err != nil {
		http.Error(w, "unable to parse id value", http.StatusBadRequest)
		return
	}
	
	res, err := s.Db.GetDiscordWebHooksByID(*s.ctx, uuid)
	if err != nil {
		http.Error(w, "no record found", http.StatusBadRequest)
		return
	}

	bres, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "unable to convert to json", http.StatusBadRequest)
		panic(err)
	}

	w.Write(bres)
}

// NewDiscordWebHook
// @Summary  Creates a new record for a discord web hook to post data to.
// @Param    url      query  string  true  "url"
// @Param    server   query  string  true  "Server name"
// @Param    channel  query  string  true  "Channel name"
// @Tags     Config, Discord, Webhook
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
	if _server ==  ""{
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
