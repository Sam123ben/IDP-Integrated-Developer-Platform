// router.go
package router

import (
	"backend/services/fetch_internal_env_details/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(internalHandler *handlers.InternalEnvHandler) *gin.Engine {
	router := gin.Default()

	// Apply CORS settings
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type"},
	}))

	// Define API routes
	api := router.Group("/api")
	{
		api.GET("/internal-env-details", internalHandler.FetchInternalEnvDetails)
	}

	return router
}
