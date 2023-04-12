package repository

import (
	"context"

	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByID(ctx context.Context, id string) (*entity.User, error)
	UpdateUserByID(ctx context.Context, id string, user *entity.User) (*entity.User, error)
	DeleteUserByID(ctx context.Context, id string) error
}
