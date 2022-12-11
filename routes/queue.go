package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jtom38/newsbot/collector/domain/models"
)

func (s *Server) GetQueueRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/discord/webhooks", s.ListDiscordWebhookQueue)

	return r
}

type ListDiscordWebHooksQueueResults struct {
	ApiStatusModel
	Payload []models.DiscordQueueDto `json:"payload"`
}

// GetDiscordQueue
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Tags     Queue
// @Router   /queue/discord/webhooks [get]
// @Success  200  {object}  ListDiscordWebHooksQueueResults  "ok"
func (s *Server) ListDiscordWebhookQueue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := ListDiscordWebHooksQueueResults{
		ApiStatusModel: ApiStatusModel{
			Message:    "OK",
			StatusCode: http.StatusOK,
		},
	}

	// Get the raw resp from sql
	res, err := s.Db.ListDiscordQueueItems(*s.ctx, 100)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// convert to dto
	for _, item := range res {
		p.Payload = append(p.Payload, models.ConvertToDiscordQueueDto(item))
	}

	// convert to json
	b, err := json.Marshal(p)
	if err != nil {
		s.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
