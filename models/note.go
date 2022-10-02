package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Note struct {
	gorm.Model
	UUID      string     `gorm:"primaryKey" json:"id"`
	Title     string     `json:"title" validate:"nonzero,min=2,max=100"`
	Body      string     `json:"body" validate:"nonzero,min=2,max=1500"`
	UserUUID  string     `gorm:"foreignKey" json:"user_id"`
	User      *User      `json:"user,omitempty"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (n *Note) BeforeCreate(tx *gorm.DB) (err error) {
	n.UUID = uuid.New().String()
	return nil
}
