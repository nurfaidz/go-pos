package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Reference string `json:"reference" gorm:"not null"`
	Qty       int    `json:"qty"`
	Total     int    `json:"total"`
	Buyer     string `json:"buyer"`
	UserId    uint   `json:"user_id"`
	User      User
}
