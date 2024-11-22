package service

import (
	"github.com/justinbather/life/cli/internal/http"
	"github.com/justinbather/life/cli/model"
)

func CreateMeal(meal model.Meal) (model.Meal, error) {
	return http.CreateMeal(meal)
}
