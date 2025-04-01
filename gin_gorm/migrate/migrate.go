package main

import (
	"gin_gorm/initializers"
	"gin_gorm/models"
	"log"

	"github.com/google/uuid"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	post := models.Post{
		ID:    uuid.New(),
		Title: "My First Post",
		Body:  "This is the body of the post.",
	}

	if err := initializers.DB.AutoMigrate(&models.Post{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := initializers.DB.Create(&post).Error; err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}

	log.Println("Migration and insert successful")
}
