package repository

import (
	"context"

	"github.com/justinbather/life/life-server/db/sqlc"
	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/prettylog"
)

type UserRespository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserById(ctx context.Context, id string) (model.User, error)
	GetUserByUsernameAndPass(ctx context.Context, username, password string) (model.User, error)
}

type userRepository struct {
	queries *sqlc.Queries
	logger  *prettylog.Logger
}

func NewUserRepository(db sqlc.DBTX, logger *prettylog.Logger) UserRespository {
	return &userRepository{queries: sqlc.New(db), logger: logger}
}

func (r *userRepository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	record, err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{ID: user.Id, Username: user.Username, Password: user.Password})
	if err != nil {
		return model.User{}, err
	}

	return mapUser(record), nil
}

func (r *userRepository) GetUserById(ctx context.Context, id string) (model.User, error) {
	return model.User{}, nil
}

func (r *userRepository) GetUserByUsernameAndPass(ctx context.Context, username, password string) (model.User, error) {
	record, err := r.queries.GetUserByUsernameAndPass(ctx, sqlc.GetUserByUsernameAndPassParams{Username: username, Password: password})
	if err != nil {
		return model.User{}, err
	}

	return mapUser(record), nil
}

func mapUser(u sqlc.User) model.User {
	return model.User{
		Id:       u.ID,
		Username: u.Username,
		Password: u.Password,
	}
}
