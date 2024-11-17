// models/models.go
package models

type Company struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
}

// TableName overrides the default table name
func (Company) TableName() string {
	return "company"
}