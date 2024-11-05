package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standardized error response structure.
type ErrorResponse struct {
	Error string `json:"error"`
}

// HandleHTTPError writes a JSON error response for non-Gin HTTP contexts.
func HandleHTTPError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

// RespondWithError writes a JSON error response for Gin contexts.
func RespondWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{Error: message})
}
