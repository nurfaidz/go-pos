package controllers

import (
	"github.com/gin-gonic/gin"
	"go-pos/config"
	"go-pos/exceptions"
	"go-pos/models"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	var Product models.Product

	if err := c.ShouldBindJSON(&Product); err != nil {
		exceptions.BadRequestException(c, err.Error())

		return
	}

	if err := config.Connection().Create(&Product).Error; err != nil {
		exceptions.InternalServerErrorException(c, err.Error())

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created",
		"data":    &Product,
	})
}

func GetProductList(c *gin.Context) {
	var Product models.Product

	if err := config.Connection().Find(&Product).Error; err != nil {
		exceptions.InternalServerErrorException(c, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &Product,
	})
}

func GetProductDetail(c *gin.Context) {
	var Product models.Product
	id := c.Param("id")

	if config.Connection().First(&Product, id).Error != nil {
		exceptions.NotFoundException(c)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &Product,
	})
}

func UpdateProduct(c *gin.Context) {
	var Product models.Product
	id := c.Param("id")

	if config.Connection().First(&Product, id).Error != nil {
		exceptions.NotFoundException(c)

		return
	}

	if err := c.ShouldBindJSON(&Product); err != nil {
		exceptions.BadRequestException(c, err.Error())

		return
	}

	if err := config.Connection().Updates(&Product).Error; err != nil {
		exceptions.InternalServerErrorException(c, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated",
		"data":    &Product,
	})
}

func DeleteProduct(c *gin.Context) {
	var Product models.Product
	id := c.Param("id")

	if config.Connection().First(&Product, id).Error != nil {
		exceptions.NotFoundException(c)

		return
	}

	if err := config.Connection().Delete(&Product).Error; err != nil {
		exceptions.InternalServerErrorException(c, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})
}
