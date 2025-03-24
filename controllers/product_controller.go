package controllers

import (
	"fmt"
	"go-crud-app/config"
	"go-crud-app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Create Product
func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Println("Product after binding:", product)

	if errorMessages, err := product.ValidateProduct(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	config.DB.Create(&product)
	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": product})
}

// Get All Products
func GetProducts(c *gin.Context) {
	var products []models.Product
	// Get page and limit from query params
	page := c.DefaultQuery("page", "1")    // Default to page 1
	limit := c.DefaultQuery("limit", "10") // Default to 10 products per page

	// Convert to integers
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	// Calculate offset
	offset := (pageInt - 1) * limitInt
	// Get products with limit and offset for pagination
	config.DB.Limit(limitInt).Offset(offset).Find(&products)

	c.JSON(http.StatusOK, gin.H{"products": products})
}

// Get Product by ID
func GetProductByID(c *gin.Context) {
	// Retrieve the id from URL parameter
	id := c.Param("id")

	var product models.Product

	// Validate that the id is numeric (if your id is an integer)
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID format"})
		return
	}

	// Find the product by ID, select specific fields (id, name, price)
	if err := config.DB.Select("id, name, price").First(&product, productID).Error; err != nil {
		// Check if the product was not found
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
		}
		return
	}

	// Return the product data with a status message
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}
