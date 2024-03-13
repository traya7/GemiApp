package handler

import (
	"GemiApp/handler/api"
	"GemiApp/handler/pages"
	"GemiApp/handler/ws"

	"GemiApp/services/auth"
	"GemiApp/services/game"
	"GemiApp/services/wallet"

	"net/http"

	"github.com/gorilla/mux"
)

type Params struct {
	AuthService   *auth.AuthService
	WalletService *wallet.WalletService
	GameService   *game.GameService
}

func NewRouter(p Params) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(handleStaticFiles())

	pages.New(p.AuthService, p.WalletService, p.GameService).Route(r)

	api.New(p.AuthService).Route(r)

	ws.New(p.AuthService, p.GameService).Route(r)

	return r
}

func handleStaticFiles() http.Handler {
	return http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/")))
}
