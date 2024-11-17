package repository

import (
	"backend/services/fetch_company_details/models"
	"errors"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}

// GetCompany retrieves the company details from the database
func (repo *CompanyRepository) GetCompany() (models.Company, error) {
	var company models.Company
	if err := repo.DB.First(&company).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return company, nil // Return an empty record if not found
		}
		return company, err
	}
	return company, nil
}

// UpsertCompany inserts or updates the company record in the database
func (repo *CompanyRepository) UpsertCompany(company models.Company) error {
	return repo.DB.Save(&company).Error
}