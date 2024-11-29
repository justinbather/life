package service

import (
	"context"

	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/life/life-server/pkg/repository"
	"github.com/oklog/ulid/v2"
)

type UserService interface {
	CreateUser(ctx context.Context, username, password string) (model.User, error)
	GetUserById(ctx context.Context, id string) (model.User, error)
	GetUserByUsernameAndPass(ctx context.Context, username, password string) (model.User, error)
}

func NewUserService(respository repository.UserRespository) UserService {
	return &userService{repository: respository}
}

type userService struct {
	repository repository.UserRespository
}

func (s *userService) CreateUser(ctx context.Context, username, password string) (model.User, error) {
	ulid := ulid.Make().String()

	user, err := s.repository.CreateUser(ctx, model.User{Id: ulid, Username: username, Password: password})
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
func (s *userService) GetUserById(ctx context.Context, id string) (model.User, error) {
	return model.User{}, nil
}
func (s *userService) GetUserByUsernameAndPass(ctx context.Context, username, password string) (model.User, error) {
	user, err := s.repository.GetUserByUsernameAndPass(ctx, username, password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
