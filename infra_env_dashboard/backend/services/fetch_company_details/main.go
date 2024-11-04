package main

import (
	"log"

	"backend/common/postgress"
	_ "backend/services/fetch_company_details/docs" // Import the generated docs
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
	// Initialize PostgreSQL database
	db, err := postgress.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize repository and handler
	companyRepo := repository.NewCompanyRepository(db)
	companyHandler := handlers.NewCompanyHandler(companyRepo)

	// Set up router and add Swagger route
	r := router.SetupRouter(companyHandler)

	// Start the server
	r.Run(":8080")
}
