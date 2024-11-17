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

// GetCustomerID retrieves the ID for a given customer name
func (repo *CustomerRepository) GetCustomerID(customerName string) (int, error) {
	var customer models.Customer
	err := repo.DB.Where("LOWER(name) = LOWER(?)", customerName).First(&customer).Error
	if err != nil {
		return 0, err
	}
	return customer.ID, nil
}

// GetProductID retrieves the ID for a given product name
func (repo *CustomerRepository) GetProductID(productName string) (int, error) {
	var product models.Product
	err := repo.DB.Where("LOWER(name) = LOWER(?)", productName).First(&product).Error
	if err != nil {
		return 0, err
	}
	return product.ID, nil
}

// GetOrCreateEnvironment retrieves or creates the environment
func (repo *CustomerRepository) GetOrCreateEnvironment(customerID, productID int, envName string) (int, error) {
	var environment models.Environment
	err := repo.DB.Where("customer_id = ? AND product_id = ? AND LOWER(name) = LOWER(?)", customerID, productID, envName).
		First(&environment).Error

	if err == gorm.ErrRecordNotFound {
		environment = models.Environment{
			CustomerID: customerID,
			ProductID:  productID,
			Name:       envName,
		}
		if err := repo.DB.Create(&environment).Error; err != nil {
			return 0, err
		}
		return environment.ID, nil
	}

	return environment.ID, err
}

// GetCustomerEnvironmentDetails retrieves customer environment details
func (repo *CustomerRepository) GetCustomerEnvironmentDetails(customer string, product string) ([]models.EnvironmentDetail, error) {
	var environmentDetails []models.EnvironmentDetail
	err := repo.DB.
		Table("environment_details").
		Joins("JOIN environments ON environments.id = environment_details.environment_id").
		Joins("JOIN products ON products.id = environments.product_id").
		Joins("JOIN customers ON customers.id = environments.customer_id").
		Where("LOWER(customers.name) = LOWER(?) AND LOWER(products.name) = LOWER(?)", customer, product).
		Select("environment_details.*").
		Find(&environmentDetails).Error

	if err != nil {
		return nil, err
	}

	return environmentDetails, nil
}

// UpsertEnvironmentDetails inserts or updates environment details
func (repo *CustomerRepository) UpsertEnvironmentDetails(envDetail models.EnvironmentDetail) error {
	if err := repo.DB.Where("id = ?", envDetail.ID).Assign(envDetail).FirstOrCreate(&envDetail).Error; err != nil {
		return err
	}
	return nil
}