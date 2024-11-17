// handlers/handler.go
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

	// Fetch the list of valid products from the database
	products, err := h.Repo.GetAllProducts()
	if err != nil {
		logrus.Errorf("Failed to fetch products: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	// Validate product name
	isValidProduct := false
	for _, p := range products {
		if strings.EqualFold(product, p.Name) {
			product = p.Name // Use the correct casing from the database
			isValidProduct = true
			break
		}
	}

	if !isValidProduct {
		logrus.Warnf("Invalid product name provided: %s. Valid options are: %v", product, products)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product name. Please provide a valid product name."})
		return
	}

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