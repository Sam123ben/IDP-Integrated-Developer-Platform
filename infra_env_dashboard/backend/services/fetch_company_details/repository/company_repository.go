// repository/repository.go
package repository

import (
	"backend/services/fetch_company_details/models"

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
		return company, err
	}
	return company, nil
}