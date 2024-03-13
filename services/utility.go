package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func GameServiceRequest(to string, body any) (map[string]any, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("cannot start game. [err:001]")
	}
	r, err := http.NewRequest("POST", to, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.New("cannot start game. [err:002]")
	}
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	rs, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return nil, errors.New("cannot start game. [err:003]")
	}

	defer rs.Body.Close()
	var result map[string]any
	if err := json.NewDecoder(rs.Body).Decode(&result); err != nil {
		return nil, errors.New("cannot start game. [err:004]")
	}
	if rs.StatusCode != 200 {
		return nil, errors.New(result["message"].(string))
	}
	return result, nil
}

func MakePostRequest(to string, body map[string]any, cookie *http.Cookie) (map[string]any, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("cannot parse request.")
	}
	r, err := http.NewRequest("POST", to, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.New("cannot prepare request.")
	}
	r.Header.Add("Content-Type", "application/json")
	if cookie != nil {
		r.AddCookie(cookie)
	}
	client := &http.Client{}
	rs, err := client.Do(r)
	if err != nil {
		return nil, errors.New("cannot call server")
	}
	defer rs.Body.Close()
	var result map[string]any
	if err := json.NewDecoder(rs.Body).Decode(&result); err != nil {
		return nil, errors.New("cannot parse response data.")
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
