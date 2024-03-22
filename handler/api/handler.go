package api

import (
	"GemiApp/services/auth"
	"github.com/gorilla/mux"
)

type ApiHandler struct {
	svc *auth.AuthService
}

func New(s *auth.AuthService) *ApiHandler {
	return &ApiHandler{
		svc: s,
	}
}
func (h *ApiHandler) Route(r *mux.Router) {
	r.HandleFunc("/api/user/login", h.handleUserLogin)
	r.HandleFunc("/api/user/reset", h.handleUserResetPwd)
	r.HandleFunc("/api/user/logout", h.handleUserLogout)
}
