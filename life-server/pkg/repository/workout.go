package repository

import (
	"context"
	"fmt"

	"github.com/justinbather/life/life-server/db/sqlc"
)

type WorkoutRepository interface {
	CreateWorkout(ctx context.Context, workoutType string) (sqlc.Workout, error)
	GetWorkoutsByType(ctx context.Context, workoutType string) ([]sqlc.Workout, error)
}

type repository struct {
	queries *sqlc.Queries
}

func (r *repository) CreateWorkout(ctx context.Context, workoutType string) (sqlc.Workout, error) {
	workout, err := r.queries.CreateWorkout(ctx, workoutType)
	if err != nil {
		fmt.Printf("Error in WorkoutRepository.CreateWorkout: %s", err)
		return sqlc.Workout{}, err
	}

	return workout, nil
}

func (r *repository) GetWorkoutsByType(ctx context.Context, workoutType string) ([]sqlc.Workout, error) {
	workouts, err := r.queries.GetWorkoutsByType(ctx, workoutType)
	if err != nil {
		fmt.Printf("Error in WorkoutRepository.GetWorkoutsByType: %s", err)
		return nil, err
	}

	return workouts, nil
}

func NewWorkoutRepository(db sqlc.DBTX) WorkoutRepository {
	return &repository{queries: sqlc.New(db)}
}
