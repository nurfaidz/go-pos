package controllers

import (
	"github.com/gin-gonic/gin"
	"go-pos/config"
	"go-pos/exceptions"
	"go-pos/models"
	"net/http"
)

func GetUserList(c *gin.Context) {
	var User models.User

	getData := config.Connection().Find(&User)
	if getData.Error != nil {
		exceptions.InternalServerErrorException(c, getData.Error)

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
		exceptions.BadRequestException(c, err.Error)

		return
	}

	updateData := config.Connection().Updates(&User)
	if updateData.Error != nil {
		exceptions.InternalServerErrorException(c, updateData.Error)

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

	User.Password = "admin123"

	createData := config.Connection().Create(&User)
	if createData.Error != nil {
		exceptions.InternalServerErrorException(c, createData.Error)

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

	deleteData := config.Connection().Delete(&User)
	if deleteData.Error != nil {
		exceptions.InternalServerErrorException(c, deleteData.Error)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User removed",
	})
}
