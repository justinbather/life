package service

import (
	"fmt"
	"time"

	"github.com/justinbather/life/cli/model"
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

	/*
		given an array of workouts and an array of meals
		I want an array of macros, which has aggregate data from both resources

		since both resources are date based and we want the data aggregated by date
		we could normalize dates and create a map of workout arrays and a map of meals, indexed by dates

		then to build our macro array, we need to loop over the keys for each, and create a macro struct for each unique date
	*/

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

	macrosOld := MacroMap{"calsBurned": totCalsBurned, "calsIn": totCalsIn, "protein": totProtein, "carbs": totCarbs, "fat": totFat}

	return macrosOld, nil
}

func AggregateMacros(workouts []model.Workout, meals []model.Meal) []Macro {

	workoutMap := workoutMap(workouts)
	mealMap := mealMap(meals)

	macros := make([]Macro, 0)

	// theres some weird stuff going on here with the logic of how we choose which list to base our map indices off of
	// just because there may be 4 unique dates for meals, and 4 for workouts, but they may be different dates
	// we can just loop over the workout map, add the relevent data to the date, then do the same after

	// gather workout data
	macroMp := make(map[time.Time]*Macro)
	// for each unique date in workouts, add all workout data and create entry in map
	for date, workouts := range workoutMap {
		cals := 0
		for _, w := range workouts {
			cals += w.CaloriesBurned
		}

		macroMp[normalizeDate(date)] = &Macro{Date: normalizeDate(date), Workouts: len(workouts), CalsBurned: cals}
	}

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
