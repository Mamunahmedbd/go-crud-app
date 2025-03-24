package controllers

import (
	"go-crud-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go-crud-app/config"
)

// Create Product
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate product using go-playground/validator/v10
	if err := product.ValidateProduct(); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Error()
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	config.DB.Create(&product)
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": product})
}

// Get All Products
func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}