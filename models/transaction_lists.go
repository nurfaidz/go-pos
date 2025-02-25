package models

import "gorm.io/gorm"

type TransactionList struct {
	gorm.Model
	Transaction   Transaction
	TransactionId uint `json:"transaction_id"`
	Product       Product
	ProductID     uint `json:"product_id"`
	Qty           int  `json:"qty"`
	Amount        int  `json:"amount"`
	Total         int  `json:"total"`
}
