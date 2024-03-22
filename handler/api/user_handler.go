package api

import (
	"GemiApp/helpers"
	"encoding/json"
	"html/template"
	"net/http"
)

func (h *ApiHandler) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/views/comps/auth_error.tmpl"))
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
	usr, err := h.svc.UserLogin(cred.Username, cred.Password)
	if err != nil {
		res.ErrorMessage = err.Error()
		tmpl.Execute(w, res)
		return
	}
	usr.Cookie.Path = "/"
	http.SetCookie(w, usr.Cookie)
	w.Header().Set("HX-Redirect", "/")
	w.Write(nil)
}

func (h *ApiHandler) handleUserResetPwd(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/views/comps/auth_error.tmpl"))
	var res struct {
		ErrorMessage string
	}
	var cred struct {
		OldPassword  string `json:"opwd"`
		NewPassword  string `json:"npwd"`
		RNewPassword string `json:"rnpwd"`
	}

	cookie, err := helpers.AuthMiddleware(r)
	if err != nil {
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		res.ErrorMessage = "error cannot parse info."
		tmpl.Execute(w, res)
		return
	}
	if cred.OldPassword == "" || cred.NewPassword == "" || cred.RNewPassword == "" {
		res.ErrorMessage = "all fields are required."
		tmpl.Execute(w, res)
		return
	}
	if cred.NewPassword != cred.RNewPassword {
		res.ErrorMessage = "repeated password not correct."
		tmpl.Execute(w, res)
		return
	}
	err = h.svc.UserResetPwd(cookie, cred.OldPassword, cred.NewPassword)
	if err != nil {
		res.ErrorMessage = err.Error()
		tmpl.Execute(w, res)
		return
	}
	success := template.Must(template.ParseFiles("./web/views/comps/auth_ok.tmpl"))
	success.Execute(w, map[string]any{"Message": "Password updated."})
}

func (h *ApiHandler) handleUserLogout(w http.ResponseWriter, r *http.Request) {
	cookie := helpers.NewEmptyCookie()
	http.SetCookie(w, cookie)
	w.Header().Set("HX-Redirect", "/user/login")
	w.Write(nil)
}
