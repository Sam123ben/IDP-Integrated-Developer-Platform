// backend/main.go

package main

import (
	"backend/models"
	"backend/services"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Define subcommands
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	terraformCmd := flag.NewFlagSet("terraform", flag.ExitOnError)

	// Define flags for 'generate' subcommand
	company := generateCmd.String("company", "", "Company name (required)")
	product := generateCmd.String("product", "", "Product name (required)")
	provider := generateCmd.String("provider", "", "Provider name (required)")
	modules := generateCmd.String("modules", "", "Comma-separated list of modules")
	customers := generateCmd.String("customers", "", "Comma-separated list of customers")

	// Define flags for 'terraform' subcommand
	tfCommand := terraformCmd.String("command", "", "Terraform command to execute (init, validate, plan, apply, build, destroy, print)")
	tfCompany := terraformCmd.String("company", "", "Company name (required)")
	tfProduct := terraformCmd.String("product", "", "Product name (required)")
	tfProvider := terraformCmd.String("provider", "", "Provider name (required)")
	tfInfratype := terraformCmd.String("infratype", "", "Infrastructure type (prod or nonprod)")

	// Ensure a subcommand is provided
	if len(os.Args) < 2 {
		fmt.Println("Expected 'generate' or 'terraform' subcommands")
		os.Exit(1)
	}

	// Parse subcommands
	switch os.Args[1] {
	case "generate":
		generateCmd.Parse(os.Args[2:])
		if generateCmd.Parsed() {
			handleGenerateCommand(*company, *product, *provider, *modules, *customers)
		}

	case "terraform":
		terraformCmd.Parse(os.Args[2:])
		if terraformCmd.Parsed() {
			handleTerraformCommand(*tfCommand, *tfCompany, *tfProduct, *tfProvider, *tfInfratype)
		}

	default:
		fmt.Println("Expected 'generate' or 'terraform' subcommands")
		os.Exit(1)
	}
}

// handleGenerateCommand processes the 'generate' subcommand
func handleGenerateCommand(company, product, provider, modules, customers string) {
	// Validate required flags
	if company == "" || product == "" || provider == "" {
		fmt.Println("Error: --company, --product, and --provider are required")
		os.Exit(1)
	}

	// Create a GenerateRequest
	req := models.GenerateRequest{
		OrganisationName: company,
		ProductName:      product,
		Provider:         provider,
		Modules:          []string{},
	}

	// Handle modules
	if modules != "" {
		moduleNames := strings.Split(modules, ",")
		for _, moduleName := range moduleNames {
			req.Modules = append(req.Modules, strings.TrimSpace(moduleName))
		}
	}

	// Handle customers
	if customers != "" {
		req.Customers = strings.Split(customers, ",")
		for i := range req.Customers {
			req.Customers[i] = strings.TrimSpace(req.Customers[i])
		}
	}

	// Generate Terraform code
	if err := services.GenerateTerraform(&req); err != nil {
		fmt.Printf("Error generating Terraform code: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Terraform code generated successfully")
}

// handleTerraformCommand processes the 'terraform' subcommand
func handleTerraformCommand(command, company, product, provider, infratype string) {
	// Validate required flags
	if command == "" || company == "" || product == "" || provider == "" {
		fmt.Println("Error: --command, --company, --product, and --provider are required")
		os.Exit(1)
	}

	// Certain commands require 'infratype'
	if (command == "init" || command == "build" || command == "print" || command == "validate") && infratype == "" {
		fmt.Println("Error: --infratype is required for 'init', 'build', 'print', and 'validate' commands")
		os.Exit(1)
	}

	// Handle 'print' command separately
	if command == "print" {
		printTerraformCommands(command, company, product, provider, infratype)
	} else {
		// Execute other Terraform commands
		if err := runTerraformCommand(command, company, product, provider, infratype); err != nil {
			log.Fatalf("Error executing Terraform command: %v\n", err)
		}
	}
}

// runTerraformCommand executes the specified Terraform command
func runTerraformCommand(command, company, product, provider, infratype string) error {
	terraformDir := filepath.Join("output", "terraform", company, product)
	if _, err := os.Stat(terraformDir); os.IsNotExist(err) {
		return fmt.Errorf("Terraform directory %s does not exist", terraformDir)
	}

	// Change to the Terraform directory
	if err := os.Chdir(terraformDir); err != nil {
		return fmt.Errorf("error changing directory to %s: %v", terraformDir, err)
	}

	switch command {
	case "init":
		// Removed -backend-config
		args := []string{"init", "-no-color", "-get=true", "-force-copy"}
		return executeCommand("terraform", args)

	case "validate":
		args := []string{"validate", "-no-color"}
		return executeCommand("terraform", args)

	case "plan":
		args := []string{"plan", "-no-color", "-input=false", "-lock=true", "-refresh=true", "-var-file=./vars.tfvars"}
		return executeCommand("terraform", args)

	case "apply":
		args := []string{"apply", "-no-color", "-input=false", "-auto-approve=true", "-lock=true", "-lock-timeout=7200s", "-refresh=true", "-var-file=./vars.tfvars"}
		return executeCommand("terraform", args)

	case "destroy":
		args := []string{"destroy", "-no-color", "-auto-approve=true", "-var-file=./vars.tfvars"}
		return executeCommand("terraform", args)

	case "build":
		// Sequentially execute init, validate, plan, and apply
		buildCommands := [][]string{
			{"init", "-no-color", "-get=true", "-force-copy"},
			{"validate", "-no-color"},
			{"plan", "-no-color", "-input=false", "-lock=true", "-refresh=true", "-var-file=./vars.tfvars"},
			{"apply", "-no-color", "-input=false", "-auto-approve=true", "-lock=true", "-lock-timeout=7200s", "-refresh=true", "-var-file=./vars.tfvars"},
		}

		for _, args := range buildCommands {
			if err := executeCommand("terraform", args); err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("unsupported Terraform command: %s", command)
	}

	return nil
}

// printTerraformCommands prints the Terraform commands without executing them
func printTerraformCommands(command, company, product, provider, infratype string) {
	terraformDir := filepath.Join("output", "terraform", company, product)
	fmt.Printf("Working directory: %s\n", terraformDir)

	switch command {
	case "init":
		fmt.Println("terraform init -no-color -get=true -force-copy")
	case "validate":
		fmt.Println("terraform validate -no-color")
	case "plan":
		fmt.Println("terraform plan -no-color -input=false -lock=true -refresh=true -var-file=./vars.tfvars")
	case "apply":
		fmt.Println("terraform apply -no-color -input=false -auto-approve=true -lock=true -lock-timeout=7200s -refresh=true -var-file=./vars.tfvars")
	case "destroy":
		fmt.Println("terraform destroy -no-color -auto-approve=true -var-file=./vars.tfvars")
	case "build":
		fmt.Println("terraform init -no-color -get=true -force-copy")
		fmt.Println("terraform validate -no-color")
		fmt.Println("terraform plan -no-color -input=false -lock=true -refresh=true -var-file=./vars.tfvars")
		fmt.Println("terraform apply -no-color -input=false -auto-approve=true -lock=true -lock-timeout=7200s -refresh=true -var-file=./vars.tfvars")
	default:
		fmt.Printf("Unsupported command: %s\n", command)
	}
}

// executeCommand runs a shell command and streams its output
func executeCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("Running command: %s %s\n", command, strings.Join(args, " "))
	return cmd.Run()
}
