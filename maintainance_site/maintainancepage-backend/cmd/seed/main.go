// backend/cmd/seed/main.go
package main

import (
	"log"
	"maintainancepage/internal/config"
	"maintainancepage/internal/database"
	"maintainancepage/internal/models"
	"time"
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

	// AutoMigrate to ensure schema is up-to-date
	if err := models.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Get current time in UTC
	currentTimeUTC := time.Now().UTC()
	log.Printf("Creating maintenance window starting at (UTC): %v", currentTimeUTC)

	// Seed Maintenance Window
	maintenanceWindow := models.MaintenanceWindow{
		StartTime:         currentTimeUTC,
		EstimatedDuration: 120, // 2 hours
		Description:       "Scheduled system maintenance to improve stability and performance.",
	}

	if err := db.Create(&maintenanceWindow).Error; err != nil {
		log.Fatalf("Failed to create maintenance window: %v", err)
	}

	// Seed System Components
	components := []models.SystemComponent{
		{Name: "API Services", Type: "api", Status: models.StatusMaintenance},
		{Name: "UI Services", Type: "ui", Status: models.StatusMaintenance},
		{Name: "Database", Type: "database", Status: models.StatusOperational},
		{Name: "CDN", Type: "cdn", Status: models.StatusOperational},
		{Name: "Azure Services", Type: "azs", Status: models.StatusOperational},
	}

	if err := db.Create(&components).Error; err != nil {
		log.Fatalf("Failed to create components: %v", err)
	}

	// Associate Components with Maintenance Window
	if err := db.Model(&maintenanceWindow).Association("Components").Replace(components); err != nil {
		log.Fatalf("Failed to associate components: %v", err)
	}

	// Seed Maintenance Updates
	updates := []models.MaintenanceUpdate{
		{MaintenanceWindowID: maintenanceWindow.ID, Message: "Started system maintenance and security updates."},
		{MaintenanceWindowID: maintenanceWindow.ID, Message: "Database backup completed successfully."},
		{MaintenanceWindowID: maintenanceWindow.ID, Message: "Deploying system updates."},
		{MaintenanceWindowID: maintenanceWindow.ID, Message: "Finalizing maintenance. Expected completion in 30 minutes."},
	}

	if err := db.Create(&updates).Error; err != nil {
		log.Fatalf("Failed to create updates: %v", err)
	}

	// Log results in UTC
	log.Println("Seed data inserted successfully (all times in UTC):")
	log.Printf("Maintenance Window ID: %d", maintenanceWindow.ID)
	log.Printf("Start Time (UTC): %v", maintenanceWindow.StartTime)
	log.Printf("End Time (UTC): %v", maintenanceWindow.EndTime())

	log.Println("Components:")
	for _, component := range components {
		log.Printf("  - %s [%s]: %s", component.Name, component.Type, component.Status)
	}

	log.Println("Updates:")
	for _, update := range updates {
		log.Printf("  - %s", update.Message)
	}
}
