package pages

import (
	"GemiApp/app/helpers"
	"GemiApp/app/middleware"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
)

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
		template.New("indx.tmpl").Funcs(sprig.FuncMap()).Funcs(funcMap).ParseFiles(
			"./app/web/views/comps/base_menu.tmpl",
			"./app/web/views/index.tmpl",
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "index.tmpl", data)
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
		template.New("indx.tmpl").Funcs(sprig.FuncMap()).Funcs(funcMap).ParseFiles(
			"./app/web/views/comps/base_menu.tmpl",
			"./app/web/views/index.tmpl",
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.ExecuteTemplate(w, "index.tmpl", data)
}
