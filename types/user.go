package types

import "net/http"

type User struct {
	ID       string
	Username string
	Balance  float64
	ImgUri   string
	Cookie   *http.Cookie
}
