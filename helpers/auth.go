package helpers

import (
	"errors"
	"net/http"
	"time"
)

func AuthMiddleware(r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie("sid")
	if err != nil || cookie.Value == "" {
		return nil, errors.New("cooike not found")
	}
	return cookie, nil
}

func NewEmptyCookie() *http.Cookie {
	return &http.Cookie{
		Name:    "sid",
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	}
}
