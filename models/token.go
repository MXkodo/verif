package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	UserID       string `gorm:"not null"`
	RefreshToken string `gorm:"unique;not null"`
	IP           string `gorm:"not null"`
}
