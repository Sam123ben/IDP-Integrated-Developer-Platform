package main

import (
	"log"
	"time"

	"backend/common/postgress"
	_ "backend/services/fetch_infra_types/docs" // Import the generated docs
	"backend/services/fetch_infra_types/handlers"
	"backend/services/fetch_infra_types/repository"
	"backend/services/fetch_infra_types/router"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Infrastructure Types API
// @version 1.0
// @description API Documentation for Infrastructure Types Service
// @host localhost:8081
// @BasePath /api
func main() {
	// Initialize PostgreSQL database with GORM
	db, err := postgress.InitDB() // Ensure this returns *gorm.DB
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize repository and handler
	infraRepo := repository.NewInfraRepository(db)
	infraHandler := handlers.NewInfraHandler(infraRepo)

	// Set up router
	r := router.SetupRouter(infraHandler)

	// Configure and apply CORS middleware
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
	log.Println("Server running on port 8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
