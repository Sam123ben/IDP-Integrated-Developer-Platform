package controllers

import (
	"company_service/models"
	"encoding/json"
	"net/http"
)

func GetCompanyName(w http.ResponseWriter, r *http.Request) {
	company, err := models.GetCompanyName()
	if err != nil {
		http.Error(w, "Error fetching company name", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(company); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
