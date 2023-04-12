package main

import (
	"context"

	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/app/usecases"
	m "github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/infra/databases/mongodb"
	h "github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/infra/http"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/infra/http/handler"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/pkg/mongodb"
)

func main() {
	mclient := mongodb.StartMongo()
	defer mclient.Disconnect(context.Background())
	dbName := "gamesdb"
	gameCollectionName := "games"
	gameRepository := m.NewGameRepositoryDB(mclient, dbName, gameCollectionName)
	gameUsecase := usecases.NewGameUsecase(gameRepository)
	gameHandler := handler.NewGameHandler(gameUsecase)
	server := h.NewServer()
	server.Run(gameHandler)
}
