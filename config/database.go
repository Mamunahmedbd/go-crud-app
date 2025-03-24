package config

import (
	"fmt"
	"go-crud-app/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres password=postgres dbname=go-app port=5432 sslmode=disable TimeZone=Asia/Dhaka",
		PreferSimpleProtocol: true,
	 }), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	DB = database
	fmt.Println("Connected to database")
	DB.AutoMigrate(&models.User{}, &models.Product{})
}
