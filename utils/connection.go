package utils

import (
	"basic-crud-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		return
	}
}
