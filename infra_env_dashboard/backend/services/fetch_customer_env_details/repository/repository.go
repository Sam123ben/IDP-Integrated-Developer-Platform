package repository

import (
	"backend/services/fetch_customer_env_details/models"

	"gorm.io/gorm"
)

type CustomerEnvRepository struct {
	DB *gorm.DB
}

func NewCustomerEnvRepository(db *gorm.DB) *CustomerEnvRepository {
	return &CustomerEnvRepository{DB: db}
}

// GetEnvironmentDetails retrieves environment details by vendor and product
func (repo *CustomerEnvRepository) GetEnvironmentDetails(vendor, product string) ([]models.EnvironmentDetail, error) {
	var environmentDetails []models.EnvironmentDetail

	err := repo.DB.
		Table("environment_details").
		Joins("JOIN products ON products.id = environment_details.product_id").
		Joins("JOIN vendors ON vendors.id = products.vendor_id").
		Where("LOWER(vendors.name) = LOWER(?) AND LOWER(products.name) = LOWER(?)", vendor, product).
		Select("environment_details.name, environment_details.url, environment_details.last_updated, environment_details.status, environment_details.contact, environment_details.app_version, environment_details.db_version, environment_details.comments").
		Find(&environmentDetails).Error

	if err != nil {
		return nil, err
	}

	return environmentDetails, nil
}