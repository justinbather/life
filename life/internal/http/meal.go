package http

import (
	"fmt"

	"github.com/justinbather/life/life/model"
)

var mealUri string = "/meals"

func CreateMeal(meal model.Meal, jwt string) (model.Meal, error) {
	return post(meal, mealUri)
}

func GetMeals(user string, tf map[string]string, jwt string) ([]model.Meal, error) {
	fullPath := mealUri + fmt.Sprintf("/%s/%s/%s", user, tf["start"], tf["end"])
	return get(fullPath, model.Meal{})
}
