package test

import (
    "infra_env_dashboard/internal/models"
    "testing"
)

func TestGetLatestData(t *testing.T) {
    db := setupTestDB()
    defer db.Close()

    environments, err := environments.GetLatestData(db)
    if err != nil {
        t.Fatalf("Failed to fetch data: %v", err)
    }

    // Example: Check for one of the new fields
    for _, env := range environments {
        if env.EnvironmentType == "" {
            t.Errorf("Expected environment type, got empty string")
        }
        // Add more assertions for other new fields if needed
    }
}