package routes

import (
	"encoding/json"
	"net/http"
)

// GetDiscordQueue
// @Summary  Returns the top 100 entries from the queue to be processed.
// @Produce  application/json
// @Tags     DiscordQueue
// @Router   /discord/queue/ [get]
func (s *Server) GetDiscordQueue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	res, err := s.Db.GetDiscordQueueItems(*s.ctx, 100)
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