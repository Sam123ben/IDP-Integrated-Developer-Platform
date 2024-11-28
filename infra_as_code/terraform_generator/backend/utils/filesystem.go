// backend/utils/filesystem.go

package utils

import (
	"backend/models"
	"encoding/json"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Add function
func add(a, b int) int {
	return a + b
}

func CreateDirectories(paths []string) error {
	for _, path := range paths {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func WriteFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0644)
}

func GenerateFileFromTemplate(templatePath, destinationPath string, data interface{}) error {
	// Register custom functions
	funcMap := template.FuncMap{
		"title": cases.Title(language.Und).String,
		"add":   add, // Ensure 'add' is included
	}

	// Parse the template with the function map
	tmpl, err := template.New(filepath.Base(templatePath)).Funcs(funcMap).ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Ensure the destination directory exists
	destDir := filepath.Dir(destinationPath)
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return err
	}

	destFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Execute the template
	return tmpl.Execute(destFile, data)
}

func LoadConfig(path string) (*models.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config models.Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
