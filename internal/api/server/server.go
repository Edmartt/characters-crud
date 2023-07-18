package server

import (
	routers "github.com/edmartt/characters-crud/internal/api/characters/routes"
	"github.com/edmartt/characters-crud/internal/api/server/ping"
	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")

	ping.SetServerHealthRoute(apiGroup)
	routers.SetCharactersRoutes(apiGroup)

	return router
}

func Init(port string) {
	router := setRouter()
	router.Run(":" + port)
}
