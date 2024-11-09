// main.go
package main

import (
	"log"
	"notifications-test-1/database"
	"notifications-test-1/models"
)

func main() {
	// Initialize the database
	database.Init()

	// Run migrations
	err := database.DB.AutoMigrate(
		&models.Team{},
		&models.User{},
		&models.Product{},
		&models.Incident{},
		&models.Subscription{},
		&models.Message{},
	)
	if err != nil {
		log.Fatal("failed to run migrations:", err)
	}

	log.Println("Database migration completed successfully.")
}
