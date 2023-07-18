package database

import (
	"log"

	"github.com/edmartt/characters-crud/internal/api/characters/models"
)

type Migrations struct {
	DB SQLDB
}

func (m Migrations) MigrateData() {
	connection, connError := m.DB.GetConnection()

	if connError != nil {
		log.Printf("%s", connError.Error())
	}

	log.Println("====MIGRATING====")

	connection.AutoMigrate(&models.Character{})
}

func Init() {
	migrations := Migrations{
		DB: &SQLite{dbName: "test"},
	}

	migrations.MigrateData()
}
