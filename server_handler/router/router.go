package router

import (
	"net/http"


	"github.com/maguro-alternative/discord_go_bot/server_handler"
	"github.com/maguro-alternative/discord_go_bot/service"

	"github.com/bwmarrin/discordgo"
)

func NewRouter(discordSession *discordgo.Session) *http.ServeMux {
	// *service.IndexService型変数を作成する。
	var indexService = service.NewIndexService(discordSession)

	// register routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", serverHandler.NewIndexHandler(indexService).ServeHTTP)
	return mux
}