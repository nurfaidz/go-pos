package controllers

import (
	"github.com/gin-gonic/gin"
	"go-pos/config"
	"go-pos/models"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	var Product models.Product

	if err := c.ShouldBindJSON(&Product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})

		return
	}

	create := config.Connection().Create(&Product)

	if create.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Product is not created",
			"err":     create.Error,
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created",
		"data":    &Product,
	})
}

func GetProductList(c *gin.Context) {
	var Product models.Product

	c.JSON(http.StatusOK, gin.H{
		"data": &Product,
	})
}

func GetProductDetail(c *gin.Context) {
	var Product models.Product
	id := c.Param("id")

	if config.Connection().First(&Product, id).Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &Product,
	})
}

func UpdateProduct(c *gin.Context) {
	var Product models.Product
	id := c.Param("id")

	if config.Connection().First(&Product, id).Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
	}

	if err := c.ShouldBindJSON(&Product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"err":     err.Error(),
		})

		return
	}

	update := config.Connection().Updates(&Product)

	if update.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Product is not updated",
			"err":     update.Error,
		})

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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
	}

	remove := config.Connection().Delete(&Product)

	if remove.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Product is not deleted",
			"err":     remove.Error,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
	})
}
