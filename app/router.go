package app

import (
	"GemiApp/app/handler/api"
	"GemiApp/app/handler/pages"
	"GemiApp/app/handler/ws"
	"GemiApp/services/auth"
	"GemiApp/services/wallet"
	"net/http"

	"github.com/gorilla/mux"
)

type Params struct {
	AuthService   *auth.AuthService
	WalletService *wallet.WalletService
}

func NewRouter(p Params) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(handleStaticFiles())

	pages.New(p.AuthService, p.WalletService).Route(r)
	api.New(p.AuthService).Route(r)
	ws.New(p.AuthService).Route(r)

	return r
}

func handleStaticFiles() http.Handler {
	return http.StripPrefix("/assets/", http.FileServer(http.Dir("./app/web/assets/")))
}
