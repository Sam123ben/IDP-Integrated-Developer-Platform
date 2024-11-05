package models

type InfraType struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

func (InfraType) TableName() string {
	return "infra_types"
}
