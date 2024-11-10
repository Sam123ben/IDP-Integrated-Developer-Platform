package router

import (
	"backend/services/fetch_infra_types/handlers"
	"backend/services/fetch_infra_types/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupInfraRoutes(api *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewInfraRepository(db)
	handler := handlers.NewInfraHandler(repo)

	api.GET("/infra-types", handler.GetInfraTypes)
}
