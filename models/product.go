package models

import "github.com/go-playground/validator/v10"

type Product struct {
	Name        string  `validate:"required,min=3,max=50"`
	Description string  `validate:"required,min=10,max=200"`
	Price       float64 `validate:"required,gt=0"`
	Stock       int     `validate:"required,gte=0"`
}

// Validator instance
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateProduct validates the Product struct
func (p *Product) ValidateProduct() error {
	return validate.Struct(p)
}
