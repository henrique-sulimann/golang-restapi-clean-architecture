package repository

import (
	"context"

	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/entity"
)

type GameRepository interface {
	Create(ctx context.Context, game *entity.Game) (*entity.Game, error)
	FindAll(ctx context.Context) ([]*entity.Game, error)
	FindByID(ctx context.Context, id string) (*entity.Game, error)
	UpdateGameByID(ctx context.Context, id string, game *entity.Game) (*entity.Game, error)
	DeleteGameByID(ctx context.Context, id string) error
}
