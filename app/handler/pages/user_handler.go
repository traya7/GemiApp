package pages

import (
	"GemiApp/app/helpers"
	"GemiApp/app/middleware"
	"html/template"
	"net/http"
	"time"

	"github.com/Masterminds/sprig/v3"
)

func (h *PageHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.AuthMiddleware(r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := template.Must(
		template.ParseFiles("./app/web/views/user/login.tmpl"),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}
func (h *PageHandler) handleRestPwd(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{}
	data["user"] = r.Context().Value("data")
	data["menu_items"] = helpers.MenuBuilder()

	var funcMap = template.FuncMap{"safeURL": SafeURL}
	tmpl := template.Must(
		template.New("resetpwd.tmpl").Funcs(sprig.FuncMap()).Funcs(funcMap).ParseFiles(
			"./app/web/views/comps/base_menu.tmpl",
			"./app/web/views/user/resetpwd.tmpl",
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
}
func (h *PageHandler) handleTransactions(w http.ResponseWriter, r *http.Request) {
	usr := r.Context().Value("data").(map[string]any)
	data := map[string]any{}
	data["user"] = usr
	data["menu_items"] = helpers.MenuBuilder()
	data["transactions"] = h.wltSrv.GetMyTransactions(usr["Username"].(string))

	var funcMap = template.FuncMap{"safeURL": SafeURL}
	tmpl := template.Must(
		template.New("transactions.tmpl").Funcs(sprig.FuncMap()).Funcs(funcMap).ParseFiles(
			"./app/web/views/comps/base_menu.tmpl",
			"./app/web/views/user/transactions.tmpl",
		),
	)

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
}

func t() {
	time.Now().Format("15:04")
}
func (h *PageHandler) handleGameHistory(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{}
	data["user"] = r.Context().Value("data")
	data["menu_items"] = helpers.MenuBuilder()

	var funcMap = template.FuncMap{"safeURL": SafeURL}
	tmpl := template.Must(
		template.New("game_history.tmpl").Funcs(sprig.FuncMap()).Funcs(funcMap).ParseFiles(
			"./app/web/views/comps/base_menu.tmpl",
			"./app/web/views/user/game_history.tmpl",
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
}
