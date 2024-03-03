package pages

import (
	"GemiApp/app/middleware"
	"html/template"
	"net/http"
)

func (h *PageHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.AuthMiddleware(r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := template.Must(template.ParseFiles("./app/web/views/login.tmpl"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}
