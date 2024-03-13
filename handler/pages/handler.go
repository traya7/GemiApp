package pages

import (
	"GemiApp/helpers"
	"GemiApp/services/auth"
	"GemiApp/services/game"
	"GemiApp/services/wallet"
	"GemiApp/types"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

var PagesPath = map[string]string{
	"Menu": "./web/views/comps/base_menu.tmpl",

	"HomePage":        "./web/views/index.tmpl",
	"GameDetailsPage": "./web/views/game_details.tmpl",
	"GameViewPage":    "./web/views/game_view.tmpl",

	"LoginPage":       "./web/views/user/login.tmpl",
	"ResetPwdPage":    "./web/views/user/resetpwd.tmpl",
	"TransactionPage": "./web/views/user/transactions.tmpl",
	"GameHistoryPage": "./web/views/user/game_history.tmpl",
}

type PageHandler struct {
	AuthSrv *auth.AuthService
	wltSrv  *wallet.WalletService
	gmeSrv  *game.GameService
}

func New(s1 *auth.AuthService, s2 *wallet.WalletService, s3 *game.GameService) *PageHandler {
	return &PageHandler{
		AuthSrv: s1,
		wltSrv:  s2,
		gmeSrv:  s3,
	}
}

func (h *PageHandler) Route(r *mux.Router) {
	r.HandleFunc("/", h.WithAuth(h.handleHome))
	r.HandleFunc("/games/{game_id}", h.WithAuth(h.handleGameDetails))
	r.HandleFunc("/game", h.WithAuth(h.handleGameView))

	r.HandleFunc("/user/login", h.handleLogin)
	r.HandleFunc("/user/resetpwd", h.WithAuth(h.handleRestPwd))
	r.HandleFunc("/user/transactions", h.WithAuth(h.handleTransactions))
	r.HandleFunc("/user/history", h.WithAuth(h.handleGameHistory))
}

func (h *PageHandler) WithAuth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val, err := helpers.AuthMiddleware(r)
		if err != nil {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		var data *types.User
		if data, err = h.AuthSrv.UserStatus(val); err != nil {
			http.SetCookie(w, helpers.NewEmptyCookie())
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			w.Write(nil)
			return
		}
		ctx := context.WithValue(r.Context(), "data", data)
		f(w, r.WithContext(ctx))
	}
}
