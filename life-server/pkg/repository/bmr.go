package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/justinbather/life/life-server/db/sqlc"
	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/prettylog"
)

type BmrRepository interface {
	CreateBmr(ctx context.Context, meal model.Bmr) (model.Bmr, error)
	GetBmrById(ctx context.Context, id int) (model.Bmr, error)
	GetBmrFromDateRange(ctx context.Context, user string, from time.Time, to time.Time) ([]model.Bmr, error)
}

type bmrRepository struct {
	queries *sqlc.Queries
	logger  *prettylog.Logger
}

func NewBmrRepository(db sqlc.DBTX, logger *prettylog.Logger) BmrRepository {
	return &bmrRepository{queries: sqlc.New(db), logger: logger}
}

func (r *bmrRepository) CreateBmr(ctx context.Context, bmr model.Bmr) (model.Bmr, error) {
	date := pgtype.Timestamp{Time: bmr.CreatedAt, Valid: true}
	record, err := r.queries.CreateBmr(ctx, sqlc.CreateBmrParams{
		UserID:        bmr.UserId,
		CreatedAt:     date,
		TotalCalories: int32(bmr.TotalCalories),
		NumWorkouts:   int32(bmr.NumWorkouts)})

	if err != nil {
		r.logger.Errorf("Error saving meal: %s", err)
		return model.Bmr{}, nil
	}
	return mapBmr(record), nil
}

func (r *bmrRepository) GetBmrById(ctx context.Context, id int) (model.Bmr, error) {
	record, err := r.queries.GetBmrById(ctx, int32(id))
	if err != nil {
		r.logger.Errorf("Error getting meal by id=%d. Err: %s", id, err)
		return model.Bmr{}, err
	}
	return mapBmr(record), nil
}

func (r *bmrRepository) GetBmrFromDateRange(ctx context.Context, user string, from time.Time, to time.Time) ([]model.Bmr, error) {
	r.logger.Infof("Fetching meals from %s to %s", from, to)
	records, err := r.queries.GetBmrFromDateRange(ctx, sqlc.GetBmrFromDateRangeParams{
		UserID:      user,
		CreatedAt:   mapDate(from),
		CreatedAt_2: mapDate(to)})

	if err != nil {
		r.logger.Errorf("Error getting meals from date range: ", err)
		return nil, err
	}

	return mapBmrs(records), nil
}

func mapBmrs(records []sqlc.Bmr) []model.Bmr {
	var bmrs []model.Bmr
	for _, bmr := range records {
		bmrs = append(bmrs, mapBmr(bmr))
	}
	return bmrs
}

func mapBmr(record sqlc.Bmr) model.Bmr {
	return model.Bmr{
		Id:            int(record.ID),
		UserId:        record.UserID,
		TotalCalories: int(record.TotalCalories),
		NumWorkouts:   int(record.NumWorkouts),
		CreatedAt:     record.CreatedAt.Time}
}
