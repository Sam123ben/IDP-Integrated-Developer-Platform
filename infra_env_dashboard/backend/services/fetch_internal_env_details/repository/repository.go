package repository

import (
	"backend/services/fetch_internal_env_details/models"

	"gorm.io/gorm"
)

type InternalRepository struct {
	DB *gorm.DB
}

func NewInternalRepository(db *gorm.DB) *InternalRepository {
	return &InternalRepository{DB: db}
}

// GetAllProducts retrieves all products from the database
func (repo *InternalRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := repo.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// Existing function to get environment details
func (repo *InternalRepository) GetEnvironmentDetailsByEnvName(product string, envName string) ([]models.EnvironmentDetail, error) {
	var environmentDetails []models.EnvironmentDetail
	err := repo.DB.
		Table("environment_details").
		Joins("JOIN environments ON environments.id = environment_details.environment_id").
		Joins("JOIN products ON products.id = environments.product_id").
		Where("LOWER(products.name) = LOWER(?) AND LOWER(environments.name) = LOWER(?)", product, envName).
		Select("environment_details.*").
		Find(&environmentDetails).Error

	if err != nil {
		return nil, err
	}
	return environmentDetails, nil
}
