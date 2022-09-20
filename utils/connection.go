package utils

import (
	"basic-crud-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		panic("Unable to migrate database")
	}
	return db
}
