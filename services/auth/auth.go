package auth

import (
	"GemiApp/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type AuthService struct {
	AuthUri string
}

func NewAuthService(AuthUri string) *AuthService {
	return &AuthService{
		AuthUri: AuthUri,
	}
}

func (s *AuthService) UserLogin(username, password string) (*types.User, error) {
	to := fmt.Sprintf("%s%s", s.AuthUri, "/login")
	body := map[string]any{"username": username, "password": password}

	result, err := MakePostRequest(to, body, nil)
	if err != nil {
		return nil, err
	}
	return &types.User{
		ID:       result["ID"].(string),
		Username: result["Username"].(string),
		Balance:  result["Balance"].(float64),
		ImgUri:   result["ImgUri"].(string),
		Cookie:   result["cookie"].(*http.Cookie),
	}, nil
}

func (s *AuthService) UserResetPwd(cookie *http.Cookie, oldp, newp string) error {
	to := fmt.Sprintf("%s%s", s.AuthUri, "/resetpwd")
	body := map[string]any{
		"old_pwd": oldp,
		"new_pwd": newp,
	}
	_, err := MakePostRequest(to, body, cookie)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) UserStatus(cookie *http.Cookie) (*types.User, error) {
	to := fmt.Sprintf("%s%s", s.AuthUri, "/status")
	body := map[string]any{}
	result, err := MakePostRequest(to, body, cookie)
	if err != nil {
		return nil, err
	}
	return &types.User{
		ID:       result["ID"].(string),
		Username: result["Username"].(string),
		Balance:  result["Balance"].(float64),
		ImgUri:   result["ImgUri"].(string),
	}, nil
}

func MakePostRequest(to string, body map[string]any, cookie *http.Cookie) (map[string]any, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("internal error [001].")
	}
	r, err := http.NewRequest("POST", to, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.New("internal error [002].")
	}
	r.Header.Add("Content-Type", "application/json")
	if cookie != nil {
		r.AddCookie(cookie)
	}
	client := &http.Client{}
	rs, err := client.Do(r)
	if err != nil {
		return nil, errors.New("internal error [003]")
	}
	defer rs.Body.Close()
	var result map[string]any
	if err := json.NewDecoder(rs.Body).Decode(&result); err != nil {
		return nil, errors.New("internal error [004].")
	}
	if rs.StatusCode != 200 {
		return nil, errors.New(result["message"].(string))
	}
	sid := rs.Header.Get("Set-Cookie")
	if sid != "" {
		result["cookie"] = rs.Cookies()[0]
	}
	return result, nil
}
