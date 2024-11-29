package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Learn more about doing custom errors like this, where to put them etc
var ERR_USER_NOT_FOUND = errors.New("User not found")
var ERR_INVALID_DATE = errors.New("Invalid date")

func encode[T any](w http.ResponseWriter, _ *http.Request, status int, v T) error {
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

func parseDateParams(r *http.Request) (map[string]time.Time, error) {
	params := mux.Vars(r)
	from := params["from"]
	to := params["to"]

	if from == "" || to == "" {
		return nil, ERR_INVALID_DATE
	}

	fromDate, err := time.Parse(time.RFC3339, from)
	if err != nil {
		return nil, err
	}

	toDate, err := time.Parse(time.RFC3339, to)
	if err != nil {
		return nil, err
	}

	mp := make(map[string]time.Time)
	mp["from"] = fromDate
	mp["to"] = toDate

	return mp, nil

}
