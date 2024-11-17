// main.go
package main

import (
	"backend/common/logger"
	"backend/common/postgress"
	companyRouter "backend/services/fetch_company_details/router"
	internalEnvRouter "backend/services/fetch_internal_env_details/router"
	customerEnvRouter "backend/services/fetch_customer_env_details/router"
	"fmt"
	"os"

	_ "backend/docs" // Swagger docs

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Unified Backend API
// @version 1.0
// @description Unified API for Company, Internal Environment Details, and Customer Environment Details services.
// @host localhost:8080
// @BasePath /api
func main() {
	// Initialize logger
	logger.Logger.Info("Starting backend services")

	// Initialize the database connection
	db, err := postgress.InitDB()
	if err != nil {
		logger.Logger.Fatalf("Database connection failed: %v", err)
	}

	// Set up the main router
	router := gin.Default()

	// Apply CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Set Swagger host dynamically
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "localhost:8080"
	}
	if err := os.Setenv("SWAGGER_HOST", swaggerHost); err != nil {
		logger.Logger.Warnf("Failed to set SWAGGER_HOST: %v", err)
	}

	// Swagger setup
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register API routes
	api := router.Group("/api")
	{
		companyRouter.SetupCompanyRoutes(api, db)         // For GET and PUT company details
		internalEnvRouter.SetupInternalEnvRoutes(api, db) // For internal environment details
		customerEnvRouter.SetupCustomerEnvRoutes(api, db) // For customer environment details
	}

	// Get port from environment variable or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := "0.0.0.0"

	logger.Logger.Infof("Server running on %s:%s", host, port)
	if err := router.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		logger.Logger.Fatalf("Failed to run server: %v", err)
	}
}