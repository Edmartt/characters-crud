package routes

import (
	"github.com/edmartt/http_test/internal/api/characters/handlers"
	"github.com/gin-gonic/gin"
)

func SetCharactersRoutes(router *gin.RouterGroup) {

	handlersObject := handlers.CharactersHandler{}
	router.POST("/characters", handlersObject.PostCharacter)
	router.GET("/characters", handlersObject.GetCharacters)
	router.GET("/characters/:id", handlersObject.GetCharacter)
}
