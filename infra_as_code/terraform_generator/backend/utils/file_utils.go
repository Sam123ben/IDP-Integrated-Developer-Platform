// backend/utils/file_utils.go

package utils

import (
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

// ToJSON converts a value to a JSON string
func ToJSON(value interface{}) (string, error) {
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// formatValue dynamically formats values based on their types.
func formatValue(value interface{}, varType string) string {
	switch varType {
	case "string":
		// Quote the value if it's a string and not a variable reference
		if strVal, ok := value.(string); ok {
			if strings.HasPrefix(strVal, "var.") {
				return strVal
			}
			return fmt.Sprintf("\"%s\"", strVal)
		}
		return "null"

	case "bool", "number":
		// Render booleans and numbers as-is
		return fmt.Sprintf("%v", value)

	case "map(string)":
		// Render map values
		if mapVal, ok := value.(map[string]interface{}); ok {
			if len(mapVal) == 0 {
				return "null"
			}
			var entries []string
			for key, val := range mapVal {
				entries = append(entries, fmt.Sprintf("\"%s\" = \"%v\"", key, val))
			}
			return fmt.Sprintf("{ %s }", strings.Join(entries, ", "))
		}
		return "null"

	case "list(string)", "set(string)":
		// Render lists or sets
		if listVal, ok := value.([]interface{}); ok {
			var items []string
			for _, item := range listVal {
				items = append(items, fmt.Sprintf("\"%v\"", item))
			}
			if varType == "set(string)" {
				return fmt.Sprintf("toset([%s])", strings.Join(items, ", "))
			}
			return fmt.Sprintf("[%s]", strings.Join(items, ", "))
		}
		return "[]"

	case "object", "tuple":
		// Handle objects or tuples
		if mapVal, ok := value.(map[string]interface{}); ok {
			var entries []string
			for key, val := range mapVal {
				entries = append(entries, fmt.Sprintf("\"%s\" = %v", key, val))
			}
			return fmt.Sprintf("{ %s }", strings.Join(entries, ", "))
		}
		return "{}"

	default:
		// Default fallback for unknown types
		return fmt.Sprintf("%v", value)
	}
}

// FormatDefault formats the default value of a variable
func FormatDefault(varDef models.Variable) string {
	switch varDef.Type {
	case "bool", "number":
		return fmt.Sprintf("%v", varDef.Default)
	case "string":
		// Check if default is an expression
		if expr, ok := varDef.Default.(string); ok && strings.HasPrefix(expr, "var.") {
			return expr // Expression
		}
		return fmt.Sprintf("\"%v\"", varDef.Default)
	case "list(string)", "set(string)":
		list, ok := varDef.Default.([]interface{})
		if !ok {
			return "[]"
		}
		var items []string
		for _, item := range list {
			items = append(items, fmt.Sprintf("\"%v\"", item))
		}
		if varDef.Type == "set(string)" {
			return fmt.Sprintf("toset([%s])", strings.Join(items, ", "))
		}
		return fmt.Sprintf("[%s]", strings.Join(items, ", "))
	case "map(string)":
		var entries []string
		switch v := varDef.Default.(type) {
		case map[string]interface{}:
			for key, val := range v {
				entries = append(entries, fmt.Sprintf("\"%s\" = \"%v\"", key, val))
			}
		case map[string]string:
			for key, val := range v {
				entries = append(entries, fmt.Sprintf("\"%s\" = \"%s\"", key, val))
			}
		}
		return fmt.Sprintf("{ %s }", strings.Join(entries, ", "))
	case "object({ provision_vm_agent = bool, enable_automatic_upgrades = bool })",
		"object({ publisher = string, offer = string, sku = string, version = string })",
		"object({ name = string, caching = string, create_option = string, managed_disk_type = string })":
		// Assume default is a map[string]interface{}
		objMap, ok := varDef.Default.(map[string]interface{})
		if !ok {
			return "{}"
		}
		var items []string
		for key, val := range objMap {
			switch v := val.(type) {
			case string:
				items = append(items, fmt.Sprintf("\"%s\" = \"%v\"", key, v))
			default:
				items = append(items, fmt.Sprintf("\"%s\" = %v", key, v))
			}
		}
		return fmt.Sprintf("{ %s }", strings.Join(items, ", "))
	case "tuple":
		tuple, ok := varDef.Default.([]interface{})
		if !ok {
			return "[]"
		}
		var items []string
		for _, item := range tuple {
			switch v := item.(type) {
			case string:
				items = append(items, fmt.Sprintf("\"%v\"", v))
			default:
				items = append(items, fmt.Sprintf("%v", v))
			}
		}
		return fmt.Sprintf("[%s]", strings.Join(items, ", "))
	default:
		return fmt.Sprintf("%v", varDef.Default)
	}
}

// GenerateFileFromTemplate generates a file from a template
func GenerateFileFromTemplate(templatePath, destinationPath string, data interface{}) error {
	funcMap := template.FuncMap{
		"title": cases.Title(language.Und).String,
		"add":   func(a, b int) int { return a + b },
		"toJSON": func(value interface{}) string {
			jsonString, err := ToJSON(value)
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
		"formatValue": func(value interface{}, varType string) string {
			switch varType {
			case "bool", "number":
				return fmt.Sprintf("%v", value)
			case "string":
				expr, ok := value.(string)
				if ok && strings.HasPrefix(expr, "var.") {
					return expr
				}
				return fmt.Sprintf("\"%v\"", value)
			case "list(string)", "set(string)":
				list, ok := value.([]interface{})
				if !ok {
					return "[]"
				}
				var items []string
				for _, item := range list {
					items = append(items, fmt.Sprintf("\"%v\"", item))
				}
				if varType == "set(string)" {
					return fmt.Sprintf("toset([%s])", strings.Join(items, ", "))
				}
				return fmt.Sprintf("[%s]", strings.Join(items, ", "))
			case "map(string)":
				var entries []string
				switch v := value.(type) {
				case map[string]interface{}:
					for key, val := range v {
						entries = append(entries, fmt.Sprintf("\"%s\" = \"%v\"", key, val))
					}
				case map[string]string:
					for key, val := range v {
						entries = append(entries, fmt.Sprintf("\"%s\" = \"%s\"", key, val))
					}
				}
				return fmt.Sprintf("{ %s }", strings.Join(entries, ", "))
			default:
				return fmt.Sprintf("%v", value)
			}
		},
		"formatDefault": FormatDefault, // Existing functions
		"formatType":    formatType,    // Existing functions
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
