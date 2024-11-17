// models/models.go
package models

type Company struct {
    ID   int    `gorm:"primaryKey" json:"id"`
    Name string `json:"name"`
}

// TableName overrides the table name used by Company to `company`
func (Company) TableName() string {
    return "company"
}