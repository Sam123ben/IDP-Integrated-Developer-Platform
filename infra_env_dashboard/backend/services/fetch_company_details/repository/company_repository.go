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

func (repo *CompanyRepository) Getcompany() ([]models.Company, error) {
	var company []models.Company
	err := repo.DB.Find(&company).Error
	return company, err
}
