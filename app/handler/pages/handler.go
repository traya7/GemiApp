package pages

import (
	"GemiApp/services/auth"
	"github.com/gorilla/mux"
)

type PageHandler struct {
	srv *auth.AuthService
}

func New(s *auth.AuthService) *PageHandler {
	return &PageHandler{
		srv: s,
	}
}

func (h *PageHandler) Route(r *mux.Router) {
	r.HandleFunc("/", h.handleHome)
	r.HandleFunc("/login", h.handleLogin)
	r.HandleFunc("/games/{game_id}", h.handleGameDetails)

	r.HandleFunc("/game/ludo/lobby", h.handleLudoLobby)
}
