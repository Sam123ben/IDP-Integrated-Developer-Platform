package main

import (
	"backend/common/logger"
	"backend/common/postgress"
	companyRouter "backend/services/fetch_company_details/router"          // Import Company Router
	infraRouter "backend/services/fetch_infra_types/router"                // Import Infra Router
	internalEnvRouter "backend/services/fetch_internal_env_details/router" // Import Internal Env Router

	_ "backend/docs" // Import Swagger docs

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Unified Backend API
// @version 1.0
// @description Unified API for Company, Infra Types, and Internal Environment Details services.
// @host localhost:8080
// @BasePath /api
func main() {
	// Initialize the logger
	logger.Logger.Info("Starting backend services")

	// Initialize the database connection
	db, err := postgress.InitDB()
	if err != nil {
		logger.Logger.Fatalf("Database connection failed: %v", err)
	}

	// Set up the main router
	router := gin.Default()

	// Apply CORS middleware globally
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Swagger setup
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register service-specific routes under /api
	api := router.Group("/api")
	{
		// Call the setup functions from each router, passing the api group and database connection
		companyRouter.SetupCompanyRoutes(api, db)
		infraRouter.SetupInfraRoutes(api, db)
		internalEnvRouter.SetupInternalEnvRoutes(api, db)
	}

	// Start the server on port 8080
	logger.Logger.Info("Server running on port 8080")
	if err := router.Run(":8080"); err != nil {
		logger.Logger.Fatalf("Failed to run server: %v", err)
	}
}
