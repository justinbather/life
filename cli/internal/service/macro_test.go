package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/justinbather/life/cli/model"
)

// make sure we can have a macro built by joined meals and workouts, a macro from just a workout, and one from just a meal
// def want to test this deeper
func TestMacro_aggregateMacros(t *testing.T) {
	workouts := []model.Workout{
		{Date: time.Now(), CaloriesBurned: 1000},
		{Date: time.Now().AddDate(0, 0, -1), CaloriesBurned: 500},
		{Date: time.Now().AddDate(0, 0, -2), CaloriesBurned: 1000},
		{Date: time.Now().AddDate(0, 0, -4), CaloriesBurned: 1000},
		{Date: time.Now().AddDate(0, 0, -5), CaloriesBurned: 1000},
	}
	meals := []model.Meal{
		{Date: time.Now(), Calories: 100, Carbs: 50, Fat: 50},
		{Date: time.Now().AddDate(0, 0, -1), Calories: 100, Carbs: 50, Fat: 50},
		{Date: time.Now().AddDate(0, 0, -2), Calories: 1000, Carbs: 100, Fat: 100},
		{Date: time.Now().AddDate(0, 0, -3), Calories: 1000, Carbs: 100, Fat: 100}}

	macros := AggregateMacros(workouts, meals)

	fmt.Println(macros)

	if len(macros) != 6 {
		t.Fatalf("Expected 5 macros in result. Got %d", len(macros))
	}

}
