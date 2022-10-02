package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UUID      string     `gorm:"primaryKey" json:"id"`
	FirstName string     `json:"first_name" validate:"nonzero,min=2,max=100"`
	LastName  string     `json:"last_name" validate:"nonzero,min=2,max=100"`
	Email     string     `gorm:"unique" json:"email" validate:"nonzero,min=6,max=100,email"`
	Password  string     `json:"password" validate:"nonzero,min=6"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	Notes     []Note     `json:"notes"`
}

func (u *User) CheckAlreadyExists(db *gorm.DB) bool {
	var count int64
	db.Model(&User{}).Where("email = ?", u.Email).Count(&count)
	return count > 0
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()
	return nil
}
