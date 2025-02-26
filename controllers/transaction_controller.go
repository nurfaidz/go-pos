package controllers

import (
	"github.com/gin-gonic/gin"
	"go-pos/config"
	"go-pos/exceptions"
	"go-pos/models"
	"net/http"
)

func GetTransactionList(c *gin.Context) {
	var Transaction models.Transaction

	if err := config.Connection().Find(&Transaction); err.Error != nil {
		exceptions.InternalServerErrorException(c, err)

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
		exceptions.BadRequestException(c, err.Error)

		return
	}

	if err := config.Connection().Create(&Transaction); err.Error != nil {
		exceptions.InternalServerErrorException(c, err.Error)

		return
	}
}
