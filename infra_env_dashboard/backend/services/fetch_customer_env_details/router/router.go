// router/router.go
package router

import (
	"backend/services/fetch_customer_env_details/handlers"
	"backend/services/fetch_customer_env_details/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCustomerEnvRoutes(api *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewCustomerRepository(db)
	handler := handlers.NewCustomerEnvHandler(repo)

	api.GET("/customer-env-details", handler.FetchCustomerEnvDetails)
}