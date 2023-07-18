package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	dbName string
}

func (s *SQLite) GetConnection() (*gorm.DB, error) {
	s.dbName = os.Getenv("DB_NAME")

	connection, connError := gorm.Open(sqlite.Open("./test_db"))

	if connError != nil {
		return nil, connError
	}

	return connection, nil
}
