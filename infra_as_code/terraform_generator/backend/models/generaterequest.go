// backend/models/generaterequest.go

package models

type GenerateRequest struct {
	OrganisationName string   `json:"organisation_name"`
	ProductName      string   `json:"product_name"`
	Customers        []string `json:"customers,omitempty"`
	Provider         string   `json:"provider"`
	Modules          []string `json:"modules"`
}
