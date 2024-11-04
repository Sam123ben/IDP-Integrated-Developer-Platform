package models

type Company struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (Company) TableName() string {
	return "company"
}
