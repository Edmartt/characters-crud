package database

import "gorm.io/gorm"

type SQLDB interface {
	GetConnection() (*gorm.DB, error)
}
