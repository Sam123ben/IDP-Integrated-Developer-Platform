// backend/handlers/generate_handler.go

package handlers

import (
	"backend/models"
	"backend/services"
	"encoding/json"
	"net/http"
)

// GenerateTerraformHandler handles HTTP requests to generate Terraform files.
func GenerateTerraformHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := services.GenerateTerraform(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Terraform code generated successfully"))
}
