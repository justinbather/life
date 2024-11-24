package http

import (
	"fmt"

	"github.com/justinbather/life/life/model"
)

var mealUri string = "/meals"

func CreateMeal(meal model.Meal) (model.Meal, error) {
	return create(meal, mealUri)
}

func GetMeals(user string, tf map[string]string) ([]model.Meal, error) {
	fullPath := mealUri + fmt.Sprintf("/%s/%s/%s", user, tf["start"], tf["end"])
	return get(fullPath, model.Meal{})
}
