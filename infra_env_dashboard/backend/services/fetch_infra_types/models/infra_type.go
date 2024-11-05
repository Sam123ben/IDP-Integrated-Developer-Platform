package models

type InfraType struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Type     string `json:"type"`
	Sections string `json:"sections"` // Assuming sections is stored as a JSON or string in DB
}

func (InfraType) TableName() string {
	return "infra_types"
}
