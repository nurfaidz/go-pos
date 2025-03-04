package models

import "gorm.io/gorm"

type TransactionList struct {
	gorm.Model
	TransactionId uint `json:"transaction_id"`
	ProductID     uint `json:"product_id"`
	Qty           int  `json:"qty"`
	Amount        int  `json:"amount"`
	Total         int  `json:"total"`
}
