package service

import (
	"context"
	"time"

	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/life/life-server/pkg/repository"
	"github.com/justinbather/prettylog"
)

type service struct {
	logger     *prettylog.Logger
	repository repository.WorkoutRepository
}

type WorkoutService interface {
	CreateWorkout(ctx context.Context, workout model.Workout) (model.Workout, error)
	GetWorkoutsByType(ctx context.Context, user string, workoutType string) ([]model.Workout, error)
	GetAllWorkouts(ctx context.Context, user string) ([]model.Workout, error)
	GetWorkoutsFromDateRange(ctx context.Context, user string, from, to time.Time) ([]model.Workout, error)
}

func NewWorkoutService(repository repository.WorkoutRepository, logger *prettylog.Logger) WorkoutService {
	return &service{repository: repository, logger: logger}
}

func (s *service) CreateWorkout(ctx context.Context, workout model.Workout) (model.Workout, error) {
	workout, err := s.repository.CreateWorkout(ctx, workout)
	if err != nil {
		return model.Workout{}, err
	}

	s.logger.Infof("Created new workout")
	return workout, nil
}

func (s *service) GetWorkoutsByType(ctx context.Context, user, workoutType string) ([]model.Workout, error) {
	workouts, err := s.repository.GetWorkoutsByType(ctx, user, workoutType)
	if err != nil {
		return nil, err
	}

	s.logger.Infof("Fetched %d workouts", len(workouts))
	return workouts, nil
}

func (s *service) GetAllWorkouts(ctx context.Context, user string) ([]model.Workout, error) {
	workouts, err := s.repository.GetAllWorkouts(ctx, user)
	if err != nil {
		return nil, err
	}

	s.logger.Infof("Fetched %d workouts", len(workouts))
	return workouts, nil
}

func (s *service) GetWorkoutsFromDateRange(ctx context.Context, user string, from, to time.Time) ([]model.Workout, error) {
	workouts, err := s.repository.GetWorkoutsFromDateRange(ctx, user, from, to)
	if err != nil {
		return nil, err
	}

	s.logger.Infof("Fetched %d workouts", len(workouts))
	return workouts, nil
}
