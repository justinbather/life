package repository

import (
	"context"

	"github.com/justinbather/life/life-server/db/sqlc"
	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/prettylog"
)

type WorkoutRepository interface {
	CreateWorkout(ctx context.Context, workoutType string) (model.Workout, error)
	GetWorkoutsByType(ctx context.Context, workoutType string) ([]model.Workout, error)
}

type repository struct {
	queries *sqlc.Queries
	logger  *prettylog.Logger
}

func (r *repository) CreateWorkout(ctx context.Context, workoutType string) (model.Workout, error) {
	workout, err := r.queries.CreateWorkout(ctx, workoutType)
	if err != nil {
		r.logger.Errorf("Error in WorkoutRepository.CreateWorkout: %s", err)
		return model.Workout{}, err
	}

	return mapWorkout(workout), nil
}

func (r *repository) GetWorkoutsByType(ctx context.Context, workoutType string) ([]model.Workout, error) {
	records, err := r.queries.GetWorkoutsByType(ctx, workoutType)
	if err != nil {
		r.logger.Errorf("Error in WorkoutRepository.GetWorkoutsByType: %s", err)
		return nil, err
	}

	return mapWorkouts(records), nil
}

func NewWorkoutRepository(db sqlc.DBTX, logger *prettylog.Logger) WorkoutRepository {
	return &repository{queries: sqlc.New(db), logger: logger}
}

func mapWorkouts(w []sqlc.Workout) []model.Workout {
	var workouts []model.Workout
	for _, workout := range w {
		workouts = append(workouts, mapWorkout(workout))
	}

	return workouts
}

func mapWorkout(w sqlc.Workout) model.Workout {
	return model.Workout{Id: int(w.ID), Type: w.Type}
}
