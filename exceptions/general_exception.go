package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundException(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": "Record not found",
	})
}

func InternalServerErrorException[E any](c *gin.Context, err E) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "Process interrupted",
		"err":     err,
	})
}

func BadRequestException[E any](c *gin.Context, err E) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": "Bad request",
		"err":     err,
	})
}
