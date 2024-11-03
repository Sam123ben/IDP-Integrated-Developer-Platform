package models

import "time"

// Environment represents an environment record
type Environment struct {
    ID              int       `json:"id"`
    EnvironmentName string    `json:"environment_name"`
    Description     string    `json:"description"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
    EnvironmentType string    `json:"environment_type"`
    GroupName       string    `json:"group_name"`
    CustomerName    string    `json:"customer_name"`
    URL             string    `json:"url"`
    Status          string    `json:"status"`
    Contact         string    `json:"contact"`
    AppVersion      string    `json:"app_version"`
    DBVersion       string    `json:"db_version"`
    Comments        string    `json:"comments"`
}
