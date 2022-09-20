package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
	Email     string
}

type Note struct {
	gorm.Model
	Title     string
	Body      string
	CreatedAt string
	User      *User
}
