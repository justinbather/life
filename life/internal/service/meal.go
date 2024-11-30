package service

import (
	"github.com/justinbather/life/life/internal/http"
	"github.com/justinbather/life/life/model"
)

func CreateMeal(meal model.Meal, jwt string) (model.Meal, error) {
	return http.CreateMeal(meal, jwt)
}

func GetMeals(user string, tf map[string]string, jwt string) ([]model.Meal, error) {
	return http.GetMeals(user, tf, jwt)
}
