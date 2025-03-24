package models

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"required,gte=18,lte=60"`
	Password string `json:"password" validate:"required,min=8"`
}

func (u *User) ValidateUser() (map[string]string, error) {
	err := validate.Struct(u)
	if err != nil {
		errorMessages := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages[e.Field()] = fmt.Sprintf("%s is required", e.Field())
		}
		return errorMessages, err
	}
	return nil, nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	u.Password = hashPassword(u.Password)
	return
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
