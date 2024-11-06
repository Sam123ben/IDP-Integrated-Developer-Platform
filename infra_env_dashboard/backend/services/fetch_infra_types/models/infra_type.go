package models

import (
	"github.com/lib/pq"
)

type InfraType struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Sections []Section `json:"sections" gorm:"foreignKey:InfraTypeID;references:ID"` // Define relationship
}

type Section struct {
	ID           int            `gorm:"primaryKey" json:"id"`
	InfraTypeID  int            `json:"infra_type_id"`
	Name         string         `json:"name"`
	Environments pq.StringArray `gorm:"type:text[]" json:"environments"` // Use pq.StringArray for PostgreSQL TEXT[]
}