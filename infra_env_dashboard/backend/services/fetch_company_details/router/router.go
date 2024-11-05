package router

import (
	_ "backend/services/fetch_company_details/docs" // Import the generated docs
	"backend/services/fetch_company_details/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swaggerFiles is required to serve Swagger
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(companyHandler *handlers.CompanyHandler) *gin.Engine {
	router := gin.Default()

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := router.Group("/api")
	{
		api.POST("/company", companyHandler.CreateCompany)
		api.GET("/company", companyHandler.Getcompany)
	}

	return router
}
