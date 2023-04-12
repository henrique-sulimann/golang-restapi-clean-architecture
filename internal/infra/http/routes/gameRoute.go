package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/infra/http/handler"
)

func GameRoutes(router *gin.Engine, handler *handler.GameHandler) *gin.Engine {
	games := router.Group("api/v1/games")
	{
		games.GET("/:id", handler.FindByID)
		games.GET("/", handler.FindAllGames)
		games.POST("/", handler.CreateGame)
		games.PUT("/:id", handler.UpdateGameByID)
		games.DELETE("/:id", handler.DeleteGameByID)
	}
	return router
}
