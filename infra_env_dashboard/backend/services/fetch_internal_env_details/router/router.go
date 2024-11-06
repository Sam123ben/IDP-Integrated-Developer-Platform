package router

import (
	"backend/services/fetch_internal_env_details/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(internalHandler *handlers.InternalHandler) *gin.Engine {
	router := gin.Default()

	// Apply CORS settings
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Define API routes
	api := router.Group("/api")
	{
		api.GET("/internal-env-details", internalHandler.GetInternalEnvDetails)
	}
	return router
}
