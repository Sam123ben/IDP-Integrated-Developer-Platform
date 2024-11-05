package router

import (
	"backend/services/fetch_infra_types/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(infraHandler *handlers.InfraHandler) *gin.Engine {
	router := gin.Default()

	// Apply CORS settings
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust as needed
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Define API routes
	api := router.Group("/api")
	{
		api.GET("/infra-types", infraHandler.GetInfraTypes)
	}
	return router
}
