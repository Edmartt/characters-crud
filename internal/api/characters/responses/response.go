package responses

import "github.com/edmartt/http_test/internal/api/characters/models"

type HttpOKArray struct {
	Response *[]models.Character `json:"response"`
}

type HttpOK struct {
	Response models.Character `json:"response"`
}

type HttpBad struct {
	Response string `json:"bad_request"`
}

type Created struct {
	CharacterData string `json:"response"`
}
