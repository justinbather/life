package service

import (
	"context"

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
	GetWorkoutsByType(ctx context.Context, workoutType string) ([]model.Workout, error)
	GetAllWorkouts(ctx context.Context) ([]model.Workout, error)
}

func NewWorkoutService(repository repository.WorkoutRepository, logger *prettylog.Logger) WorkoutService {
	return &service{repository: repository, logger: logger}
}

func (s *service) CreateWorkout(ctx context.Context, workout model.Workout) (model.Workout, error) {
	workout, err := s.repository.CreateWorkout(ctx, workout)
	if err != nil {
		s.logger.Errorf("Error creating workout with type: %s. Err: %s", workout, err)
		return model.Workout{}, err
	}

	s.logger.Infof("Created new workout")
	return workout, nil
}

func (s *service) GetWorkoutsByType(ctx context.Context, workoutType string) ([]model.Workout, error) {
	workouts, err := s.repository.GetWorkoutsByType(ctx, workoutType)
	if err != nil {
		s.logger.Errorf("Error getting workouts with type: %s. Err: %s", workoutType, err)
		return nil, err
	}

	s.logger.Infof("Fetched %d workouts", len(workouts))
	return workouts, nil
}

func (s *service) GetAllWorkouts(ctx context.Context) ([]model.Workout, error) {
	workouts, err := s.repository.GetAllWorkouts(ctx)
	if err != nil {
		s.logger.Errorf("Error getting all workouts: %s", err)
		return nil, err
	}

	s.logger.Infof("Fetched %d workouts", len(workouts))
	return workouts, nil
}
