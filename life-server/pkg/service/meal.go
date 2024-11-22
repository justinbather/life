package service

import (
	"context"
	"time"

	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/life/life-server/pkg/repository"
	"github.com/justinbather/prettylog"
)

type MealService interface {
	CreateMeal(ctx context.Context, meal model.Meal) (model.Meal, error)
	GetMealById(ctx context.Context, id int) (model.Meal, error)
	GetMealsFromDateRange(ctx context.Context, user string, from time.Time, to time.Time) ([]model.Meal, error)
}

type mealService struct {
	logger     *prettylog.Logger
	repository repository.MealRepository
}

func NewMealService(repository repository.MealRepository, logger *prettylog.Logger) MealService {
	return &mealService{logger: logger, repository: repository}
}

func (s *mealService) CreateMeal(ctx context.Context, meal model.Meal) (model.Meal, error) {
	meal, err := s.repository.CreateMeal(ctx, meal)
	if err != nil {
		return model.Meal{}, err
	}

	return meal, nil
}

func (s *mealService) GetMealById(ctx context.Context, id int) (model.Meal, error) {
	meal, err := s.repository.GetMealById(ctx, id)
	if err != nil {
		return model.Meal{}, err
	}

	return meal, nil
}

func (s *mealService) GetMealsFromDateRange(ctx context.Context, user string, from time.Time, to time.Time) ([]model.Meal, error) {
	s.logger.Infof("Fetching meals between %s and %s", from, to)
	meals, err := s.repository.GetMealsFromDateRange(ctx, from, to)
	if err != nil {
		return nil, err
	}

	return meals, nil
}
