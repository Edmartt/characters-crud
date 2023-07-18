package data

import "github.com/edmartt/http_test/internal/api/characters/models"

type CharacterDataAccessInterface interface {
	CreateCharacter(character models.Character) (string, error)
	GetCharacters() ([]models.Character, error)
	GetCharacter(id string) (*models.Character, error)
}
