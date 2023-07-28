package tests

import (
	"github.com/edmartt/characters-crud/internal/api/characters/models"
)

// MockCharacter contains http response from http request to characters get method. This embbeded struct help us to get the needed data cause the data is inside an object called response
type MockCharacter struct {
	Response struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Serie     string `json:"serie"`
		Superhero bool   `json:"superhero"`
		Gender    string `json:"gender"`
	} `json:"response"`
}

// MockDataAccess is for character data access layer
type MockDataAccess struct{}

func (m MockDataAccess) GetCharacter(id string) (*models.Character, error) {
	return &models.Character{
		ID:        "123",
		Name:      "mocked char",
		Serie:     "serie",
		Superhero: true,
		Gender:    "male",
	}, nil
}

func (m MockDataAccess) GetCharacters() ([]models.Character, error) {
	return []models.Character{}, nil
}

func (m MockDataAccess) CreateCharacter(character models.Character) (string, error) {
	return "1234", nil
}
