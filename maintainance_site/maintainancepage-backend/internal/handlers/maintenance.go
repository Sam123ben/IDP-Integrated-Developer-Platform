// backend/internal/handlers/maintenance.go
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"maintainancepage/internal/models"

	"gorm.io/gorm"
)

type MaintenanceHandler struct {
	db *gorm.DB
}

func NewMaintenanceHandler(db *gorm.DB) *MaintenanceHandler {
	return &MaintenanceHandler{db: db}
}

// GetActiveMaintenance fetches the first active or fixed maintenance window
func (h *MaintenanceHandler) GetActiveMaintenance(w http.ResponseWriter, r *http.Request) {
	log.Println("GetActiveMaintenance called")

	var maintenance models.MaintenanceWindow

	// Current time in UTC
	currentTimeUTC := time.Now().UTC()
	log.Printf("Current UTC Time: %s", currentTimeUTC)

	// Client timezone
	location := h.getClientLocation(r)
	currentTimeLocal := currentTimeUTC.In(location)
	log.Printf("Client Local Time: %s", currentTimeLocal)

	// Query for the active maintenance window
	err := h.db.Preload("Components").Preload("Updates").
		Where("status = ?", models.StatusActive).
		First(&maintenance).Error

	// If no active maintenance window is found, attempt to fetch the most recent fixed one
	if err == gorm.ErrRecordNotFound {
		log.Println("No active maintenance window found. Trying to fetch the most recent fixed maintenance window.")

		err = h.db.Preload("Components").Preload("Updates").
			Where("status = ?", models.StatusFixed).
			Order("updated_at DESC").
			First(&maintenance).Error

		if err == gorm.ErrRecordNotFound {
			log.Println("No fixed maintenance window found either.")
			h.sendEmptyResponse(w, currentTimeUTC, currentTimeLocal)
			return
		}

		if err != nil {
			log.Printf("Database error during fixed maintenance query: %v", err)
			http.Error(w, "Failed to fetch maintenance window", http.StatusInternalServerError)
			return
		}
	} else if err != nil {
		log.Printf("Database error during active maintenance query: %v", err)
		http.Error(w, "Failed to fetch maintenance window", http.StatusInternalServerError)
		return
	}

	// log.Printf("Maintenance Window Found: %+v", maintenance)

	response := models.MaintenanceResponse{
		CurrentTimeUTC:   currentTimeUTC.Format(time.RFC3339),
		CurrentTimeLocal: currentTimeLocal.Format(time.RFC3339),
		IsActive:         maintenance.Status == models.StatusActive,
		RemainingTime:    maintenance.RemainingMinutes(currentTimeUTC),
		Maintenance: models.MaintenanceWindowData{
			ID:          maintenance.ID,
			StartTime:   maintenance.StartTime.Format(time.RFC3339),
			Description: maintenance.Description,
			Components:  separateComponentsByStatus(maintenance.Components),
			Updates:     getPaginatedUpdates(maintenance.Updates, r),
		},
	}

	h.sendJSONResponse(w, response)
}

