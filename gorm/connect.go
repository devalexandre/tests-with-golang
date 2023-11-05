package gorm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnection(database string) *Database {
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Database{db}
}
