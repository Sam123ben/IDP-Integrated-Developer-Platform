package main

import (
	"time"

	"backend/common/logger"
	"backend/common/postgress"
	"backend/services/fetch_internal_env_details/handlers"
	"backend/services/fetch_internal_env_details/repository"
	"backend/services/fetch_internal_env_details/router"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Internal Environment Details API
// @version 1.0
// @description API Documentation for Internal Environment Details Service
// @host localhost:8082
// @BasePath /api
func main() {
	// Initialize PostgreSQL database
	db, err := postgress.InitDB()
	if err != nil {
		logger.Logger.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize repository and handler
	internalRepo := repository.NewInternalRepository(db)
	internalHandler := handlers.NewInternalHandler(internalRepo)

	// Set up router
	r := router.SetupRouter(internalHandler)

	// Configure CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	logger.Logger.Info("Server running on port 8082")
	if err := r.Run(":8082"); err != nil {
		logger.Logger.Fatalf("Could not start server: %s\n", err.Error())
	}
}
