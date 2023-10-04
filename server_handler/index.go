package server_handler

import (
	"net/http"

	"github.com/maguro-alternative/discord_go_bot/service"
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

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
