package service

import (
	"fmt"
	"time"

	"github.com/justinbather/life/life/model"
)

/*
Full view should look like this
Really this is an array of Macros structs ? sorted by date
+==============================================================================+
|     Date     |  - Cals  |  + Cals   |  + Carbs   |   + Protein   |   + Fat   |
|==============+==========+===========+============+===============+===========|
| Sept 3, 2024 |    923   |    1599   |    1999    |      103      |     93    |
|--------------+----------+-----------+------------+---------------+-----------|
|
|
*/
type MacroMap map[string]int

type Macro struct {
	Date       time.Time
	CalsBurned int
	CalsIn     int
	Protein    int
	Carbs      int
	Fat        int
	Workouts   int
}

func (m Macro) String() string {
	return fmt.Sprintf("Macro: {\nDate=%s,\nCalsBurned=%d,\nCalsIn=%d,\nProtein=%d,\nCarbs=%d,\nFat=%d,\nWorkouts=%d\n}\n", m.Date, m.CalsBurned, m.CalsIn, m.Protein, m.Carbs, m.Fat, m.Workouts)
}

func (m MacroMap) String() string {
	return fmt.Sprintf("Total Macros\n+Cals: %d\n-Cals: %d\nCals (sum): %d\nProtein: %d\nCarbs: %d\nFat: %d\n", m["calsIn"], m["calsBurned"], m["calsIn"]-m["calsBurned"], m["protein"], m["carbs"], m["fat"])
}

func GetMacros(user string, tf map[string]string) ([]Macro, error) {
	workouts, err := GetWorkouts(user, tf)
	if err != nil {
		return nil, err
	}

	meals, err := GetMeals(user, tf)
	if err != nil {
		return nil, err
	}

	macros := AggregateMacros(workouts, meals)

	return macros, nil
}

func AggregateMacros(workouts []model.Workout, meals []model.Meal) []Macro {
	// Need this in maps so we can grab workouts/ meals unique to a normalized date
	workoutMap := workoutMap(workouts)
	mealMap := mealMap(meals)

	macros := make([]Macro, 0)

	// Want macros to be unique by date
	macroMp := make(map[time.Time]*Macro)

	// Workouts grouped by by date, add the caloried burned to date in macros map
	for date, workouts := range workoutMap {
		cals := 0
		for _, w := range workouts {
			cals += w.CaloriesBurned
		}

		macroMp[normalizeDate(date)] = &Macro{Date: normalizeDate(date), Workouts: len(workouts), CalsBurned: cals}
	}

	// For meals, grouped by date, add them to corresponding date in the macros map
	for date, meals := range mealMap {
		cals := 0
		carbs := 0
		fat := 0
		protein := 0
		for _, m := range meals {
			cals += m.Calories
			carbs += m.Carbs
			fat += m.Fat
			protein += m.Protein
		}

		macro, ok := macroMp[normalizeDate(date)]
		if ok {
			macro.CalsIn = cals
			macro.Protein = protein
			macro.Carbs = carbs
			macro.Fat = fat
		} else {
			macroMp[normalizeDate(date)] = &Macro{Date: normalizeDate(date), CalsIn: cals, Protein: protein, Carbs: carbs, Fat: fat}
		}
	}

	// will want to sort this
	for _, macro := range macroMp {
		macros = append(macros, *macro)
	}

	return macros
}

// get a map of workouts indexed by date, with a normalized time ie. 2024-10-12T00:00:00Z
func workoutMap(workouts []model.Workout) map[time.Time][]model.Workout {
	mp := make(map[time.Time][]model.Workout)
	for _, w := range workouts {
		date := normalizeDate(w.Date)
		mp[date] = append(mp[date], w)
	}

	return mp
}

func mealMap(meals []model.Meal) map[time.Time][]model.Meal {
	mp := make(map[time.Time][]model.Meal)
	for _, w := range meals {
		date := normalizeDate(w.Date)
		mp[date] = append(mp[date], w)
	}

	return mp
}

func normalizeDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
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
