package service

import (
	"github.com/justinbather/life/life/internal/http"
	"github.com/justinbather/life/life/model"
)

func CreateWorkout(workout model.Workout, jwt, apiUrl string) (model.Workout, error) {
	return http.CreateWorkout(workout, jwt, apiUrl)
}

func GetWorkouts(user string, dateRange map[string]string, jwt, apiUrl string) ([]model.Workout, error) {
	return http.GetWorkouts(user, dateRange, jwt, apiUrl)
}
