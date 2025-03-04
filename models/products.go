package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Price int    `json:"price" gorm:"not null"`
}
