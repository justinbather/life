// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateMeal(ctx context.Context, arg CreateMealParams) (Meal, error)
	CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (Workout, error)
	GetAllWorkouts(ctx context.Context, username string) ([]Workout, error)
	GetMealById(ctx context.Context, id int32) (Meal, error)
	GetMealsByType(ctx context.Context, arg GetMealsByTypeParams) ([]Meal, error)
	GetMealsFromDateRange(ctx context.Context, arg GetMealsFromDateRangeParams) ([]Meal, error)
	GetWorkoutsByType(ctx context.Context, arg GetWorkoutsByTypeParams) ([]Workout, error)
	GetWorkoutsFromDateRange(ctx context.Context, arg GetWorkoutsFromDateRangeParams) ([]Workout, error)
}

var _ Querier = (*Queries)(nil)
