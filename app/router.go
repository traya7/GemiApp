package app

import (
	"GemiApp/app/handler/api"
	"GemiApp/app/handler/pages"
	"GemiApp/services/auth"
	"net/http"

	"github.com/gorilla/mux"
)

type Params struct {
	AuthService *auth.AuthService
}

func NewRouter(p Params) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(handleStaticFiles())

	pages.New(p.AuthService).Route(r)
	api.New(p.AuthService).Route(r)

	return r
}

func handleStaticFiles() http.Handler {
	return http.StripPrefix("/assets/", http.FileServer(http.Dir("./app/web/assets/")))
}
