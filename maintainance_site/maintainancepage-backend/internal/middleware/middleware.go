// backend/internal/middleware/middleware.go
package middleware

import (
	"log"
	"net/http"
	"time" // Import the time package

	"github.com/gin-gonic/gin"
)

// CorsMiddleware sets up CORS for the server
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Expose-Headers", "Link")

		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

// LoggerMiddleware logs only essential information
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request
		start := time.Now()
		c.Next()

		// Only log if there's an error or if the request took longer than expected
		duration := time.Since(start)
		if c.Writer.Status() >= 400 || duration > time.Second {
			log.Printf("[%d] %s %s %v",
				c.Writer.Status(),
				c.Request.Method,
				c.Request.URL.Path,
				duration,
			)
		}
	}
}

// RecoveryMiddleware recovers from panics and returns a 500 error
func RecoveryMiddleware() gin.HandlerFunc {
	return gin.RecoveryWithWriter(gin.DefaultErrorWriter, func(c *gin.Context, err any) {
		log.Printf("Recovered from panic: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	})
}