// CreateOrUpdateMaintenanceWindow handles creating or updating the maintenance window
func (h *MaintenanceHandler) CreateOrUpdateMaintenanceWindow(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateOrUpdateMaintenanceWindow called")

	var maintenance models.MaintenanceWindow

	// Decode the incoming request payload
	if err := json.NewDecoder(r.Body).Decode(&maintenance); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure StartTime is in UTC
	maintenance.StartTime = maintenance.StartTime.UTC()
	log.Printf("StartTime converted to UTC: %s", maintenance.StartTime)

	// Fetch all existing components
	var components []models.SystemComponent
	if err := h.db.Find(&components).Error; err != nil {
		log.Printf("Error fetching components: %v", err)
		http.Error(w, "Failed to fetch components", http.StatusInternalServerError)
		return
	}
	log.Printf("Fetched Components: %+v", components)

	// Check if a maintenance window already exists
	var existingMaintenance models.MaintenanceWindow
	err := h.db.First(&existingMaintenance).Error

	if err != nil {
		// Create a new maintenance window if none exists
		if err == gorm.ErrRecordNotFound {
			log.Println("No existing maintenance window found. Creating a new one.")

			// Default values for new maintenance windows
			if maintenance.Status == "" {
				maintenance.Status = models.StatusActive
			}
			if !maintenance.IssueFixed {
				maintenance.IssueFixed = false
			}

			// Link components to the new maintenance window
			maintenance.Components = components

			// Create the new maintenance window
			if createErr := h.db.Create(&maintenance).Error; createErr != nil {
				log.Printf("Error creating maintenance window: %v", createErr)
				http.Error(w, "Failed to create maintenance window", http.StatusInternalServerError)
				return
			}

			log.Println("Maintenance window created successfully")
			w.WriteHeader(http.StatusCreated)
			h.sendJSONResponse(w, map[string]string{"message": "Maintenance window created successfully"})
			return
		}

		log.Printf("Database error during maintenance window query: %v", err)
		http.Error(w, "Failed to query maintenance window", http.StatusInternalServerError)
		return
	}

	// Update existing maintenance window
	log.Printf("Existing Maintenance Window Found: %+v", existingMaintenance)

	// Update values from payload
	existingMaintenance.StartTime = maintenance.StartTime
	existingMaintenance.EstimatedDuration = maintenance.EstimatedDuration
	existingMaintenance.Description = maintenance.Description
	existingMaintenance.Components = components

	// Check and handle `issue_fixed` flag
	if maintenance.IssueFixed {
		existingMaintenance.IssueFixed = true
		existingMaintenance.Status = models.StatusFixed
	} else {
		existingMaintenance.IssueFixed = false
		existingMaintenance.Status = models.StatusActive
	}

	// Save the updated maintenance window
	if updateErr := h.db.Save(&existingMaintenance).Error; updateErr != nil {
		log.Printf("Error updating maintenance window: %v", updateErr)
		http.Error(w, "Failed to update maintenance window", http.StatusInternalServerError)
		return
	}

	log.Println("Maintenance window updated successfully")
	h.sendJSONResponse(w, map[string]string{"message": "Maintenance window updated successfully"})
}

