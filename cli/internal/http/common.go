package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var baseUri string = "http://localhost:8080"

func create[T any](v T, uri string) (T, error) {
	var idk T

	data, err := json.Marshal(v)
	if err != nil {
		return idk, err
	}

	r, err := http.NewRequest("POST", baseUri+uri, bytes.NewBuffer(data))
	if err != nil {
		return idk, err
	}

	r.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		return idk, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Got non 201 response from create endpoint. got a %d", resp.StatusCode)
		return idk, fmt.Errorf("Error creating workout")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return idk, err
	}

	var created T
	err = json.Unmarshal(body, &created)

	return created, nil
}

func get[T any](uri string, _ T) ([]T, error) {

	req, err := http.NewRequest("GET", baseUri+uri, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error fetching workouts")
	}

	var items []T

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
