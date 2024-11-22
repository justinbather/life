package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/justinbather/life/cli/model"
)

var URL string = "http://localhost:8080/workouts"

func CreateWorkout(workout model.Workout) (model.Workout, error) {
	data, err := json.Marshal(workout)
	if err != nil {
		return model.Workout{}, err
	}

	r, err := http.NewRequest("POST", URL, bytes.NewBuffer(data))
	if err != nil {
		return model.Workout{}, err
	}

	r.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		return model.Workout{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Got non 201 response from create endpoint. got a %d", resp.StatusCode)
		return model.Workout{}, fmt.Errorf("Error creating workout")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Workout{}, err
	}

	var created model.Workout
	err = json.Unmarshal(body, &created)

	return created, nil
}
