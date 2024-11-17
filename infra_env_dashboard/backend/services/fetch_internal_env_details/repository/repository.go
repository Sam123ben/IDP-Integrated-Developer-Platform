// repository/repository.go
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

// GetProductID retrieves the ID for a given product name
func (repo *InternalRepository) GetProductID(productName string) (int, error) {
	var product models.Product
	err := repo.DB.Where("LOWER(name) = LOWER(?)", productName).First(&product).Error
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

// GetOrCreateEnvironment retrieves the environment ID or creates one if it doesn't exist
func (repo *InternalRepository) GetOrCreateEnvironment(productID int, groupName string) (int, error) {
	var environment models.Environment
	err := repo.DB.Where("product_id = ? AND LOWER(name) = LOWER(?)", productID, groupName).
		First(&environment).Error

	if err == gorm.ErrRecordNotFound {
		environment = models.Environment{
			ProductID: productID,
			Name:      groupName,
		}
		if err := repo.DB.Create(&environment).Error; err != nil {
			return 0, err
		}
		return environment.ID, nil
	}

	return environment.ID, err
}

// GetEnvironmentDetailsByEnvName retrieves internal environment details
func (repo *InternalRepository) GetEnvironmentDetailsByEnvName(product string, group string) ([]models.EnvironmentDetail, error) {
	var environmentDetails []models.EnvironmentDetail
	err := repo.DB.
		Table("environment_details").
		Joins("JOIN environments ON environments.id = environment_details.environment_id").
		Joins("JOIN products ON products.id = environments.product_id").
		Where("LOWER(products.name) = LOWER(?) AND LOWER(environments.name) = LOWER(?)", product, group).
		Select("environment_details.*").
		Find(&environmentDetails).Error

	return environmentDetails, err
}

// UpsertEnvironmentDetails inserts or updates environment details
func (repo *InternalRepository) UpsertEnvironmentDetails(envDetail models.EnvironmentDetail) error {
	if err := repo.DB.Where("id = ?", envDetail.ID).Assign(envDetail).FirstOrCreate(&envDetail).Error; err != nil {
		return err
	}
	return nil
}