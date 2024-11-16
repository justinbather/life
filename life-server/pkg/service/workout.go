package service

import (
	"context"
	"fmt"

	"github.com/justinbather/life/life-server/db/sqlc"
	"github.com/justinbather/life/life-server/pkg/repository"
	"github.com/justinbather/prettylog"
)

type service struct {
	logger     *prettylog.Logger
	repository repository.WorkoutRepository
}

type WorkoutService interface {
	CreateWorkout(ctx context.Context, workoutType string) (sqlc.Workout, error)
	GetWorkoutsByType(ctx context.Context, workoutType string) ([]sqlc.Workout, error)
}

func NewWorkoutService(repository repository.WorkoutRepository) WorkoutService {
	return &service{logger: prettylog.New(), repository: repository}
}

func (s *service) CreateWorkout(ctx context.Context, workoutType string) (sqlc.Workout, error) {
	workout, err := s.repository.CreateWorkout(ctx, workoutType)
	if err != nil {
		fmt.Printf("Error creating workout with type: %s. Err: %s", workoutType, err)
		return sqlc.Workout{}, err
	}

	return workout, nil
}

func (s *service) GetWorkoutsByType(ctx context.Context, workoutType string) ([]sqlc.Workout, error) {
	workouts, err := s.repository.GetWorkoutsByType(ctx, workoutType)
	if err != nil {
		fmt.Printf("Error getting workouts with type: %s. Err: %s", workoutType, err)
		return nil, err
	}

	return workouts, nil
}
