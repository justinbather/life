package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/justinbather/life/life-server/db/sqlc"
	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/prettylog"
)

type MealRepository interface {
	CreateMeal(ctx context.Context, meal model.Meal) (model.Meal, error)
	GetMealById(ctx context.Context, id int) (model.Meal, error)
	GetMealsFromDateRange(ctx context.Context, user string, from time.Time, to time.Time) ([]model.Meal, error)
}

type mealRepository struct {
	queries *sqlc.Queries
	logger  *prettylog.Logger
}

func NewMealRepository(db sqlc.DBTX, logger *prettylog.Logger) MealRepository {
	return &mealRepository{queries: sqlc.New(db), logger: logger}
}

func (r *mealRepository) CreateMeal(ctx context.Context, meal model.Meal) (model.Meal, error) {
	date := pgtype.Timestamp{Time: meal.Date, Valid: true}
	record, err := r.queries.CreateMeal(ctx, sqlc.CreateMealParams{Type: meal.Type, Calories: int32(meal.Calories), Protein: int32(meal.Protein), Carbs: int32(meal.Carbs), Fat: int32(meal.Fat), Description: &meal.Description, Date: date})
	if err != nil {
		r.logger.Errorf("Error saving meal: %s", err)
		return model.Meal{}, nil
	}
	return mapMeal(record), nil
}

func (r *mealRepository) GetMealById(ctx context.Context, id int) (model.Meal, error) {
	record, err := r.queries.GetMealById(ctx, int32(id))
	if err != nil {
		r.logger.Errorf("Error getting meal by id=%d. Err: %s", id, err)
	}
	return mapMeal(record), nil
}

func (r *mealRepository) GetMealsFromDateRange(ctx context.Context, user string, from time.Time, to time.Time) ([]model.Meal, error) {
	records, err := r.queries.GetMealsFromDateRange(ctx, sqlc.GetMealsFromDateRangeParams{Username: user, Date: mapDate(from), Date_2: mapDate(to)})
	if err != nil {
		r.logger.Errorf("Error getting meals from date range: ", err)
		return nil, err
	}
	return mapMeals(records), nil
}

func mapDate(d time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{Time: d, Valid: true}
}

func mapMeals(m []sqlc.Meal) []model.Meal {
	var meals []model.Meal
	for _, meal := range m {
		meals = append(meals, mapMeal(meal))
	}
	return meals
}

func mapMeal(m sqlc.Meal) model.Meal {
	return model.Meal{Id: int(m.ID), Type: m.Type, Calories: int(m.Calories), Protein: int(m.Protein), Carbs: int(m.Carbs), Fat: int(m.Fat), Description: *m.Description, Date: m.Date.Time}
}
