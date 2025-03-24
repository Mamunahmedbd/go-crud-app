package routes

import (
	"go-crud-app/controllers"
	"go-crud-app/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
	}

	userGroup := r.Group("/users")
	{

		userGroup.Use(middleware.AuthMiddleware())
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
		userGroup.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "You are authorized!"})
		})
	}
}
