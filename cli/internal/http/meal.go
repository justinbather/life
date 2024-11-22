package http

import "github.com/justinbather/life/cli/model"

var mealUri string = "/meals"

func CreateMeal(meal model.Meal) (model.Meal, error) {
	return create(meal, mealUri)
}
