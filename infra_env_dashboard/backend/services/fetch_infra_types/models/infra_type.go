package models

// StringArray is a type alias for []string to help Swagger generate documentation.
type StringArray []string

type InfraType struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Sections []Section `json:"sections" gorm:"foreignKey:InfraTypeID;references:ID"`
}

type Section struct {
	ID           int         `gorm:"primaryKey" json:"id"`
	InfraTypeID  int         `json:"infra_type_id"`
	Name         string      `json:"name"`
	Environments StringArray `gorm:"type:text[]" json:"environments"`
}
