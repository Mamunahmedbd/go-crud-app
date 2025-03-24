package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,gte=18,lte=60"`
}
