package handlers

import (
	"backend/services/fetch_customer_env_details/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CustomerEnvHandler struct {
	Repo *repository.CustomerEnvRepository
}

func NewCustomerEnvHandler(repo *repository.CustomerEnvRepository) *CustomerEnvHandler {
	return &CustomerEnvHandler{Repo: repo}
}

// FetchCustomerEnvDetails godoc
// @Summary Fetch customer environment details
// @Description Retrieves customer-specific environment details for a vendor and product
// @Tags Customer Environment
// @Produce json
// @Param vendor query string true "Vendor Name"
// @Param product query string true "Product Name"
// @Success 200 {array} models.EnvironmentDetail
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /api/customer-env-details [get]
func (h *CustomerEnvHandler) FetchCustomerEnvDetails(c *gin.Context) {
	vendor := c.Query("vendor")
	product := c.Query("product")

	// Validate required parameters
	if vendor == "" || product == "" {
		logrus.Error("Vendor and product parameters are required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vendor and product parameters are required"})
		return
	}

	// Retrieve environment details using vendor and product
	environmentDetails, err := h.Repo.GetEnvironmentDetails(vendor, product)
	if err != nil {
		logrus.Errorf("Failed to fetch customer environment details: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customer environment details"})
		return
	}

	c.JSON(http.StatusOK, environmentDetails)
}