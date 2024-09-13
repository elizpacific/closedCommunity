package main

import (
	"closedCommunity/httpHandlers"
	"closedCommunity/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=elizpacific password=mypassword dbname=closedCom port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Міграції
	db.AutoMigrate(&models.User{}, &models.UserProfile{}, &models.Hobby{}, &models.Link{})

	router := gin.Default()

	// Маршрут для створення користувача
	router.POST("/users", func(c *gin.Context) {
		httpHandlers.CreateUser(c, db)
	})

	router.Run(":8080")
}
