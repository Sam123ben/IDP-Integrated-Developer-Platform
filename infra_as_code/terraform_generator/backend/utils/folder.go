package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFolderStructure creates the required Terraform folder structure
func CreateFolderStructure(baseOutputDir string) error {
	// Define folder structure
	folders := []string{
		filepath.Join(baseOutputDir, "environments", "dev"),
		filepath.Join(baseOutputDir, "modules", "resource_group"),
		filepath.Join(baseOutputDir, "global-modules", "network"),
	}

	// Map folders to default files
	files := map[string][]string{
		"environments/dev":       {"variables.tf", "outputs.tf", "backend.tf"},
		"modules/resource_group": {"main.tf", "variables.tf", "outputs.tf"},
		"global-modules/network": {"main.tf", "variables.tf", "outputs.tf"},
	}

	// Create folders and associated files
	for _, folder := range folders {
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			if err := os.MkdirAll(folder, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create folder %s: %v", folder, err)
			}
		}

		// Create default files in the folder
		for _, file := range files[filepath.Base(folder)] {
			filePath := filepath.Join(folder, file)
			if _, err := os.Create(filePath); err != nil {
				return fmt.Errorf("failed to create file %s: %v", filePath, err)
			}
		}
	}

	// Create main.tf and provider.tf in the product folder
	productFiles := []string{"main.tf", "provider.tf"}
	for _, file := range productFiles {
		filePath := filepath.Join(baseOutputDir, file)
		if _, err := os.Create(filePath); err != nil {
			return fmt.Errorf("failed to create file %s: %v", filePath, err)
		}
	}

	return nil
}
