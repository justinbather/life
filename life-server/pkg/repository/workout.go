package repository

import (
	"context"

	"github.com/justinbather/life/life-server/db/sqlc"
	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/prettylog"
)

type WorkoutRepository interface {
	CreateWorkout(ctx context.Context, workout model.Workout) (model.Workout, error)
	GetWorkoutsByType(ctx context.Context, workoutType string) ([]model.Workout, error)
	GetAllWorkouts(ctx context.Context) ([]model.Workout, error)
}

type workoutRepository struct {
	queries *sqlc.Queries
	logger  *prettylog.Logger
}

func (r *workoutRepository) CreateWorkout(ctx context.Context, workout model.Workout) (model.Workout, error) {
	res, err := r.queries.CreateWorkout(ctx, sqlc.CreateWorkoutParams{Type: workout.Type, Duration: int32(workout.Duration), CaloriesBurned: int32(workout.CaloriesBurned), Workload: int32(workout.Workload), Description: &workout.Description})
	if err != nil {
		r.logger.Errorf("Error in WorkoutRepository.CreateWorkout: %s", err)
		return model.Workout{}, err
	}

	return mapWorkout(res), nil
}

func (r *workoutRepository) GetWorkoutsByType(ctx context.Context, workoutType string) ([]model.Workout, error) {
	records, err := r.queries.GetWorkoutsByType(ctx, workoutType)
	if err != nil {
		r.logger.Errorf("Error in WorkoutRepository.GetWorkoutsByType: %s", err)
		return nil, err
	}

	return mapWorkouts(records), nil
}

func (r *workoutRepository) GetAllWorkouts(ctx context.Context) ([]model.Workout, error) {
	records, err := r.queries.GetAllWorkouts(ctx)
	if err != nil {
		r.logger.Errorf("Error in WorkoutRepository.GetAllWorkouts: %s", err)
		return nil, err
	}

	return mapWorkouts(records), nil
}

func NewWorkoutRepository(db sqlc.DBTX, logger *prettylog.Logger) WorkoutRepository {
	return &workoutRepository{queries: sqlc.New(db), logger: logger}
}

func mapWorkouts(w []sqlc.Workout) []model.Workout {
	var workouts []model.Workout
	for _, workout := range w {
		workouts = append(workouts, mapWorkout(workout))
	}

	return workouts
}

func mapWorkout(w sqlc.Workout) model.Workout {
	return model.Workout{Id: int(w.ID), Type: w.Type, Duration: int(w.Duration), CreatedAt: w.CreatedAt.Time, Workload: int(w.Workload), CaloriesBurned: int(w.CaloriesBurned), Description: *w.Description}
}
