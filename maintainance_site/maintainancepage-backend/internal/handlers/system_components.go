// backend/internal/handlers/system_components.go
package handlers

import (
	"log"
	"net/http"

	"maintainancepage/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SystemComponentHandler struct {
	db *gorm.DB
}

// NewSystemComponentHandler initializes a new handler for system components
func NewSystemComponentHandler(db *gorm.DB) *SystemComponentHandler {
	return &SystemComponentHandler{db: db}
}

// CreateSystemComponent handles the creation of a new system component
func (h *SystemComponentHandler) CreateSystemComponent(c *gin.Context) {
	var component models.SystemComponent

	// Parse the incoming JSON payload
	if err := c.ShouldBindJSON(&component); err != nil {
		log.Printf("Error parsing payload: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Set default status to 'Operational' if not provided
	if component.Status == "" {
		component.Status = models.StatusOperational
	}

	// Validate required fields
	if component.Name == "" || component.Type == "" || string(component.Status) == "" {
		log.Println("Validation failed: missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields: name, type, or status"})
		return
	}

	// Check if the component already exists
	var existingComponent models.SystemComponent
	if err := h.db.Where("name = ?", component.Name).First(&existingComponent).Error; err == nil {
		// Component with the same name already exists
		log.Printf("Component with name '%s' already exists", component.Name)
		c.JSON(http.StatusConflict, gin.H{
			"error":   "Component with the same name already exists",
			"message": "Duplicate component not allowed",
		})
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		// Handle other database errors
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Create a new component
	if err := h.db.Create(&component).Error; err != nil {
		log.Printf("Failed to create component: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create system component"})
		return
	}

	log.Printf("System component created: %+v", component)
	c.JSON(http.StatusCreated, gin.H{
		"message":   "System component created successfully",
		"component": component,
	})
}

// UpdateComponentDetails handles the PUT request to update type and status for an existing component by name
func (h *SystemComponentHandler) UpdateComponentDetails(c *gin.Context) {
	var payload struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Status string `json:"status"`
	}

	// Parse the incoming JSON payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Printf("Error parsing payload: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Validate required fields
	if payload.Name == "" || payload.Type == "" || payload.Status == "" {
		log.Println("Validation failed: missing required fields", payload.Name, payload.Type, payload.Status)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields: name, type, or status"})
		return
	}

	// Fetch the existing component by name
	var component models.SystemComponent
	err := h.db.Where("name = ?", payload.Name).First(&component).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Component not found: %s", payload.Name)
			c.JSON(http.StatusNotFound, gin.H{"error": "Component not found"})
			return
		}
		log.Printf("Failed to fetch component: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch component"})
		return
	}

	// Update the type and status
	component.Type = payload.Type
	component.Status = models.SystemStatus(payload.Status)

	// Save the changes
	if err := h.db.Save(&component).Error; err != nil {
		log.Printf("Failed to update component: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update component"})
		return
	}

	// Respond with success
	log.Printf("Component updated successfully: %+v", component)
	c.JSON(http.StatusOK, gin.H{
		"message":   "Component updated successfully",
		"component": component,
	})
}

// GetSystemComponents handles fetching all system components
func (h *SystemComponentHandler) GetSystemComponents(c *gin.Context) {
	var components []models.SystemComponent

	// Fetch all components
	if err := h.db.Find(&components).Error; err != nil {
		log.Printf("Error fetching system components: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch system components"})
		return
	}

	// Prepare response
	response := make([]map[string]interface{}, len(components))
	for i, component := range components {
		response[i] = map[string]interface{}{
			"id":   component.ID,
			"name": component.Name,
			"type": component.Type,
			"status": map[string]interface{}{
				"key":   component.Status,
				"value": string(component.Status),
			},
		}
	}

	// Send response
	c.JSON(http.StatusOK, response)
}
