package http

import (
	"fmt"

	"github.com/justinbather/life/life/model"
)

var workoutUri string = "/workouts"

func CreateWorkout(workout model.Workout) (model.Workout, error) {
	return create(workout, workoutUri)
}

func GetWorkouts(user string, dateRange map[string]string) ([]model.Workout, error) {
	fullPath := workoutUri + fmt.Sprintf("/%s/%s/%s", user, dateRange["start"], dateRange["end"])
	return get(fullPath, model.Workout{})
}
