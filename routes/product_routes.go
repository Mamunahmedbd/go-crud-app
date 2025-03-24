package routes

import (
	"go-crud-app/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	productGroup := r.Group("/products")
	{
		productGroup.POST("/", controllers.CreateProduct)
		productGroup.GET("/", controllers.GetProducts)
	}
}