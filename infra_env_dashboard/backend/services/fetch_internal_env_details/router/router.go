package router

import (
	"backend/services/fetch_internal_env_details/handlers"
	"backend/services/fetch_internal_env_details/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupInternalEnvRoutes(api *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewInternalRepository(db)
	handler := handlers.NewInternalEnvHandler(repo)

	api.GET("/internal-env-details", handler.FetchInternalEnvDetails)
}
