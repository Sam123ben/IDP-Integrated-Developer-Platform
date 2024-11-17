// handlers/handler.go
package handlers

import (
	"backend/services/fetch_customer_env_details/repository"
	"net/http"
	"strings"

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
// @Description Retrieves environment details for a specific customer, product, and environment group
// @Tags Customer Environment
// @Produce json
// @Param customer query string true "Customer Name"
// @Param product query string true "Product Name"
// @Param group query string true "Environment Group"
// @Success 200 {object} map[string]interface{} "environmentDetails"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /customer-env-details [get]
func (h *CustomerEnvHandler) FetchCustomerEnvDetails(c *gin.Context) {
	customer := c.Query("customer")
	product := c.Query("product")
	group := c.Query("group") // Environment group

	// Validate required parameters
	if customer == "" || product == "" || group == "" {
		logrus.Error("Customer, product, and group parameters are required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer, product, and group parameters are required"})
		return
	}

	// Validate customer name
	customers, err := h.Repo.GetAllCustomers()
	if err != nil {
		logrus.Errorf("Failed to fetch customers: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customers"})
		return
	}

	isValidCustomer := false
	for _, c := range customers {
		if strings.EqualFold(customer, c.Name) {
			customer = c.Name // Use the correct casing from the database
			isValidCustomer = true
			break
		}
	}

	if !isValidCustomer {
		logrus.Warnf("Invalid customer name provided: %s. Valid options are: %v", customer, customers)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer name. Please provide a valid customer name."})
		return
	}

	// Fetch environment details
	environmentDetails, err := h.Repo.GetCustomerEnvironmentDetails(customer, product, group)
	if err != nil {
		logrus.Errorf("Failed to fetch customer environment details: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customer environment details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"environmentDetails": environmentDetails})
}