package service

import (
	"github.com/justinbather/life/cli/internal/http"
	"github.com/justinbather/life/cli/model"
)

func CreateWorkout(workout model.Workout) (model.Workout, error) {
	return http.CreateWorkout(workout)
}
