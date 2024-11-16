package models

type EnvironmentDetail struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	LastUpdated string `json:"last_updated"`
	Status      string `json:"status"`
	Contact     string `json:"contact"`
	AppVersion  string `json:"app_version"`
	DBVersion   string `json:"db_version"`
	Comments    string `json:"comments"`
}