package controllers

import (
	"github.com/gin-gonic/gin"
	"go-pos/config"
	"go-pos/exceptions"
	"go-pos/helpers"
	"go-pos/models"
	"net/http"
)

func Login(c *gin.Context) {
	var User models.User

	if err := c.ShouldBindJSON(&User); err != nil {
		exceptions.BadRequestException(c, err.Error())

		return
	}

	password := User.Password

	err := config.Connection().Where("username = ?", User.Username).First(&User).Error
	comparePass := helpers.ComparePass([]byte(password), []byte(User.Password))

	if err != nil || !comparePass {
		exceptions.BadRequestException(c, "Invalid username or password")

		return
	}

	token := helpers.GenerateToken(User.ID, User.Username)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"data": gin.H{
			"token": token,
			"user":  User,
		},
	})
}
