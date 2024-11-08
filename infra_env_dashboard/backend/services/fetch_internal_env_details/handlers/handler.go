// handler.go
package handlers

import (
	"backend/services/fetch_internal_env_details/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type InternalEnvHandler struct {
	Repo *repository.InternalRepository
}

// NewInternalEnvHandler initializes a new handler with the repository
func NewInternalEnvHandler(repo *repository.InternalRepository) *InternalEnvHandler {
	return &InternalEnvHandler{Repo: repo}
}

// FetchInternalEnvDetails retrieves environment details for a specific product and environment group
func (h *InternalEnvHandler) FetchInternalEnvDetails(c *gin.Context) {
	product := c.Query("product")
	envName := c.Query("EnvName")

	// Define valid product names
	validProducts := []string{"Product 1", "Product 2"}

	// Check if provided product name matches valid products
	isValidProduct := false
	for _, validProduct := range validProducts {
		if strings.EqualFold(product, validProduct) {
			product = validProduct // Use the correct casing
			isValidProduct = true
			break
		}
	}

	if !isValidProduct {
		logrus.Warnf("Invalid product name provided: %s. Valid options are: %v", product, validProducts)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product name. Please use 'Product 1' or 'Product 2'."})
		return
	}

	// Log the incoming parameters to confirm they are being received correctly
	logrus.Infof("Fetching environment details for product: %s, environment: %s", product, envName)

	// Validate required parameters
	if product == "" || envName == "" {
		logrus.Error("Product and EnvName parameters are required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product and EnvName parameters are required"})
		return
	}

	// Retrieve environment details
	environmentDetails, err := h.Repo.GetEnvironmentDetailsByEnvName(product, envName)
	if err != nil {
		logrus.Errorf("Failed to fetch environment details: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch environment details"})
		return
	}

	if len(environmentDetails) == 0 {
		logrus.Warnf("No environment details found for product: %s, environment: %s", product, envName)
	}

	c.JSON(http.StatusOK, gin.H{"environmentDetails": environmentDetails})
}
