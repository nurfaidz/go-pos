package routes

import (
	"github.com/gin-gonic/gin"
	"go-pos/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
	}

	product := r.Group("/products")
	{
		product.POST("", controllers.CreateProduct)
		product.GET("", controllers.GetProductList)
		product.GET("/:id", controllers.GetProductDetail)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}

	user := r.Group("/users")
	{
		user.POST("", controllers.CreateUser)
		user.GET("", controllers.GetUserList)
		user.GET("/:id", controllers.GetUserDetail)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	return r
}
