package repository

import (
	"backend/services/fetch_infra_types/models"

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
	if err := repo.DB.Find(&infraTypes).Error; err != nil {
		return nil, err
	}
	return infraTypes, nil
}
