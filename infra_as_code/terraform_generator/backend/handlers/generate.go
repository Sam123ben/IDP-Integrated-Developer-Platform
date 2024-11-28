// backend/handlers/generate.go

package handlers

import (
	"backend/models"
	"backend/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

// GenerateTerraformHandler handles HTTP requests to generate Terraform files.
func GenerateTerraformHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GenerateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := GenerateTerraform(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Terraform code generated successfully"))
}

// GenerateTerraform processes the request to generate Terraform files.
func GenerateTerraform(req *models.GenerateRequest) error {
	if req.OrganisationName == "" || req.ProductName == "" || req.Provider == "" {
		return errors.New("organisation_name, product_name, and provider are required")
	}

	// Load configuration from terraform-generator.json
	config, err := utils.LoadConfig("configs/terraform-generator.json")
	if err != nil {
		return fmt.Errorf("error loading configuration: %w", err)
	}

	// Filter provider data based on the input provider
	providerData := filterProviderData(config.Providers, req.Provider)
	if providerData == nil {
		return fmt.Errorf("specified provider '%s' not found in configuration", req.Provider)
	}

	// Resolve module dependencies
	modules, err := resolveModuleDependencies(req.Modules, config.Modules)
	if err != nil {
		return fmt.Errorf("error resolving module dependencies: %w", err)
	}

	// Update basePath to include 'output' directory
	basePath := filepath.Join("output", "terraform", req.OrganisationName)

	// Generate module files
	if err := generateModuleFiles(basePath, modules, req.Provider); err != nil {
		return fmt.Errorf("error generating module files: %w", err)
	}

	// Generate files for a single product or customers
	if len(req.Customers) > 0 {
		return processCustomers(req, config, basePath, providerData, modules)
	}

	// Generate product-specific files
	productPath := filepath.Join(basePath, req.ProductName)
	if err := utils.CreateDirectories([]string{filepath.Join(productPath, "backend")}); err != nil {
		return fmt.Errorf("error creating directories for product: %w", err)
	}

	return generateProductFiles(req, config, productPath, providerData, modules)
}

// filterProviderData filters provider details based on the specified provider name.
func filterProviderData(providers []models.Provider, providerName string) *models.Provider {
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

// resolveModuleDependencies resolves all dependencies for the requested modules.
func resolveModuleDependencies(requestedModules []string, availableModules []models.Module) ([]models.Module, error) {
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

// generateModuleFiles creates module directories and files.
func generateModuleFiles(basePath string, modules []models.Module, provider string) error {
	for _, module := range modules {
		modulePath := filepath.Join(basePath, "modules", module.ModuleName)
		if err := utils.CreateDirectories([]string{modulePath}); err != nil {
			return err
		}

		data := map[string]interface{}{
			"Module":       module,
			"ResourceName": module.ModuleName,
		}

		files := []struct {
			Template string
			Dest     string
		}{
			{
				Template: filepath.Join("templates", provider, module.ModuleName, "main.tf.tmpl"),
				Dest:     filepath.Join(modulePath, "main.tf"),
			},
			{
				Template: filepath.Join("templates", provider, module.ModuleName, "variables.tf.tmpl"),
				Dest:     filepath.Join(modulePath, "variables.tf"),
			},
		}

		// Include outputs.tf if outputs are defined
		if len(module.Outputs) > 0 {
			files = append(files, struct {
				Template string
				Dest     string
			}{
				Template: filepath.Join("templates", provider, module.ModuleName, "outputs.tf.tmpl"),
				Dest:     filepath.Join(modulePath, "outputs.tf"),
			})
		}

		// Generate files
		for _, file := range files {
			if err := utils.GenerateFileFromTemplate(file.Template, file.Dest, data); err != nil {
				return fmt.Errorf("error generating file %s: %w", file.Dest, err)
			}
		}
	}
	return nil
}

// generateProductFiles creates Terraform files for a single product.
func generateProductFiles(req *models.GenerateRequest, config *models.Config, productPath string, provider *models.Provider, modules []models.Module) error {
	data := prepareTemplateData(req, config, provider, "", modules)

	// Generate files
	if err := generateTerraformFiles(productPath, data, req.Provider, req.ProductName); err != nil {
		return err
	}

	// Generate backend tfvars files
	return generateBackendTfvarsFiles(productPath, data, req.ProductName)
}

// processCustomers generates Terraform files for multiple customers.
func processCustomers(req *models.GenerateRequest, config *models.Config, basePath string, provider *models.Provider, modules []models.Module) error {
	for _, customer := range req.Customers {
		customer = strings.TrimSpace(customer)
		customerPath := filepath.Join(basePath, customer)
		paths := []string{
			filepath.Join(customerPath, "backend"),
			filepath.Join(customerPath, "vars"),
		}

		// Create directories
		if err := utils.CreateDirectories(paths); err != nil {
			return err
		}

		// Generate files for the customer
		if err := generateCustomerFiles(req, config, customerPath, customer, provider, modules); err != nil {
			return err
		}
	}
	return nil
}

// generateCustomerFiles creates Terraform files for a single customer.
func generateCustomerFiles(req *models.GenerateRequest, config *models.Config, customerPath, customerName string, provider *models.Provider, modules []models.Module) error {
	data := prepareTemplateData(req, config, provider, customerName, modules)

	// Generate files
	if err := generateTerraformFiles(customerPath, data, req.Provider, customerName); err != nil {
		return err
	}

	// Generate backend and vars tfvars files
	return generateBackendAndVarsTfvarsFiles(customerPath, data, customerName)
}

// prepareTemplateData prepares data for the templates.
func prepareTemplateData(req *models.GenerateRequest, config *models.Config, provider *models.Provider, customerName string, modules []models.Module) map[string]interface{} {
	// Extract generic variables from config
	genericVariables := config.Variables

	// Prepare module variables for module calls in main.tf
	moduleVariables := make(map[string]map[string]string)
	for _, module := range modules {
		vars := make(map[string]string)
		for varName, varDef := range module.Variables {
			vars[varName] = fmt.Sprintf("%v", varDef.Value) // Convert value to string
		}
		moduleVariables[module.ModuleName] = vars
	}

	return map[string]interface{}{
		"Provider":         provider,
		"TerraformVersion": config.TerraformVersion,
		"Modules":          modules,
		"ModuleVariables":  moduleVariables, // Include populated module variables
		"OrganisationName": req.OrganisationName,
		"ProductName":      req.ProductName,
		"CustomerName":     customerName,
		"Region":           config.Region,
		"Environment":      config.Environment,
		"Backend":          config.Backend,
		"Variables":        genericVariables,
	}
}

// generateTerraformFiles creates Terraform files like providers.tf, main.tf, variables.tf, and vars.tfvars.
func generateTerraformFiles(path string, data map[string]interface{}, provider, entityName string) error {
	files := []struct {
		Template string
		Dest     string
	}{
		{Template: filepath.Join("templates", "generic", "providers.tf.tmpl"), Dest: filepath.Join(path, "providers.tf")},
		{Template: filepath.Join("templates", provider, "main.tf.tmpl"), Dest: filepath.Join(path, "main.tf")},
		{Template: filepath.Join("templates", "generic", "variables.tf.tmpl"), Dest: filepath.Join(path, "variables.tf")},
		{Template: filepath.Join("templates", "generic", "vars.tfvars.tmpl"), Dest: filepath.Join(path, "vars.tfvars")},
	}

	for _, file := range files {
		if err := utils.GenerateFileFromTemplate(file.Template, file.Dest, data); err != nil {
			return err
		}
	}
	return nil
}

// generateBackendTfvarsFiles creates backend tfvars files for a product.
func generateBackendTfvarsFiles(path string, data map[string]interface{}, productName string) error {
	environments := []string{"nonprod", "prod"}
	for _, env := range environments {
		data["Environment"] = env
		filename := productName + "_" + env + ".tfvars"
		destPath := filepath.Join(path, "backend", filename)
		if err := utils.GenerateFileFromTemplate(filepath.Join("templates", "generic", "backend.tfvars.tmpl"), destPath, data); err != nil {
			return err
		}
	}
	return nil
}

// generateBackendAndVarsTfvarsFiles creates backend and vars tfvars files for a customer.
func generateBackendAndVarsTfvarsFiles(path string, data map[string]interface{}, customerName string) error {
	environments := []string{"nonprod", "prod"}
	for _, env := range environments {
		data["Environment"] = env
		files := []struct {
			Template string
			Dest     string
		}{
			{Template: filepath.Join("templates", "generic", "backend.tfvars.tmpl"), Dest: filepath.Join(path, "backend", customerName+"_"+env+".tfvars")},
			{Template: filepath.Join("templates", "generic", "vars.tfvars.tmpl"), Dest: filepath.Join(path, "vars", customerName+"_"+env+".tfvars")},
		}

		for _, file := range files {
			if err := utils.GenerateFileFromTemplate(file.Template, file.Dest, data); err != nil {
				return err
			}
		}
	}
	return nil
}
