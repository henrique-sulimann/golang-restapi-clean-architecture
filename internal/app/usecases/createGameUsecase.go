package usecases

import (
	"context"

	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/data/dto"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/data/repository"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/entity"
)

type GameUsecase interface {
	Create(ctx context.Context, input dto.GameDTOInput) (*dto.GameDTOOutput, error)
	FindAll(ctx context.Context) ([]dto.GameDTOOutput, error)
	FindByID(ctx context.Context, id string) (*dto.GameDTOOutput, error)
	UpdateGameByID(ctx context.Context, id string, input *dto.GameDTOInput) (*dto.GameDTOOutput, error)
	DeleteGameByID(ctx context.Context, id string) error
}

type gameUsecase struct {
	Repository repository.GameRepository
}

func NewGameUsecase(repo repository.GameRepository) GameUsecase {
	return &gameUsecase{Repository: repo}
}
func (g *gameUsecase) Create(ctx context.Context, input dto.GameDTOInput) (*dto.GameDTOOutput, error) {
	game := entity.NewGame()
	game.Name = input.Name
	game.Description = input.Description
	err := game.IsValid()
	if err != nil {
		return nil, err
	}
	res, err := g.Repository.Create(ctx, game)
	if err != nil {
		return nil, err
	}
	output := dto.GameDTOOutput{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
	return &output, nil
}
func (g *gameUsecase) FindByID(ctx context.Context, id string) (*dto.GameDTOOutput, error) {
	game, err := g.Repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	output := dto.GameDTOOutput{
		ID:          game.ID,
		Name:        game.Name,
		Description: game.Description,
		CreatedAt:   game.CreatedAt,
		UpdatedAt:   game.UpdatedAt,
	}
	return &output, nil
}
func (g *gameUsecase) DeleteGameByID(ctx context.Context, id string) error {
	err := g.Repository.DeleteGameByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (g *gameUsecase) FindAll(ctx context.Context) ([]dto.GameDTOOutput, error) {
	games, err := g.Repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	output := []dto.GameDTOOutput{}
	for _, game := range games {
		output = append(output, dto.GameDTOOutput{
			ID:          game.ID,
			Name:        game.Name,
			Description: game.Description,
			CreatedAt:   game.CreatedAt,
			UpdatedAt:   game.UpdatedAt,
		})
	}
	return output, nil
}
func (g *gameUsecase) UpdateGameByID(ctx context.Context, id string, game *dto.GameDTOInput) (*dto.GameDTOOutput, error) {
	existingGame, err := g.Repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if game.Name != "" {
		existingGame.Name = game.Name
	}
	if game.Description != "" {
		existingGame.Description = game.Description
	}
	res, err := g.Repository.UpdateGameByID(ctx, id, existingGame)
	if err != nil {
		return nil, err
	}
	output := dto.GameDTOOutput{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}
	return &output, nil
}
