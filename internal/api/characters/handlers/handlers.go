package handlers

import (
	"log"
	"net/http"

	"github.com/edmartt/characters-crud/internal/api/characters/data"
	"github.com/edmartt/characters-crud/internal/api/characters/models"
	"github.com/edmartt/characters-crud/internal/api/characters/responses"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CharactersHandler struct {
}

var dataAccess data.CharacterDataAccessInterface

func init() {
	dataAccess = data.CharacterDataAccessImplementation{}
}

func (c CharactersHandler) PostCharacter(context *gin.Context) {

	modelObject := models.Character{}
	modelObject.ID = uuid.NewString()

	err := context.BindJSON(&modelObject)

	if err != nil {

		badResponse := responses.HttpBad{
			Response: err.Error(),
		}

		context.JSON(http.StatusBadRequest, badResponse)
		return
	}

	accessOBject, err := dataAccess.CreateCharacter(modelObject)

	if err != nil {
		log.Println("ERROR: ", err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	created := responses.Created{
		CharacterData: accessOBject,
	}

	context.JSON(http.StatusCreated, created)
}

func (c CharactersHandler) GetCharacter(context *gin.Context) {

	id := context.Param("id")

	accessObject, err := dataAccess.GetCharacter(id)

	if err != nil {
		log.Println("DATA ERROR", err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response := responses.HttpOK{
		Response: *accessObject,
	}

	context.JSON(http.StatusOK, response)
}

func (c CharactersHandler) GetCharacters(context *gin.Context) {

	existantCharacters, err := dataAccess.GetCharacters()

	if err != nil {
		log.Println("ERROR: ", err.Error())
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	okResponse := responses.HttpOKArray{
		Response: &existantCharacters,
	}

	context.JSON(http.StatusOK, okResponse)
}
