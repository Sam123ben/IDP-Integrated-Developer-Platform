// handlers/handler.go
package handlers

import (
	"backend/services/fetch_customer_env_details/models"
	"backend/services/fetch_customer_env_details/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CustomerEnvHandler struct {
	Repo *repository.CustomerRepository
}

func NewCustomerEnvHandler(repo *repository.CustomerRepository) *CustomerEnvHandler {
	return &CustomerEnvHandler{Repo: repo}
}

// FetchCustomerEnvDetails godoc
// @Summary Fetch customer environment details
// @Description Retrieves environment details for a specific customer and product
// @Tags Customer Environment
// @Produce json
// @Param customer query string true "Customer Name"
// @Param product query string true "Product Name"
// @Success 200 {object} map[string]interface{} "environmentDetails"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /customer-env-details [get]
func (h *CustomerEnvHandler) FetchCustomerEnvDetails(c *gin.Context) {
	customer := c.Query("customer")
	product := c.Query("product")

	// Validate required parameters
	if customer == "" || product == "" {
		logrus.Error("Customer and product parameters are required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer and product parameters are required"})
		return
	}

	// Fetch environment details
	environmentDetails, err := h.Repo.GetCustomerEnvironmentDetails(customer, product)
	if err != nil {
		logrus.Errorf("Failed to fetch customer environment details: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customer environment details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"environmentDetails": environmentDetails})
}

// UpdateEnvironmentDetails godoc
// @Summary Update or Insert customer environment details
// @Description Updates an existing environment detail or inserts a new one if it doesn't exist
// @Tags Customer Environment
// @Accept json
// @Produce json
// @Param environment body models.CustomerEnvUpdate true "Environment Detail with Customer and Product"
// @Success 200 {object} map[string]string "message"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /customer-env-details [put]
func (h *CustomerEnvHandler) UpdateEnvironmentDetails(c *gin.Context) {
	var envUpdate models.CustomerEnvUpdate

	// Bind JSON payload to the struct
	if err := c.ShouldBindJSON(&envUpdate); err != nil {
		logrus.Errorf("Invalid request payload: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Get the customer ID
	customerID, err := h.Repo.GetCustomerID(envUpdate.CustomerName)
	if err != nil {
		logrus.Errorf("Invalid customer name: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer name"})
		return
	}

	// Get the product ID
	productID, err := h.Repo.GetProductID(envUpdate.ProductName)
	if err != nil {
		logrus.Errorf("Invalid product name: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product name"})
		return
	}

	// Get or create the environment ID
	environmentID, err := h.Repo.GetOrCreateEnvironment(customerID, productID, envUpdate.Name)
	if err != nil {
		logrus.Errorf("Failed to resolve environment: %v", err)
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