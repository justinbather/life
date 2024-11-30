package http

import (
	"fmt"

	"github.com/justinbather/life/life/model"
)

var workoutUri string = "/workouts"

func CreateWorkout(workout model.Workout, jwt, apiUrl string) (model.Workout, error) {
	return post(workout, apiUrl+workoutUri, jwt)
}

func GetWorkouts(user string, dateRange map[string]string, jwt, apiUrl string) ([]model.Workout, error) {
	fullPath := apiUrl + workoutUri + fmt.Sprintf("/%s/%s", dateRange["start"], dateRange["end"])
	return get(fullPath, model.Workout{}, jwt)
}
