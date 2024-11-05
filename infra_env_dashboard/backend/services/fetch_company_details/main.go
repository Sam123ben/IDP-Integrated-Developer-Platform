package main

import (
	"backend/common/logger"
	"backend/common/postgress"

	"backend/services/fetch_company_details/handlers"
	"backend/services/fetch_company_details/repository"
	"backend/services/fetch_company_details/router"
)

// @title Company Service API
// @version 1.0
// @description API Documentation for Company Service
// @host localhost:8080
// @BasePath /api
func main() {
	// Initialize the database
	db, err := postgress.InitDB()
	if err != nil {
		logger.Logger.Fatalf("Database connection failed: %v", err)
	}

	// Initialize repository and handler
	companyRepo := repository.NewCompanyRepository(db)
	companyHandler := handlers.NewCompanyHandler(companyRepo)

	// Set up the router with CORS and routes
	r := router.SetupRouter(companyHandler)

	// Start the server
	logger.Logger.Info("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		logger.Logger.Fatalf("Failed to run server: %s", err)
	}
}
