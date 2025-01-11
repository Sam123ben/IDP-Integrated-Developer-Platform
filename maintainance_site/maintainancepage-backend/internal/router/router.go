// backend/internal/router/router.go
package router

import (
	"maintainancepage/internal/handlers"
	"maintainancepage/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns the Gin engine with all routes and middlewares
func SetupRouter(maintenanceHandler *handlers.MaintenanceHandler, componentHandler *handlers.SystemComponentHandler) *gin.Engine {
	// Initialize Gin engine
	router := gin.Default()

	// Apply custom middlewares
	router.Use(middleware.RecoveryMiddleware()) // Invoke the middleware function
	router.Use(middleware.LoggerMiddleware())   // Invoke the middleware function
	router.Use(middleware.CorsMiddleware())     // Invoke the middleware function

	// API group
	api := router.Group("/api")
	{
		// Maintenance routes
		api.GET("/maintenance/active", func(c *gin.Context) {
			maintenanceHandler.GetActiveMaintenance(c.Writer, c.Request)
		})

		// System Component routes
		api.POST("/system-components", componentHandler.CreateSystemComponent)        // Pass context directly
		api.PUT("/system-components/update", componentHandler.UpdateComponentDetails) // Pass context directly

		// Maintainance Windows routes
		api.POST("/maintenance-windows", func(c *gin.Context) {
			maintenanceHandler.CreateOrUpdateMaintenanceWindow(c.Writer, c.Request)
		})

		// Maintenance Update routes
		api.POST("/maintenance-updates", func(c *gin.Context) {
			maintenanceHandler.CreateMaintenanceUpdate(c.Writer, c.Request)
		})
	}

	return router
}
