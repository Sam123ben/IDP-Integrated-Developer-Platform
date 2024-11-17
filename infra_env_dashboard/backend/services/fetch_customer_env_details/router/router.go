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
	api.PUT("/customer-env-details", handler.UpdateEnvironmentDetails)
}

// Sample Body for PUT /customer-env-details
// {
// 	"id": 1,
// 	"customer_name": "Vendor A",
// 	"product_name": "Product 1",
// 	"name": "New Environment",
// 	"url": "https://newenv.example.com",
// 	"lastUpdated": "2024-11-17T12:34:56",
// 	"status": "Online",
// 	"contact": "John Doe",
// 	"appVersion": "v2.0.0",
// 	"dbVersion": "v2.0.0",
// 	"comments": "Environment added for testing"
// }