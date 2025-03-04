package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pos/config"
	"go-pos/exceptions"
	"go-pos/models"
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
	//var list models.TransactionList

	if err := c.ShouldBindJSON(&Transaction); err != nil {
		exceptions.BadRequestException(c, err.Error())

		return
	}

	err := config.Connection().Debug().Create(&Transaction.TransactionList).Error

	fmt.Println(err)

	Transaction.Reference = fmt.Sprintf("TRX%d", time.Now().Unix())

	var totalQuantity int

	//for item := Transaction.TransactionList {
	//	totalQuantity += item.Qty
	//}
	fmt.Println(totalQuantity)

	Transaction.Qty = totalQuantity
	c.JSON(http.StatusCreated, &Transaction)
	//if err := config.Connection().Create(&Transaction).Error; err != nil {
	//	exceptions.InternalServerErrorException(c, err.Error())
	//
	//	return
	//}
	//
	//c.JSON(http.StatusCreated, gin.H{
	//	"message": "Transaction created",
	//	"data":    &Transaction,
	//})
}
