package model

import (
	"gorm.io/gorm"
)

type Schema struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
}
