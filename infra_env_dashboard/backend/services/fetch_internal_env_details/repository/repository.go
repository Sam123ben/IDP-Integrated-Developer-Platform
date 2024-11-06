package repository

import (
	"backend/services/fetch_internal_env_details/models"
	"log"

	"gorm.io/gorm"
)

type InternalRepository struct {
	DB *gorm.DB
}

func NewInternalRepository(db *gorm.DB) *InternalRepository {
	return &InternalRepository{DB: db}
}

func (repo *InternalRepository) GetInternalEnvDetails() ([]models.Product, error) {
	var products []models.Product

	// Use GORM's Preload to load environments and details
	if err := repo.DB.Preload("Environments.Details").Find(&products).Error; err != nil {
		return nil, err
	}

	log.Printf("Fetched products: %+v\n", products) // Add debug logs
	return products, nil
}
