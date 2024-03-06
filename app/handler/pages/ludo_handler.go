package pages

import (
	"GemiApp/app/middleware"
	"html/template"
	"net/http"
)

func (h *PageHandler) handleLudoLobby(w http.ResponseWriter, r *http.Request) {
	usr, err := middleware.AuthMiddleware(r)
	if err != nil {
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}

	gid := r.URL.Query().Get("gid")
	mid := r.URL.Query().Get("id")
	if !h.validateGameAndMode(gid, mid) {
		return
	}

	tmpl := template.Must(template.ParseFiles("./app/web/views/ludo/lobby.tmpl"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, map[string]any{
		"user":    usr,
		"game_id": gid,
		"mode_id": mid,
	})
}

func (h *PageHandler) validateGameAndMode(gid, mid string) bool {
	game, err := h.gmeSrv.GetGameDetails(gid)
	if err != nil || game.Status == false {
		return false
	}

	for _, v := range game.Modes {
		if v.ID == mid {
			return true
		}
	}
	return false
}
