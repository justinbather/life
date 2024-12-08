package service

import (
	"context"
	"time"

	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/life/life-server/pkg/repository"
)

type BmrService interface {
	CreateBmr(ctx context.Context, bmr model.Bmr) (model.Bmr, error)
	GetBmrsFromDateRange(ctx context.Context, user string, from time.Time, to time.Time) ([]model.Bmr, error)
}

type bmrService struct {
	repository repository.BmrRepository
}

func NewBmrService(repository repository.BmrRepository) BmrService {
	return &bmrService{repository: repository}
}

func (s *bmrService) CreateBmr(ctx context.Context, bmr model.Bmr) (model.Bmr, error) {
	return s.repository.CreateBmr(ctx, bmr)

}

func (s *bmrService) GetBmrsFromDateRange(ctx context.Context, user string, from time.Time, to time.Time) ([]model.Bmr, error) {
	return s.repository.GetBmrFromDateRange(ctx, user, from, to)
}
