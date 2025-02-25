package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Username string `json:"username" gorm:"not null"`
	Password string `gorm:"not null"`
}
