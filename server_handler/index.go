package serverHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maguro-alternative/discord_go_bot/model"
	"github.com/maguro-alternative/discord_go_bot/service"
)

type IndexHandler struct {
	svc *service.IndexService
}

// http.HandlerをベースにしたIndexHandlerを返す
func NewIndexHandler(svc *service.IndexService) *IndexHandler {
	return &IndexHandler{
		svc: svc,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&model.IndexResponse{
		Message: h.svc.DiscordSession.State.User.Username + " is running",
	})
	if err != nil {
		log.Println(err)
	}
}
