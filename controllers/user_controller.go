package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-pos/config"
	"go-pos/exceptions"
	"go-pos/models"
	"net/http"
)

func GetUserList(c *gin.Context) {
	var User models.User

	if err := config.Connection().Find(&User).Error; err != nil {
		exceptions.InternalServerErrorException(c, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &User,
	})
}

func GetUserDetail(c *gin.Context) {
	var User models.User
	id := c.Param("id")

	if config.Connection().First(&User, id).Error != nil {
		exceptions.NotFoundException(c)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &User,
	})
}

func UpdateUser(c *gin.Context) {
	var User models.User
	id := c.Param("id")

	if config.Connection().First(&User, id).Error != nil {
		exceptions.NotFoundException(c)

		return
	}

	if err := c.ShouldBindJSON(&User); err != nil {
		exceptions.BadRequestException(c, err.Error())

		return
	}

	if err := config.Connection().Updates(&User).Error; err != nil {
		exceptions.InternalServerErrorException(c, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data updated",
		"data":    &User,
	})
}

func CreateUser(c *gin.Context) {
	var User models.User

	if err := c.ShouldBindJSON(&User); err != nil {
		exceptions.BadRequestException(c, err.Error())

		return
	}

	User.Password = fmt.Sprintf("%s123", User.Username)

	if err := config.Connection().Create(&User).Error; err != nil {
		exceptions.InternalServerErrorException(c, err.Error())

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered",
		"data":    &User,
	})
}

func DeleteUser(c *gin.Context) {
	var User models.User
	id := c.Param("id")

	if config.Connection().First(&User, id).Error != nil {
		exceptions.NotFoundException(c)

		return
	}

	if err := config.Connection().Delete(&User).Error; err != nil {
		exceptions.InternalServerErrorException(c, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User removed",
	})
}
