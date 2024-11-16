package service

import "github.com/justinbather/prettylog"

type WorkoutService struct {
	Logger *prettylog.Logger
}

func NewWorkoutService() *WorkoutService {
	return &WorkoutService{Logger: prettylog.New()}
}
