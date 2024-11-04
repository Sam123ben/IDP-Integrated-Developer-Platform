package models

import (
	"company_service/common"
	"log"
)

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetCompanyName() (Company, error) {
	var company Company

	query := "SELECT id, name FROM company LIMIT 1"
	row := common.DB.QueryRow(query)
	if err := row.Scan(&company.ID, &company.Name); err != nil {
		log.Printf("Error fetching company name: %s", err)
		return company, err
	}

	return company, nil
}
