package handlers

import (
	"backend/services/fetch_company_details/models"
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

// UpsertCompany godoc
// @Summary Add or update company details
// @Description Inserts or updates company details in the database
// @Tags Company
// @Accept json
// @Produce json
// @Param company body models.Company true "Company data"
// @Success 200 {object} map[string]string "message"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /company [put]
func (h *CompanyHandler) UpsertCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.Repo.UpsertCompany(company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upsert company details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company details upserted successfully"})
}