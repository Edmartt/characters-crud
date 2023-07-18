package data

import (
	"log"

	"github.com/edmartt/characters-crud/internal/api/characters/models"
	"github.com/edmartt/characters-crud/internal/database"
)

var dbConnectionObject database.SQLDB

type CharacterDataAccessImplementation struct {
	character *models.Character
}

func init() {
	dbConnectionObject = &database.SQLite{}
}

func (c CharacterDataAccessImplementation) CreateCharacter(character models.Character) (string, error) {

	connection, connErr := dbConnectionObject.GetConnection()

	if connErr != nil {
		log.Println("ERROR: ", connErr.Error())
		return "", nil
	}

	connection.Create(&character)

	return "character(s) created", nil
}

func (c CharacterDataAccessImplementation) GetCharacter(id string) (*models.Character, error) {
	connection, connErr := dbConnectionObject.GetConnection()

	if connErr != nil {
		log.Println("CONNECTION ERROR: ", connErr.Error())
		return nil, connErr
	}

	connection.Where("id = ?", id).First(&c.character)

	if c.character.ID == "" {
		return c.character, nil
	}

	return c.character, nil
}

func (c CharacterDataAccessImplementation) GetCharacters() ([]models.Character, error) {
	connection, connErr := dbConnectionObject.GetConnection()

	if connErr != nil {
		log.Println("ERROR: ", connErr.Error())
		return []models.Character{}, connErr
	}

	var records []models.Character
	err := connection.Find(&records).Error

	if err != nil {
		log.Println("ERROR: ", err.Error())
	}

	return records, nil
}
