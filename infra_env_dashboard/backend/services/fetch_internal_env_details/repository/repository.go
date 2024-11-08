// repository.go
package repository

import (
	"backend/services/fetch_internal_env_details/models"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InternalRepository struct {
	DB *gorm.DB
}

// NewInternalRepository initializes a new repository instance
func NewInternalRepository(db *gorm.DB) *InternalRepository {
	return &InternalRepository{DB: db}
}

// GetEnvironmentDetailsByEnvName retrieves environment details for a specific product and environment name
func (repo *InternalRepository) GetEnvironmentDetailsByEnvName(product string, envName string) ([]models.EnvironmentDetail, error) {
	var environmentDetails []models.EnvironmentDetail

	// Query to filter by product name and environment name
	err := repo.DB.
		Table("environment_details").
		Joins("JOIN environments ON environments.id = environment_details.environment_id").
		Joins("JOIN products ON products.id = environments.product_id").
		Where("LOWER(products.name) = LOWER(?) AND LOWER(environments.name) = LOWER(?)", product, envName).
		Select("environment_details.*").
		Find(&environmentDetails).Error

	// Log query results for debugging purposes
	if err != nil {
		logrus.Errorf("Database query failed: %v", err)
		return nil, fmt.Errorf("failed to retrieve environment details: %w", err)
	}

	logrus.Infof("Fetched environment details: %v", environmentDetails)
	return environmentDetails, nil
}
