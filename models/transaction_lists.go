package models

import "gorm.io/gorm"

type TransactionList struct {
	gorm.Model
	Transaction   Transaction `gorm:"TransactionID"`
	TransactionID uint        `json:"transaction_id"`
	Product       Product     `gorm:"ProductID"`
	ProductID     uint        `json:"product_id"`
	Qty           int         `json:"qty"`
	Amount        int         `json:"amount"`
	Total         int         `json:"total"`
}
