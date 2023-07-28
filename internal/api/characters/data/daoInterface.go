package data

import "github.com/edmartt/characters-crud/internal/api/characters/models"

var DataAccess CharacterDataAccessInterface

type CharacterDataAccessInterface interface {
	CreateCharacter(character models.Character) (string, error)
	GetCharacters() ([]models.Character, error)
	GetCharacter(id string) (*models.Character, error)
}
