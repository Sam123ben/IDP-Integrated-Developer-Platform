// backend/cmd/migrate/main.go
package main

import (
	"log"
	"maintainancepage/internal/config"
	"maintainancepage/internal/database"
	"maintainancepage/internal/models"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	log.Println("Running database migrations...")
	if err := models.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrations completed successfully")

	// Insert seed data if the --seed flag is provided
	// TODO: Add seed data insertion logic here
}
