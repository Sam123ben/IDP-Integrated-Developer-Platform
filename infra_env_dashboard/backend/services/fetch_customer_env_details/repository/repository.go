// repository/repository.go
package repository

import (
	"backend/services/fetch_customer_env_details/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

// GetAllCustomers retrieves all customers from the database
func (repo *CustomerRepository) GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	if err := repo.DB.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

// GetCustomerEnvironmentDetails retrieves customer environment details
func (repo *CustomerRepository) GetCustomerEnvironmentDetails(customer string, product string) ([]models.EnvironmentDetail, error) {
	var environmentDetails []models.EnvironmentDetail
	err := repo.DB.
		Table("environment_details").
		Joins("JOIN environments ON environments.id = environment_details.environment_id").
		Joins("JOIN products ON products.id = environments.product_id").
		Joins("JOIN customers ON customers.id = environments.customer_id").
		Where("customers.name = ? AND products.name = ?", customer, product).
		Select("environment_details.*").
		Find(&environmentDetails).Error

	if err != nil {
		return nil, err
	}
	return environmentDetails, nil
}