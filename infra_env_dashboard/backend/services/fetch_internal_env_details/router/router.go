// router/router.go
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
	api.PUT("/internal-env-details", handler.UpdateEnvironmentDetails)
}


// Sample Body for PUT /internal-env-details
// {
// 	"id": 7,
// 	"product_name": "Product 1",
// 	"group_name": "QA",
// 	"name": "New QA Test",
// 	"url": "https://newqa.example.com",
// 	"lastUpdated": "2024-11-17T12:34:56",
// 	"status": "Online",
// 	"contact": "John Doe",
// 	"appVersion": "v3.0.0",
// 	"dbVersion": "v2.0.0",
// 	"comments": "Environment added for QA testing"
// }