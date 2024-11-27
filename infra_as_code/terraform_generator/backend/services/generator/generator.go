// backend/services/generator/generator.go

package generator

import (
	"backend/services/config"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// GenerateProviderFile generates the provider.tf file using the template from the templates folder
func GenerateProviderFile(baseOutputDir string, providerConfig config.ProviderConfig) error {
	// Define the template path
	templatePath := filepath.Join("templates", "azure", "provider", "provider.tf.tmpl")

	// Check if the template exists
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return fmt.Errorf("template not found at %s", templatePath)
	}

	// Generate the output file
	providerFile := filepath.Join(baseOutputDir, "provider.tf")
	err := renderTemplateToFile(providerFile, templatePath, providerConfig)
	if err != nil {
		return fmt.Errorf("failed to render provider.tf: %v", err)
	}

	return nil
}

// GenerateBackendFile generates the backend.tf file using the template from the templates folder
func GenerateBackendFile(baseOutputDir string, backendConfig config.BackendConfig) error {
	// Define the template path
	templatePath := filepath.Join("templates", "azure", "provider", "backend.tf.tmpl")

	// Check if the template exists
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return fmt.Errorf("template not found at %s", templatePath)
	}

	// Generate the output file
	backendFile := filepath.Join(baseOutputDir, "backend.tf")
	err := renderTemplateToFile(backendFile, templatePath, backendConfig)
	if err != nil {
		return fmt.Errorf("failed to render backend.tf: %v", err)
	}

	return nil
}

// renderTemplateToFile reads a template from the given path and renders it to the specified file
func renderTemplateToFile(filePath, templatePath string, data interface{}) error {
	// Parse the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %v", templatePath, err)
	}

	// Create the output file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer file.Close()

	// Render the template to the file
	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to render template to file %s: %v", filePath, err)
	}

	return nil
}
