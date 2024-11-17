// handlers/handler.go
package handlers

import (
	"backend/services/fetch_internal_env_details/models"
	"backend/services/fetch_internal_env_details/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type InternalEnvHandler struct {
	Repo *repository.InternalRepository
}

func NewInternalEnvHandler(repo *repository.InternalRepository) *InternalEnvHandler {
	return &InternalEnvHandler{Repo: repo}
}

// FetchInternalEnvDetails godoc
// @Summary Fetch internal environment details
// @Description Retrieves environment details for a specific product and environment group
// @Tags Internal Environment
// @Produce json
// @Param product query string true "Product Name"
// @Param group query string true "Environment Group"
// @Success 200 {object} map[string]interface{} "environmentDetails"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /internal-env-details [get]
func (h *InternalEnvHandler) FetchInternalEnvDetails(c *gin.Context) {
	product := c.Query("product")
	group := c.Query("group") // Environment group

	// Validate required parameters
	if product == "" || group == "" {
		logrus.Error("Product and group parameters are required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product and group parameters are required"})
		return
	}

	// Retrieve environment details using product and group
	environmentDetails, err := h.Repo.GetEnvironmentDetailsByEnvName(product, group)
	if err != nil {
		logrus.Errorf("Failed to fetch environment details: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch environment details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"environmentDetails": environmentDetails})
}

// UpdateEnvironmentDetails godoc
// @Summary Update or Insert internal environment details
// @Description Updates an existing environment detail or inserts a new one if it doesn't exist
// @Tags Internal Environment
// @Accept json
// @Produce json
// @Param environment body models.InternalEnvUpdate true "Environment Detail with Product and Group"
// @Success 200 {object} map[string]string "message"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /internal-env-details [put]
func (h *InternalEnvHandler) UpdateEnvironmentDetails(c *gin.Context) {
	var envUpdate models.InternalEnvUpdate

	// Bind JSON payload to the struct
	if err := c.ShouldBindJSON(&envUpdate); err != nil {
		logrus.Error("Invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Get the product ID
	productID, err := h.Repo.GetProductID(envUpdate.ProductName)
	if err != nil {
		logrus.Error("Invalid product name")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product name"})
		return
	}

	// Get or create the environment ID
	environmentID, err := h.Repo.GetOrCreateEnvironment(productID, envUpdate.GroupName)
	if err != nil {
		logrus.Error("Failed to resolve environment")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to resolve environment"})
		return
	}

	// Populate the EnvironmentDetail struct
	envDetail := models.EnvironmentDetail{
		ID:            envUpdate.ID,
		EnvironmentID: environmentID,
		Name:          envUpdate.Name,
		URL:           envUpdate.URL,
		LastUpdated:   envUpdate.LastUpdated,
		Status:        envUpdate.Status,
		Contact:       envUpdate.Contact,
		AppVersion:    envUpdate.AppVersion,
		DBVersion:     envUpdate.DBVersion,
		Comments:      envUpdate.Comments,
	}

	// Call the Upsert method in the repository
	if err := h.Repo.UpsertEnvironmentDetails(envDetail); err != nil {
		logrus.Errorf("Failed to update environment details: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update environment details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Environment details updated successfully"})
}