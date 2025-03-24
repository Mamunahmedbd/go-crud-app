package routes

import (
	"go-crud-app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}
}
