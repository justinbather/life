package service

import (
	"github.com/justinbather/life/life/internal/http"
	"github.com/justinbather/life/life/model"
)

func CreateMeal(meal model.Meal, jwt, apiUrl string) (model.Meal, error) {
	return http.CreateMeal(meal, jwt, apiUrl)
}

func GetMeals(user string, tf map[string]string, jwt, apiUrl string) ([]model.Meal, error) {
	return http.GetMeals(user, tf, jwt, apiUrl)
}
