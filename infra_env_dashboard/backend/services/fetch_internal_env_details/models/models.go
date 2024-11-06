package models

import (
	_ "github.com/lib/pq"
)

// EnvironmentDetail represents the details for each environment.
type EnvironmentDetail struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	EnvironmentID int    `json:"environment_id"` // Foreign key to Environment
	Name          string `json:"name"`
	URL           string `json:"url"`
	LastUpdated   string `json:"lastUpdated"`
	Status        string `json:"status"`
	Contact       string `json:"contact"`
	AppVersion    string `json:"appVersion"`
	DBVersion     string `json:"dbVersion"`
	Comments      string `json:"comments"`
}

// Environment represents each environment type within a product.
type Environment struct {
	ID        int                 `gorm:"primaryKey" json:"id"`
	ProductID int                 `json:"product_id"` // Foreign key to Product
	Name      string              `json:"name"`
	Details   []EnvironmentDetail `gorm:"foreignKey:EnvironmentID;references:ID" json:"details"` // Link to EnvironmentDetail
}

// Product represents a product with multiple environments.
type Product struct {
	ID           int           `gorm:"primaryKey" json:"id"`
	Name         string        `json:"name"`
	Environments []Environment `gorm:"foreignKey:ProductID;references:ID" json:"environments"` // Link to Environment
}
