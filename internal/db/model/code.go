package model

import (
	"gorm.io/gorm"
)

type Code struct {
	gorm.Model
	Content  string `gorm:"type:text;not null"`
	Language string `gorm:"not null"`
	SchemaID uint
	Schema   Schema `gorm:"foreignKey:SchemaID"`
}
