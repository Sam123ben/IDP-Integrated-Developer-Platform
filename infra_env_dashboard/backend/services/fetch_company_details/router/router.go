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

	api.POST("/company", handler.CreateCompany)
	api.GET("/company", handler.GetCompany)
}
