package pages

import (
	"GemiApp/app/middleware"
	"html/template"
	"net/http"
)

func (h *PageHandler) handleLudoLobby(w http.ResponseWriter, r *http.Request) {
	usr, err := middleware.AuthMiddleware(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("./app/web/views/ludo/lobby.tmpl"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, usr)
}
