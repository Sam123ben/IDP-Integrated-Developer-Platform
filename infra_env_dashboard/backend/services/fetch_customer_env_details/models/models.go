// models/models.go
package models

type EnvironmentDetail struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	EnvironmentID int    `json:"environment_id"`
	Name          string `json:"name"`
	URL           string `json:"url"`
	LastUpdated   string `json:"lastUpdated"`
	Status        string `json:"status"`
	Contact       string `json:"contact"`
	AppVersion    string `json:"appVersion"`
	DBVersion     string `json:"dbVersion"`
	Comments      string `json:"comments"`
}

type Product struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

// TableName overrides the default table name for Product
func (Product) TableName() string {
	return "products"
}

type Customer struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

// TableName overrides the default table name for Customer
func (Customer) TableName() string {
	return "customers"
}