// main.go
package main

import (
	"backend/common/postgress"
	"backend/services/fetch_internal_env_details/handlers"
	"backend/services/fetch_internal_env_details/repository"
	"backend/services/fetch_internal_env_details/router"
	"log"

	"gorm.io/gorm/logger"
)

// @title Internal Environment Details API
// @version 1.0
// @description API Documentation for Internal Environment Details
// @host localhost:8082
// @BasePath /api
func main() {
	// Initialize PostgreSQL database with GORM, enabling GORM logger for SQL output
	db, err := postgress.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Set GORM to log all SQL queries at the Info level
	db.Logger = logger.Default.LogMode(logger.Info)

	// Initialize repository and handler
	internalRepo := repository.NewInternalRepository(db)
	internalHandler := handlers.NewInternalEnvHandler(internalRepo)

	// Set up router
	r := router.SetupRouter(internalHandler)

	// Start the server
	log.Println("Server running on port 8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
