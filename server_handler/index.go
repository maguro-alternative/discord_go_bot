package serverHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maguro-alternative/discord_go_bot/service"
	"github.com/maguro-alternative/discord_go_bot/model"
)

type IndexHandler struct {
	svc *service.IndexService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewIndexHandler(svc *service.IndexService) *IndexHandler {
	return &IndexHandler{
		svc: svc,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&model.IndexResponse{
		Message: "OK",
	})
	if err != nil {
		log.Println(err)
	}
}
