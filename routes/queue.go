package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jtom38/newsbot/collector/domain/models"
)

type ListDiscordWebHooksQueueResults struct {
	ApiStatusModel
	Payload []models.DiscordQueueDetailsDto `json:"payload"`
}

func (s *Server) GetQueueRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/discord/webhooks", s.ListDiscordWebhookQueue)

	return r
}

// GetDiscordQueue
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Tags     Queue
// @Router   /queue/discord/webhooks [get]
// @Success  200  {object}  ListDiscordWebHooksQueueResults  "ok"
func (s *Server) ListDiscordWebhookQueue(w http.ResponseWriter, r *http.Request) {
	p := ListDiscordWebHooksQueueResults{
		ApiStatusModel: ApiStatusModel{
			Message:    "OK",
			StatusCode: http.StatusOK,
		},
	}

	// Get the raw resp from sql
	res, err := s.dto.ListDiscordWebhookQueueDetails(r.Context(), 50)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p.Payload = res
	s.WriteJson(w, p)
}
