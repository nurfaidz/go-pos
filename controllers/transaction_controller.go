package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pos/config"
	"go-pos/exceptions"
	"go-pos/models"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GetTransactionList(c *gin.Context) {
	var Transaction models.Transaction

	if err := config.Connection().Find(&Transaction).Error; err != nil {
		exceptions.InternalServerErrorException(c, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &Transaction,
	})
}

func GetTransactionDetail(c *gin.Context) {
	var Transaction models.Transaction
	id := c.Param("id")

	if config.Connection().First(&Transaction, id).Error != nil {
		exceptions.NotFoundException(c)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &Transaction,
	})
}

func CreateTransaction(c *gin.Context) {
	var Transaction models.Transaction

	if err := c.ShouldBindJSON(&Transaction); err != nil {
		exceptions.BadRequestException(c, err.Error())

		return
	}

	db := config.Connection()

	for i, item := range Transaction.TransactionList {
		var Product models.Product

		if err := db.First(&Product, item.ProductID).Error; err != nil {
			exceptions.InternalServerErrorException(c, err.Error())

			return
		}

		Transaction.TransactionList[i].Amount = Product.Price
		Transaction.TransactionList[i].Total = Transaction.TransactionList[i].Amount * item.Qty
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		Transaction.Reference = fmt.Sprintf("TR%d", time.Now().Unix())

		if err := tx.Create(&Transaction).Error; err != nil {
			return err
		}

		var totalQty, totalAmount int

		for i := range Transaction.TransactionList {
			Transaction.TransactionList[i].ID = 0
			Transaction.TransactionList[i].TransactionID = Transaction.ID

			totalQty += Transaction.TransactionList[i].Qty
			totalAmount += Transaction.TransactionList[i].Total
		}

		if err := tx.CreateInBatches(&Transaction.TransactionList, len(Transaction.TransactionList)).Error; err != nil {
			return err
		}

		if err := tx.Model(&Transaction).Updates(models.Transaction{
			Qty:   totalQty,
			Total: totalAmount,
		}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		exceptions.InternalServerErrorException(c, err.Error())

		return
	}

	db.Preload("User").
		Preload("TransactionList.Product").
		First(&Transaction, Transaction.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Transaction created successfully",
		"data":    &Transaction,
	})

}
