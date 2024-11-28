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
		return err
	}

	// Filter provider data based on the input provider
	providerData := filterProviderData(config.Providers, req.Provider)
	if providerData == nil {
		return errors.New("specified provider not found in configuration")
	}

	// Resolve module dependencies
	moduleNames := make([]string, len(req.Modules))
	for i, moduleName := range req.Modules {
		moduleNames[i] = strings.TrimSpace(moduleName)
	}
	modules, rootVariables := resolveModuleDependencies(moduleNames, config.Modules)

	basePath := filepath.Join("output", "terraform", req.OrganisationName)

	// Generate module files
	if err := generateModuleFiles(basePath, modules, req.Provider); err != nil {
		return err
	}

	if len(req.Customers) > 0 {
		return processCustomers(req, config, basePath, providerData, modules, rootVariables)
	}

	// Process for a single product
	productPath := filepath.Join(basePath, req.ProductName)
	if err := utils.CreateDirectories([]string{filepath.Join(productPath, "backend")}); err != nil {
		return err
	}

	if err := generateProductFiles(req, config, productPath, providerData, modules, rootVariables); err != nil {
		return err
	}

	return nil
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
func resolveModuleDependencies(requestedModules []string, availableModules []models.Module) ([]models.Module, map[string]models.Variable) {
	moduleMap := make(map[string]models.Module)
	for _, module := range availableModules {
		moduleMap[module.ModuleName] = module
	}

	visited := make(map[string]bool)
	var resolved []models.Module
	rootVariables := make(map[string]models.Variable)

	var resolve func(string)
	resolve = func(moduleName string) {
		if visited[moduleName] {
			return
		}
		visited[moduleName] = true

		module, exists := moduleMap[moduleName]
		if !exists {
			fmt.Printf("Warning: Module '%s' not found in available modules.\n", moduleName)
			return
		}

		// Resolve dependencies first
		for _, dependency := range module.DependsOn {
			resolve(dependency)
		}

		// Collect variables for root variables.tf
		for _, varDef := range module.Variables {
			// Assume that variables starting with "var." are root variables
			if strings.HasPrefix(varDef.Value, "var.") {
				rootVarName := strings.TrimPrefix(varDef.Value, "var.")
				if _, exists := rootVariables[rootVarName]; !exists {
					rootVariables[rootVarName] = models.Variable{
						Type:        varDef.Type,
						Description: varDef.Description,
					}
				}
			}
		}

		// Add the current module to the resolved list
		resolved = append(resolved, module)
	}

	for _, moduleName := range requestedModules {
		resolve(moduleName)
	}

	return resolved, rootVariables
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

		// Prepare list of files to generate
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
				return err
			}
		}
	}
	return nil
}

// generateProductFiles creates Terraform files for a single product.
func generateProductFiles(req *models.GenerateRequest, config *models.Config, productPath string, provider *models.Provider, modules []models.Module, rootVariables map[string]models.Variable) error {
	data := prepareTemplateData(req, config, provider, "", modules, rootVariables)

	// Generate files
	if err := generateTerraformFiles(productPath, data, req.Provider, req.ProductName); err != nil {
		return err
	}

	// Generate backend tfvars files
	return generateBackendTfvarsFiles(productPath, data, req.ProductName)
}

// processCustomers generates Terraform files for multiple customers.
func processCustomers(req *models.GenerateRequest, config *models.Config, basePath string, provider *models.Provider, modules []models.Module, rootVariables map[string]models.Variable) error {
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
		if err := generateCustomerFiles(req, config, customerPath, customer, provider, modules, rootVariables); err != nil {
			return err
		}
	}
	return nil
}

// generateCustomerFiles creates Terraform files for a single customer.
func generateCustomerFiles(req *models.GenerateRequest, config *models.Config, customerPath, customerName string, provider *models.Provider, modules []models.Module, rootVariables map[string]models.Variable) error {
	data := prepareTemplateData(req, config, provider, customerName, modules, rootVariables)

	// Generate files
	if err := generateTerraformFiles(customerPath, data, req.Provider, customerName); err != nil {
		return err
	}

	// Generate backend and vars tfvars files
	return generateBackendAndVarsTfvarsFiles(customerPath, data, customerName)
}

// prepareTemplateData prepares data for the templates.
func prepareTemplateData(req *models.GenerateRequest, config *models.Config, provider *models.Provider, customerName string, modules []models.Module, rootVariables map[string]models.Variable) map[string]interface{} {
	return map[string]interface{}{
		"Provider":         provider,
		"TerraformVersion": config.TerraformVersion,
		"Modules":          modules,
		"OrganisationName": req.OrganisationName,
		"ProductName":      req.ProductName,
		"CustomerName":     customerName,
		"Region":           config.Region,
		"Environment":      config.Environment,
		"Backend":          config.Backend,
		"Variables":        rootVariables,
	}
}

// generateTerraformFiles creates Terraform files like providers.tf, main.tf, variables.tf, and vars.tfvars.
func generateTerraformFiles(path string, data map[string]interface{}, provider, entityName string) error {
	// Generate root variables.tf
	if err := generateRootVariablesFile(path, data["Variables"].(map[string]models.Variable)); err != nil {
		return err
	}

	files := []struct {
		Template string
		Dest     string
	}{
		{Template: filepath.Join("templates", "generic", "providers.tf.tmpl"), Dest: filepath.Join(path, "providers.tf")},
		{Template: filepath.Join("templates", provider, "main.tf.tmpl"), Dest: filepath.Join(path, "main.tf")},
		{Template: filepath.Join("templates", "generic", "vars.tfvars.tmpl"), Dest: filepath.Join(path, "vars.tfvars")},
	}

	for _, file := range files {
		if err := utils.GenerateFileFromTemplate(file.Template, file.Dest, data); err != nil {
			return err
		}
	}
	return nil
}

// generateRootVariablesFile generates the root variables.tf file.
func generateRootVariablesFile(path string, variables map[string]models.Variable) error {
	data := map[string]interface{}{
		"Variables": variables,
	}

	destPath := filepath.Join(path, "variables.tf")
	templatePath := filepath.Join("templates", "generic", "variables.tf.tmpl")

	return utils.GenerateFileFromTemplate(templatePath, destPath, data)
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
