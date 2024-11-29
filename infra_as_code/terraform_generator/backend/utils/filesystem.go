// backend/utils/filesystem.go

package utils

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// LoadConfig reads the configuration from a JSON file
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

// Convert a value to a JSON string
func toJSON(value interface{}) (string, error) {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// CreateDirectories ensures that the specified directories exist
func CreateDirectories(paths []string) error {
	for _, path := range paths {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// WriteFile writes content to a specified path
func WriteFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0644)
}

// GenerateFileFromTemplate generates a file from a template
func GenerateFileFromTemplate(templatePath, destinationPath string, data interface{}) error {
	funcMap := template.FuncMap{
		"title": cases.Title(language.Und).String,
		"add":   func(a, b int) int { return a + b },
		"toJSON": func(value interface{}) string {
			jsonString, err := toJSON(value)
			if err != nil {
				return "null"
			}
			return jsonString
		},
		"typeOf": func(value interface{}) string {
			switch value.(type) {
			case string:
				return "string"
			case bool:
				return "bool"
			case int, float64:
				return "number"
			case []interface{}:
				return "list"
			case map[string]interface{}:
				return "map"
			case map[string]string:
				return "map(string)"
			default:
				return "any"
			}
		},
		"or": func(a, b interface{}) interface{} {
			if a != nil {
				return a
			}
			return b
		},
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

	// Execute the template
	var outputBuffer bytes.Buffer
	if err := tmpl.Execute(&outputBuffer, data); err != nil {
		return err
	}

	return WriteFile(destinationPath, outputBuffer.Bytes())
}
