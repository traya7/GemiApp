package api

import (
	"GemiApp/app/middleware"
	"encoding/json"
	"html/template"
	"net/http"
)

func (h *ApiHandler) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./app/web/views/comps/auth_error.tmpl"))
	var res struct {
		ErrorMessage string
	}

	var cred struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		res.ErrorMessage = "error cannot parse info."
		tmpl.Execute(w, res)
		return
	}

	if cred.Username == "" || cred.Password == "" {
		res.ErrorMessage = "username and password required."
		tmpl.Execute(w, res)
		return
	}

	usr, role, err := h.svc.UserLogin(cred.Username, cred.Password)
	if err != nil {
		res.ErrorMessage = err.Error()
		tmpl.Execute(w, res)
		return
	}

	cookie, err := middleware.NewSessionCookie(usr, role)
	if err != nil {
		res.ErrorMessage = "Server error, try again later."
		tmpl.Execute(w, res)
		return
	}

	http.SetCookie(w, &cookie)
	w.Header().Set("HX-Redirect", "/")
	w.Write(nil)
}

func (h *ApiHandler) handleUserLogout(w http.ResponseWriter, r *http.Request) {
	cookie := middleware.NewEmptyCookie()
	http.SetCookie(w, cookie)
	w.Header().Set("HX-Redirect", "/user/login")
	w.Write(nil)
}
