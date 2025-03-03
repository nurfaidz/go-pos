package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name            string            `json:"name" gorm:"not null"`
	Price           int               `json:"price" gorm:"not null"`
	TransactionList []TransactionList `gorm:"foreignKey:ProductID"`
	DeletedAt       gorm.DeletedAt    `gorm:"index" json:"deleted_at,omitempty"`
}
