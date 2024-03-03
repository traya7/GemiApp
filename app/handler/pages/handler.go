package pages

import (
	"GemiApp/app/helpers"
	"GemiApp/app/middleware"
	"GemiApp/services/auth"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
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

	r.HandleFunc("/game/ludo/lobby", h.handleLudoLobby)
}

func SafeURL(url string) template.URL {
	return template.URL(url)
}

func (h *PageHandler) handleHome(w http.ResponseWriter, r *http.Request) {
	usr, err := middleware.AuthMiddleware(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	data := map[string]any{}
	if data, err = h.srv.UserStatus(usr.UserID); err != nil {
		cookie := middleware.NewEmptyCookie()
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		w.Write(nil)
		return
	}

	data["menu_items"] = helpers.MenuBuilder()
	data["games"] = helpers.StaticGames()

	var funcMap = template.FuncMap{"safeURL": SafeURL}
	tmpl := template.Must(
		template.New("indx.tmpl").Funcs(sprig.FuncMap()).Funcs(funcMap).ParseFiles("./app/web/views/index.tmpl"),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "index.tmpl", data)
}
