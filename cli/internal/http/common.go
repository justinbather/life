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
