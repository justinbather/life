package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/justinbather/life/life/internal/config"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Jwt     string    `json:"jwt"`
	Expires time.Time `json:"expires"`
}

func Authenticate(config config.Config) (string, error) {
	req := LoginRequest{Username: config.Username, Password: config.Password}

	data, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	r, err := http.NewRequest("POST", config.ApiUrl+"/auth/login", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	r.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	auth := LoginResponse{}

	err = json.Unmarshal(body, &auth)

	return auth.Jwt, nil
}
