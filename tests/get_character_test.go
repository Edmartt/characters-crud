package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edmartt/characters-crud/internal/api/characters/data"
	"github.com/edmartt/characters-crud/internal/api/characters/handlers"
	"github.com/gin-gonic/gin"
)

func init() {
	data.DataAccess = MockDataAccess{}
}

func TestGetCharacterHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := &handlers.CharactersHandler{}

	router := gin.Default()

	router.GET("/characters/:id", handler.GetCharacter)

	req, err := http.NewRequest("GET", "/characters/123", nil)

	if err != nil {
		t.Logf("error: %v", err.Error())
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {

		t.Logf("error: %v", err.Error())
		t.Errorf("expected code %d but got %d", http.StatusOK, resp.Code)
	}

	character := &MockCharacter{}

	err = json.Unmarshal(resp.Body.Bytes(), &character)

	if err != nil {
		t.Errorf("Decoding JSON response error: %v", err)
	}

	expectedID := "123"
	expectedName := "mocked char"

	if character.Response.ID != expectedID {
		t.Errorf("expected id %s and got %s", expectedID, character.Response.ID)
	}

	if character.Response.Name != expectedName {
		t.Errorf("expected %s and got %s", expectedName, character.Response.Name)
	}

}
