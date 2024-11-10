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

func (repo *CompanyRepository) CreateCompany(company *models.Company) error {
	return repo.DB.Create(company).Error
}

// Rename Getcompany to GetCompany
func (repo *CompanyRepository) GetCompany() ([]models.Company, error) {
	var companies []models.Company
	err := repo.DB.Find(&companies).Error
	return companies, err
}
