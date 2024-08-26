package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID    string `gorm:"unique;not null"`
	Email string `gorm:"unique;not null"`
}
