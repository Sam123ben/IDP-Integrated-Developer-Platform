// backend/cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"maintainancepage/internal/config"
	croncleanup "maintainancepage/internal/cron"
	"maintainancepage/internal/database"
	"maintainancepage/internal/handlers"
	"maintainancepage/internal/models"
	"maintainancepage/internal/router"
)

func main() {
	// Setup logger with flags
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load configuration from environment or .env file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize the maintenance handler
	maintenanceHandler := handlers.NewMaintenanceHandler(db)

	// Initialize the system component handler
	systemComponentHandler := handlers.NewSystemComponentHandler(db)

	// Set up the router and apply middlewares
	r := router.SetupRouter(maintenanceHandler, systemComponentHandler)

	// Perform migrations
	log.Println("🗄️ Performing database migrations...")
	if err := models.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Start the cleanup cron job
	log.Println("⏰ Starting cleanup cron job...")
	croncleanup.StartCleanupCron(db)

	// Define the server address
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("🚀 Server starting on %s", addr)
	log.Printf("📝 CORS enabled for all origins in development mode")
	log.Printf("🔊 Request logging enabled")

	// Create and start the HTTP server
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// Start the server and handle any errors
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
