package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Learn more about doing custom errors like this, where to put them etc
var ERR_USER_NOT_FOUND = errors.New("User not found")

func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func getUser(r *http.Request) (string, error) {
	params := mux.Vars(r)
	user := params["user"]
	if user == "" {
		return "", ERR_USER_NOT_FOUND
	}

	return user, nil
}
