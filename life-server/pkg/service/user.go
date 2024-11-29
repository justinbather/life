package service

import (
	"context"
	"fmt"

	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/life/life-server/pkg/repository"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
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
	hash, err := hashPassword(password)
	if err != nil {
		return model.User{}, fmt.Errorf("Error creating user")
	}

	user, err := s.repository.CreateUser(ctx, model.User{Id: ulid, Username: username, Password: hash})
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
func (s *userService) GetUserById(ctx context.Context, id string) (model.User, error) {
	return model.User{}, nil
}
func (s *userService) GetUserByUsernameAndPass(ctx context.Context, username, password string) (model.User, error) {
	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		return model.User{}, err
	}

	ok := checkPasswordHash(password, user.Password)
	if !ok {
		return model.User{}, fmt.Errorf("No user found")
	}

	return user, nil
}

func hashPassword(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), 14)
	return string(bytes), err
}

func checkPasswordHash(secret, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err == nil
}
