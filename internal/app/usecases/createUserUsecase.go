package usecases

import (
	"context"

	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/data/repository"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/entity"
)

type UserUsecase interface {
	Create(ctx context.Context, user *entity.User) error
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByID(ctx context.Context, id string) (*entity.User, error)
	UpdateUserByID(ctx context.Context, id string, user *entity.User) (*entity.User, error)
	DeleteUserByID(ctx context.Context, id string) error
}
type userUsecase struct {
	Repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{Repository: repo}
}
func (g *userUsecase) Create(ctx context.Context, user *entity.User) error {
	err := g.Repository.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (g *userUsecase) FindByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := g.Repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (g *userUsecase) DeleteUserByID(ctx context.Context, id string) error {
	err := g.Repository.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
func (g *userUsecase) FindAll(ctx context.Context) ([]*entity.User, error) {
	user, err := g.Repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (g *userUsecase) UpdateUserByID(ctx context.Context, id string, user *entity.User) (*entity.User, error) {
	existingUser, err := g.Repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	res, err := g.Repository.UpdateUserByID(ctx, id, existingUser)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// func (g *CreateGame) Execute(input GameDTOInput) (GameDTOOutput, error) {
// 	game := entity.NewGame()
// 	game.Name = input.Name
// 	game.Description = input.Description
// 	invalidGame := game.IsValid()
// 	if invalidGame != nil {
// 		return GameDTOOutput{}, invalidGame
// 	}
// 	return g.Create(game)
// }

// func (g *CreateGame) Create(game *entity.User) (GameDTOOutput, error) {
// 	err := g.Repository.Create(game.Name, game.Description)
// 	if err != nil {
// 		return GameDTOOutput{}, err
// 	}
// 	output := GameDTOOutput{
// 		ID:          game.ID,
// 		Name:        game.Name,
// 		Description: game.Description,
// 	}
// 	return output, nil
// }
