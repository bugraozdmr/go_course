package main

import (
	"basiccrud/internal/config"
	db "basiccrud/internal/database"
	user "basiccrud/internal/user/entity"
	"log"
)

func main() {
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load environements")
	}
	database := db.ConnectPGDB(cnf)

	if err := database.AutoMigrate(&user.UserRegister{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	if err := database.AutoMigrate(&user.UserLogin{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Migration successful")
}
