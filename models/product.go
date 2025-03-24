package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Validator instance
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateProduct validates the Product struct

// ValidateProduct validates the Product struct and returns human-readable error messages
func (p *Product) ValidateProduct() (map[string]string, error) {
	err := validate.Struct(p)
	if err != nil {
		// Map to store custom error messages
		errorMessages := make(map[string]string)

		// Loop through validation errors
		for _, e := range err.(validator.ValidationErrors) {
			// Custom error messages based on the validation rule
			switch e.Tag() {
			case "required":
				errorMessages[e.Field()] = fmt.Sprintf("%s is required", e.Field())
			case "min":
				errorMessages[e.Field()] = fmt.Sprintf("%s should be at least %s characters", e.Field(), e.Param())
			case "max":
				errorMessages[e.Field()] = fmt.Sprintf("%s cannot exceed %s characters", e.Field(), e.Param())
			case "gt":
				errorMessages[e.Field()] = fmt.Sprintf("%s must be greater than zero", e.Field())
			case "gte":
				errorMessages[e.Field()] = fmt.Sprintf("%s cannot be less than zero", e.Field())
			}
		}
		return errorMessages, err
	}
	return nil, nil
}