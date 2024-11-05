package router

import (
	"backend/services/fetch_company_details/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(companyHandler *handlers.CompanyHandler) *gin.Engine {
	router := gin.Default()

	// Apply CORS settings
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust as needed
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define API routes
	api := router.Group("/api")
	{
		api.POST("/company", companyHandler.CreateCompany)
		api.GET("/company", companyHandler.Getcompany)
	}

	return router
}
