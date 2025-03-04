package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Reference       string            `json:"reference" gorm:"not null"`
	Qty             int               `json:"qty"`
	Total           int               `json:"total"`
	Buyer           string            `json:"buyer"`
	UserID          uint              `json:"user_id"`
	User            User              `gorm:"foreignKey:UserID"`
	TransactionList []TransactionList `gorm:"foreignKey:TransactionID" json:"transaction_list"`
}
