package pages

import (
	"GemiApp/app/middleware"
	"GemiApp/services/auth"
	"GemiApp/services/wallet"
	"GemiApp/types"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type PageHandler struct {
	AuthSrv *auth.AuthService
	wltSrv  *wallet.WalletService
}

func New(s *auth.AuthService, s2 *wallet.WalletService) *PageHandler {
	return &PageHandler{
		AuthSrv: s,
		wltSrv:  s2,
	}
}

func (h *PageHandler) Route(r *mux.Router) {
	r.HandleFunc("/", h.WithAuth(h.handleHome))
	r.HandleFunc("/games/{game_id}", h.WithAuth(h.handleGameDetails))

	r.HandleFunc("/user/login", h.handleLogin)
	r.HandleFunc("/user/resetpwd", h.WithAuth(h.handleRestPwd))
	r.HandleFunc("/user/transactions", h.WithAuth(h.handleTransactions))
	r.HandleFunc("/user/history", h.WithAuth(h.handleGameHistory))

	r.HandleFunc("/game/ludo/lobby", h.handleLudoLobby)
}

type httpFunc func(w http.ResponseWriter, r *http.Request)

func (h *PageHandler) WithAuth(f httpFunc) httpFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr, err := middleware.AuthMiddleware(r)
		if err != nil {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		var data *types.User
		if data, err = h.AuthSrv.UserStatus(usr.UserID); err != nil {
			cookie := middleware.NewEmptyCookie()
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			w.Write(nil)
			return
		}
		ctx := context.WithValue(r.Context(), "data", data)
		f(w, r.WithContext(ctx))
	}
}
