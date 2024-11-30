// backend/utils/terraform_utils.go

package utils

import (
	"backend/models"
	"fmt"
	"strings"
)

// FilterProviderData filters provider details based on the specified provider name.
func FilterProviderData(providers []models.Provider, providerName string) *models.Provider {
	aliases := map[string]string{
		"azure":   "azurerm",
		"aws":     "aws",
		"gcp":     "google",
		"azurerm": "azurerm",
		"google":  "google",
	}

	normalizedProvider := aliases[strings.ToLower(providerName)]

	for _, provider := range providers {
		if strings.EqualFold(provider.Name, normalizedProvider) {
			return &provider
		}
	}
	return nil
}

// ResolveModuleDependencies resolves all dependencies for the requested modules.
func ResolveModuleDependencies(requestedModules []string, availableModules []models.Module) ([]models.Module, error) {
	moduleMap := make(map[string]models.Module)
	for _, module := range availableModules {
		moduleMap[module.ModuleName] = module
	}

	visited := make(map[string]bool)
	var resolved []models.Module

	var resolve func(string) error
	resolve = func(moduleName string) error {
		if visited[moduleName] {
			return nil
		}
		visited[moduleName] = true

		module, exists := moduleMap[moduleName]
		if !exists {
			return fmt.Errorf("module '%s' not found in available modules", moduleName)
		}

		// Resolve dependencies first
		for _, dependency := range module.DependsOn {
			if err := resolve(dependency); err != nil {
				return err
			}
		}

		// Add the current module to the resolved list
		resolved = append(resolved, module)
		return nil
	}

	for _, moduleName := range requestedModules {
		if err := resolve(moduleName); err != nil {
			return nil, err
		}
	}

	return resolved, nil
}

// ExtractModuleVariables processes the module variables for a given module.
func ExtractModuleVariables(module models.Module) map[string]interface{} {
	variables := make(map[string]interface{})
	for name, varDef := range module.Variables {
		variables[name] = map[string]interface{}{
			"description": varDef.Description,
			"type":        varDef.Type,
			"default":     varDef.Default,
			"sensitive":   varDef.Sensitive,
			"validation":  varDef.Validation,
		}
	}
	return variables
}
