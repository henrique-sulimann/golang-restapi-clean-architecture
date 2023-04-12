package handler

import (
	"context"

	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/data/dto"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/usecases"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gameUsecase usecases.GameUsecase
}

func NewGameHandler(gameUsecase usecases.GameUsecase) *GameHandler {
	return &GameHandler{
		gameUsecase: gameUsecase,
	}
}

func (h *GameHandler) CreateGame(c *gin.Context) {
	var input dto.GameDTOInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid json",
		})
		return
	}
	game, err := h.gameUsecase.Create(context.Background(), input)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error to create game: " + err.Error(),
		})
		return
	}
	c.JSON(200, game)
}
func (h *GameHandler) UpdateGameByID(c *gin.Context) {
	id := c.Param("id")
	var input dto.GameDTOInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid json",
		})
		return
	}
	res, err := h.gameUsecase.UpdateGameByID(context.Background(), id, &input)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error to update game: " + err.Error(),
		})
		return
	}
	c.JSON(200, res)
}
func (h *GameHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	game, err := h.gameUsecase.FindByID(context.Background(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error to find game by id: " + err.Error(),
		})
		return
	}
	c.JSON(200, game)
}
func (h *GameHandler) DeleteGameByID(c *gin.Context) {
	id := c.Param("id")
	err := h.gameUsecase.DeleteGameByID(context.Background(), id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error to delete game by id: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Game deleted successfully",
	})
}
func (h *GameHandler) FindAllGames(c *gin.Context) {
	games, err := h.gameUsecase.FindAll(context.Background())
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error to find games: " + err.Error(),
		})
		return
	}
	c.JSON(200, games)
}
