package service

import (
	"github.com/justinbather/life/cli/internal/http"
	"github.com/justinbather/life/cli/model"
)

func CreateWorkout(workout model.Workout) (model.Workout, error) {
	return http.CreateWorkout(workout)
}

func GetWorkouts(user string, dateRange map[string]string) ([]model.Workout, error) {
	return http.GetWorkouts(user, dateRange)
}
