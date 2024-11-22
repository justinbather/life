package http

import (
	"github.com/justinbather/life/cli/model"
)

var workoutUri string = "/workouts"

func CreateWorkout(workout model.Workout) (model.Workout, error) {
	return create(workout, workoutUri)
}
