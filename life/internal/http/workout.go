package http

import (
	"fmt"

	"github.com/justinbather/life/life/model"
)

var workoutUri string = "/workouts"

func CreateWorkout(workout model.Workout, jwt string) (model.Workout, error) {
	return post(workout, workoutUri, jwt)
}

func GetWorkouts(user string, dateRange map[string]string, jwt string) ([]model.Workout, error) {
	fullPath := workoutUri + fmt.Sprintf("/%s/%s/%s", user, dateRange["start"], dateRange["end"])
	return get(fullPath, model.Workout{}, jwt)
}
