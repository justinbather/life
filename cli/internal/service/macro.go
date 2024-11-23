package service

import (
	"fmt"

	"github.com/justinbather/life/cli/model"
)

type MacroMap map[string]int

func (m MacroMap) String() string {
	return fmt.Sprintf("Total Macros\n+Cals: %d\n-Cals: %d\nCals (sum): %d\nProtein: %d\nCarbs: %d\nFat: %d\n", m["calsIn"], m["calsBurned"], m["calsIn"]-m["calsBurned"], m["protein"], m["carbs"], m["fat"])
}

func GetMacros(user string, tf map[string]string) (MacroMap, error) {
	// for a time frame, we want calories burned, gained, and p/f/c
	workouts, err := GetWorkouts(user, tf)
	if err != nil {
		return nil, err
	}

	meals, err := GetMeals(user, tf)
	if err != nil {
		return nil, err
	}

	totCalsBurned := getCalsBurned(workouts)

	totCalsIn := getSum(func(m model.Meal) int {
		return m.Calories
	}, meals)

	totProtein := getSum(func(m model.Meal) int {
		return m.Protein
	}, meals)
	totCarbs := getSum(func(m model.Meal) int {
		return m.Carbs
	}, meals)
	totFat := getSum(func(m model.Meal) int {
		return m.Fat
	}, meals)

	macros := MacroMap{"calsBurned": totCalsBurned, "calsIn": totCalsIn, "protein": totProtein, "carbs": totCarbs, "fat": totFat}

	return macros, nil
}

func getCalsBurned(workouts []model.Workout) int {
	totCalsBurned := 0
	for _, w := range workouts {
		totCalsBurned += w.CaloriesBurned
	}

	return totCalsBurned
}

func getSum(aggregate aggregate, meals []model.Meal) int {
	tot := 0

	for _, m := range meals {
		tot += aggregate(m)
	}

	return tot

}

type aggregate func(m model.Meal) int
