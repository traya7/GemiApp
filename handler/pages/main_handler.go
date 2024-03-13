package pages

import (
	"GemiApp/helpers"
	"GemiApp/types"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
)

func SafeURL(url string) template.URL {
	return template.URL(url)
}

func (h *PageHandler) handleHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{}

	data["user"] = r.Context().Value("data")
	data["menu_items"] = helpers.MenuBuilder()
	data["games"] = h.gmeSrv.GetAllGames()

	var funcMap = template.FuncMap{"safeURL": SafeURL}
	tmpl := template.Must(
		template.New("indx.tmpl").Funcs(sprig.FuncMap()).Funcs(funcMap).ParseFiles(
			PagesPath["Menu"],
			PagesPath["HomePage"],
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "index.tmpl", data)
}

func (h *PageHandler) handleGameDetails(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{}
	data["user"] = r.Context().Value("data")
	data["menu_items"] = helpers.MenuBuilder()
	data["game_status"] = true

	gd, err := h.gmeSrv.GetGameDetails("ludo_classic")
	if err != nil {
		data["game_status"] = false
	}
	data["game_detail"] = gd

	var funcMap = template.FuncMap{"safeURL": SafeURL}
	tmpl := template.Must(
		template.New("game_details.tmpl").Funcs(sprig.FuncMap()).Funcs(funcMap).ParseFiles(
			PagesPath["Menu"],
			PagesPath["GameDetailsPage"],
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "game_details.tmpl", data)
}

func (h *PageHandler) handleGameView(w http.ResponseWriter, r *http.Request) {
	usr := r.Context().Value("data").(*types.User)
	gid := r.URL.Query().Get("gid")
	session := h.gmeSrv.NewGameSession(usr.ID, gid)

	data := map[string]any{
		"user":       usr,
		"menu_items": helpers.MenuBuilder(),
		"info":       session,
	}
	if session["status"].(bool) {
		data["gameUri"] = h.gmeSrv.GameUriFrom(session["id"].(string))
	}

	tmpl := template.Must(
		template.New("game_view.tmpl").ParseFiles(
			PagesPath["Menu"],
			PagesPath["GameViewPage"],
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "game_view.tmpl", data)
}
