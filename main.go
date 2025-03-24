package main

import (
	"go-crud-app/config"
	"go-crud-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect Database
	config.InitDB()

	// root route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// Register Routes
	routes.UserRoutes(r)
	routes.ProductRoutes(r)

	// Start Server
	r.Run(":8080")
}
