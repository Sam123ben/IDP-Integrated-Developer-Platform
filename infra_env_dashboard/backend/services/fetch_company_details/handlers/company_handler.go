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

// CreateCompany godoc
// @Summary Create a new company
// @Description Create a new company and store it in the database
// @Tags Company
// @Accept json
// @Produce json
// @Param company body models.Company true "Company data"
// @Success 201 {object} models.Company
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/company [post]
func (handler *CompanyHandler) CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if err := handler.Repo.CreateCompany(&company); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create company"})
		return
	}
	c.JSON(http.StatusCreated, company)
}

// GetCompany godoc
// @Summary Get a list of companies
// @Description Retrieve a list of companies from the database
// @Tags Company
// @Produce json
// @Success 200 {array} models.Company
// @Failure 500 {object} models.ErrorResponse
// @Router /api/company [get]
func (handler *CompanyHandler) GetCompany(c *gin.Context) {
	companies, err := handler.Repo.GetCompany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch companies"})
		return
	}
	c.JSON(http.StatusOK, companies)
}
