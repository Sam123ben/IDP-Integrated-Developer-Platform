package repository

import (
	"backend/services/fetch_infra_types/models"
	"log"

	"gorm.io/gorm"
)

type InfraRepository struct {
	DB *gorm.DB
}

func NewInfraRepository(db *gorm.DB) *InfraRepository {
	return &InfraRepository{DB: db}
}

func (repo *InfraRepository) GetAllInfraTypes() ([]models.InfraType, error) {
	var infraTypes []models.InfraType

	// Use GORM's Preload to load sections for each infra type
	if err := repo.DB.Preload("Sections").Find(&infraTypes).Error; err != nil {
		return nil, err
	}

	log.Printf("Fetched infraTypes: %+v\n", infraTypes) // Add debug logs
	return infraTypes, nil
}