// CreateMaintenanceUpdate handles creating a new maintenance update
func (h *MaintenanceHandler) CreateMaintenanceUpdate(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateMaintenanceUpdate called")

	// Parse the incoming JSON payload
	var updatePayload struct {
		Message    string `json:"message"`
		IssueFixed bool   `json:"issue_fixed,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updatePayload); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Decoded Update Payload: %+v", updatePayload)

	// Fetch the active maintenance window
	var activeMaintenance models.MaintenanceWindow
	err := h.db.Where("status = ? AND issue_fixed = false", models.StatusActive).
		First(&activeMaintenance).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("No active maintenance window found")
			http.Error(w, "No active maintenance window available", http.StatusNotFound)
			return
		}
		log.Printf("Database error fetching active maintenance window: %v", err)
		http.Error(w, "Failed to fetch active maintenance window", http.StatusInternalServerError)
		return
	}
	// log.Printf("Active Maintenance Window Found: %+v", activeMaintenance)

	// Create the maintenance update
	newUpdate := models.MaintenanceUpdate{
		MaintenanceWindowID: activeMaintenance.ID,
		Message:             updatePayload.Message,
	}

	if err := h.db.Create(&newUpdate).Error; err != nil {
		log.Printf("Error creating maintenance update: %v", err)
		http.Error(w, "Failed to create maintenance update", http.StatusInternalServerError)
		return
	}

	log.Println("Maintenance update created successfully")

	// If issue_fixed is true, update the active maintenance window
	if updatePayload.IssueFixed {
		activeMaintenance.IssueFixed = true
		activeMaintenance.Status = models.StatusFixed

		if err := h.db.Save(&activeMaintenance).Error; err != nil {
			log.Printf("Error updating maintenance window issue_fixed status: %v", err)
			http.Error(w, "Failed to update maintenance window status", http.StatusInternalServerError)
			return
		}
		log.Println("Maintenance window marked as fixed")
	}

	// Respond with success
	h.sendJSONResponse(w, map[string]string{
		"message": "Maintenance update created successfully",
	})
}

// Helper Functions
func (h *MaintenanceHandler) getClientLocation(r *http.Request) *time.Location {
	userTimeZone := r.Header.Get("X-Timezone")
	log.Printf("Received Timezone Header: %s", userTimeZone)

	if userTimeZone == "" {
		log.Println("No timezone provided in header. Defaulting to UTC.")
		return time.UTC
	}
	location, err := time.LoadLocation(userTimeZone)
	if err != nil {
		log.Printf("Invalid timezone in header (%s). Defaulting to UTC. Error: %v", userTimeZone, err)
		return time.UTC
	}
	log.Printf("Resolved Timezone Location: %s", location)
	return location
}

func (h *MaintenanceHandler) validateMaintenanceWindow(m *models.MaintenanceWindow) error {
	log.Printf("Validating Maintenance Window: %+v", m)

	if m.StartTime.IsZero() {
		log.Println("Validation failed: start time is required")
		return fmt.Errorf("start time is required")
	}
	if m.EstimatedDuration <= 0 {
		log.Println("Validation failed: estimated duration must be positive")
		return fmt.Errorf("estimated duration must be positive")
	}
	if m.Description == "" {
		log.Println("Validation failed: description is required")
		return fmt.Errorf("description is required")
	}
	log.Println("Validation passed")
	return nil
}

func (h *MaintenanceHandler) sendEmptyResponse(w http.ResponseWriter, currentTimeUTC, currentTimeLocal time.Time) {
	log.Println("Sending empty maintenance response")
	response := models.MaintenanceResponse{
		CurrentTimeUTC:   currentTimeUTC.Format(time.RFC3339),
		CurrentTimeLocal: currentTimeLocal.Format(time.RFC3339),
		IsActive:         false,
		RemainingTime:    0,
		Maintenance:      models.MaintenanceWindowData{},
	}
	h.sendJSONResponse(w, response)
}

func (h *MaintenanceHandler) sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	// log.Printf("Sending JSON Response: %+v", data)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func separateComponentsByStatus(components []models.SystemComponent) map[string][]models.SystemComponent {
	// log.Printf("Separating Components by Status: %+v", components)
	categorized := map[string][]models.SystemComponent{
		"maintenance": {},
		"operational": {},
	}
	for _, component := range components {
		if component.Status == models.StatusMaintenance {
			categorized["maintenance"] = append(categorized["maintenance"], component)
		} else {
			categorized["operational"] = append(categorized["operational"], component)
		}
	}
	// log.Printf("Categorized Components: %+v", categorized)
	return categorized
}

func getPaginationParams(r *http.Request) (int, int) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}
	log.Printf("Pagination Params - Page: %d, Limit: %d", page, limit)
	return page, limit
}

func paginateUpdates(data []models.MaintenanceUpdate, page, limit int) ([]models.MaintenanceUpdate, int) {
	log.Printf("Paginating Updates: Page %d, Limit %d, Data: %+v", page, limit, data)
	total := len(data)
	start := (page - 1) * limit
	end := start + limit

	if start > total {
		log.Println("Pagination Start Index exceeds total updates")
		return []models.MaintenanceUpdate{}, total
	}
	if end > total {
		end = total
	}
	// log.Printf("Paginated Data: %+v", data[start:end])
	return data[start:end], total
}

func getPaginatedUpdates(updates []models.MaintenanceUpdate, r *http.Request) models.PaginatedUpdates {
	// Extract pagination parameters from the request
	page, limit := getPaginationParams(r)

	// Paginate the updates
	paginatedData, total := paginateUpdates(updates, page, limit)

	return models.PaginatedUpdates{
		Data: paginatedData,
		Pagination: models.Pagination{
			CurrentPage:  page,
			PerPage:      limit,
			TotalUpdates: total,
		},
	}
}
