package main

import (
	analytics "VisiGo/internal/analytics/entity"
	"VisiGo/internal/config"
	db "VisiGo/internal/database"
	"log"
)

func main() {
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load environements")
	}

	database := db.ConnectPGDB(cnf)

	if err := database.AutoMigrate(
		&analytics.Visit{},    // Visit
		&analytics.PageView{}, // PageView
		&analytics.Event{},    // Event
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Migration successful")
}
