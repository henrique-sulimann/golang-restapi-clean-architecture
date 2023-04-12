package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/infra/http/handler"
	"github.com/henrique-sulimann/golang-restapi-clean-architecture/internal/infra/http/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "3000",
		server: gin.Default(),
	}
}

func (s *Server) Run(handler *handler.GameHandler) {
	router := routes.GameRoutes(s.server, handler)

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
