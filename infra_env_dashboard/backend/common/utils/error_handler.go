package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents the structure of the error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// HandleError sends a standardized error response
func HandleError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(ErrorResponse{Error: message})
	if err != nil {
		http.Error(w, "Error encoding error response: "+err.Error(), http.StatusInternalServerError)
	}
}
