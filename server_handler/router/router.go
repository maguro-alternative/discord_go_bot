package router

import (
	"database/sql"
	"net/http"


	"github.com/maguro-alternative/discord_go_bot/server_handler"
	"github.com/maguro-alternative/discord_go_bot/service"
)

func NewRouter(indeDB *sql.DB) *http.ServeMux {
	// create a *service.TODOService type variable using the *sql.DB type variable
	var indexService = service.NewIndexService(indeDB)

	// register routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", serverHandler.NewIndexHandler(indexService).ServeHTTP)
	//mux.HandleFunc("/todos", handler.NewTODOHandler(todoService).ServeHTTP)
	return mux
}