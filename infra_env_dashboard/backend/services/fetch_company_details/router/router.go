package router

import (
	"backend/services/fetch_company_details/handlers"
	"backend/services/fetch_company_details/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCompanyRoutes(api *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewCompanyRepository(db)
	handler := handlers.NewCompanyHandler(repo)

	api.GET("/company", handler.GetCompanyDetails) // Get company details
	api.PUT("/company", handler.UpsertCompany)    // Add or update company details
}