// handlers/company_handler.go
package handlers

import (
	"backend/services/fetch_company_details/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	Repo *repository.CompanyRepository
}

func NewCompanyHandler(repo *repository.CompanyRepository) *CompanyHandler {
	return &CompanyHandler{Repo: repo}
}

// GetCompanyDetails godoc
// @Summary Get company details
// @Description Retrieves details about the company
// @Tags Company
// @Produce json
// @Success 200 {object} models.Company
// @Failure 500 {object} map[string]string "error"
// @Router /company [get]
func (h *CompanyHandler) GetCompanyDetails(c *gin.Context) {
	company, err := h.Repo.GetCompany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve company details"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"company": company})
}