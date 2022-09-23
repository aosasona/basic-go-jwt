package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `json:"first_name" validate:"nonzero,min=2,max=100"`
	LastName  string    `json:"last_name" validate:"nonzero,min=2,max=100"`
	Password  string    `json:"password" validate:"nonzero,min=6"`
	CreatedAt time.Time `json:"created_at"`
}

type Note struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `json:"title" validate:"nonzero,min=2,max=100"`
	Body      string `json:"body" validate:"nonzero,min=2"`
	CreatedAt string `json:"created_at"`
	UserId    int    `json:"user_id"`
	User      User   `json:"user"`
}
