package router

import (
	"backend/services/fetch_infra_types/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(infraHandler *handlers.InfraHandler) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/infra-types", infraHandler.GetInfraTypes)
	}
	return router
}
