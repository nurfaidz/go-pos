package routes

import (
	"github.com/gin-gonic/gin"
	"go-pos/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	product := r.Group("/products")
	{
		product.POST("", controllers.CreateProduct)
		product.GET("", controllers.GetProductList)
		product.GET("/:id", controllers.GetProductDetail)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}

	return r
}
