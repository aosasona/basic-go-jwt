package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database")
	}

	return db
}
