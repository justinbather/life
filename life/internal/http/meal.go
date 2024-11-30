package http

import (
	"fmt"

	"github.com/justinbather/life/life/model"
)

var mealUri string = "/meals"

func CreateMeal(meal model.Meal, jwt, apiUrl string) (model.Meal, error) {
	return post(meal, apiUrl+mealUri, jwt)
}

func GetMeals(user string, tf map[string]string, jwt, apiUrl string) ([]model.Meal, error) {
	fullPath := apiUrl + mealUri + fmt.Sprintf("/%s/%s", tf["start"], tf["end"])
	return get(fullPath, model.Meal{}, jwt)
}
