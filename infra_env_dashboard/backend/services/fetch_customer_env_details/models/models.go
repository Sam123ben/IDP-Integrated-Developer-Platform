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

type CustomerEnvUpdate struct {
	ID            int    `json:"id"`
	CustomerName  string `json:"customer_name" binding:"required"`
	ProductName   string `json:"product_name" binding:"required"`
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

type Customer struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type Environment struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	CustomerID int    `json:"customer_id"`
	ProductID  int    `json:"product_id"`
	Name       string `json:"name"`
}