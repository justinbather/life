// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateMeal(ctx context.Context, arg CreateMealParams) (CreateMealRow, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (CreateWorkoutRow, error)
	GetAllWorkouts(ctx context.Context, userID string) ([]Workout, error)
	GetMealById(ctx context.Context, id int32) (Meal, error)
	GetMealsByType(ctx context.Context, arg GetMealsByTypeParams) ([]Meal, error)
	GetMealsFromDateRange(ctx context.Context, arg GetMealsFromDateRangeParams) ([]Meal, error)
	GetUserById(ctx context.Context, id string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetWorkoutsByType(ctx context.Context, arg GetWorkoutsByTypeParams) ([]Workout, error)
	GetWorkoutsFromDateRange(ctx context.Context, arg GetWorkoutsFromDateRangeParams) ([]Workout, error)
}

var _ Querier = (*Queries)(nil)
