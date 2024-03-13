package pages

import (
	"GemiApp/helpers"
	"GemiApp/types"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
)

func (h *PageHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	_, err := helpers.AuthMiddleware(r)
	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl := template.Must(template.ParseFiles(PagesPath["LoginPage"]))
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
			PagesPath["Menu"],
			PagesPath["ResetPwdPage"],
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
}

func (h *PageHandler) handleTransactions(w http.ResponseWriter, r *http.Request) {
	usr := r.Context().Value("data").(*types.User)
	data := map[string]any{}
	data["user"] = usr
	data["menu_items"] = helpers.MenuBuilder()
	data["transactions"] = h.wltSrv.GetMyTransactions(usr.Username)

	var funcMap = template.FuncMap{"safeURL": SafeURL}
	tmpl := template.Must(
		template.New("transactions.tmpl").Funcs(sprig.FuncMap()).Funcs(funcMap).ParseFiles(
			PagesPath["Menu"],
			PagesPath["TransactionPage"],
		),
	)

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
}

func (h *PageHandler) handleGameHistory(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{}
	data["user"] = r.Context().Value("data")
	data["menu_items"] = helpers.MenuBuilder()

	var funcMap = template.FuncMap{"safeURL": SafeURL}
	tmpl := template.Must(
		template.New("game_history.tmpl").Funcs(sprig.FuncMap()).Funcs(funcMap).ParseFiles(
			PagesPath["Menu"],
			PagesPath["GameHistoryPage"],
		),
	)
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
}
