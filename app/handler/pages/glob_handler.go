package pages

import (
	"GemiApp/app/helpers"
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
			"./app/web/views/comps/base_menu.tmpl",
			"./app/web/views/index.tmpl",
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "index.tmpl", data)
}

func (h *PageHandler) handleGameDetails(w http.ResponseWriter, r *http.Request) {
	// Get Game id from Request

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
			"./app/web/views/comps/base_menu.tmpl",
			"./app/web/views/game_details.tmpl",
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "game_details.tmpl", data)
}
