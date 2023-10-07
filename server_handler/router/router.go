package router

import (
	"net/http"


	"github.com/maguro-alternative/discord_go_bot/server_handler"
	"github.com/maguro-alternative/discord_go_bot/service"
)

func NewRouter() *http.ServeMux {
	// *service.IndexService型変数を作成する。
	var indexService = service.NewIndexService()

	// register routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", serverHandler.NewIndexHandler(indexService).ServeHTTP)
	return mux
}