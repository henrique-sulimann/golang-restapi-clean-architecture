package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/infra/http/handler"
)

func UserRoutes(router *gin.Engine, handler *handler.GameHandler) *gin.Engine {
	users := router.Group("api/v1/users")
	{
		users.GET("/:id", handler.FindByID)
		users.GET("/", handler.FindAllGames)
		users.POST("/", handler.CreateGame)
		users.PUT("/:id", handler.UpdateGameByID)
		users.DELETE("/:id", handler.DeleteGameByID)
	}
	return router
}
