package repository

import (
	"backend/services/fetch_infra_types/models"
	"log"

	"github.com/lib/pq" // Import pq for PostgreSQL array handling
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

	// Preload sections for each infra type
	if err := repo.DB.Preload("Sections").Find(&infraTypes).Error; err != nil {
		return nil, err
	}

	// Example: Convert StringArray to pq.StringArray if needed
	for i := range infraTypes {
		for j := range infraTypes[i].Sections {
			infraTypes[i].Sections[j].Environments = models.StringArray(pq.StringArray(infraTypes[i].Sections[j].Environments))
		}
	}

	log.Printf("Fetched infraTypes: %+v\n", infraTypes)
	return infraTypes, nil
}
