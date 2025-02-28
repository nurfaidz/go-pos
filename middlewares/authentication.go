package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-pos/exceptions"
	"go-pos/helpers"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			exceptions.BadRequestException(c, err.Error())

			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
