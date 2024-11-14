package main

import (
	"backend/common/logger"
	"backend/common/postgress"
	companyRouter "backend/services/fetch_company_details/router"          // Import Company Router
	infraRouter "backend/services/fetch_infra_types/router"                // Import Infra Router
	internalEnvRouter "backend/services/fetch_internal_env_details/router" // Import Internal Env Router
	"fmt"
	"os"

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

	// Set Swagger host dynamically from environment variables for deployment flexibility
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "localhost:8080" // Default to localhost for local development
	}
	// Update the Swagger host dynamically
	ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName(swaggerHost))

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

	// Get the port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Log where the server will run
	host := "0.0.0.0" // This makes it accessible on any host interface (suitable for cloud and local)
	logger.Logger.Infof("Server running on %s:%s", host, port)

	// Start the server
	if err := router.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		logger.Logger.Fatalf("Failed to run server: %v", err)
	}
}