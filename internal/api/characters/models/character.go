package models

type Character struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Serie     string `json:"serie"`
	Superhero bool   `json:"superhero"`
	Gender    string `json:"gender"`
}
