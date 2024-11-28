// backend/main.go

package main

import (
	"backend/handlers"
	"backend/models"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	company := generateCmd.String("company", "", "Company name (required)")
	product := generateCmd.String("product", "", "Product name (required)")
	provider := generateCmd.String("provider", "", "Provider name (required)")
	modules := generateCmd.String("modules", "", "Comma-separated list of modules")
	customers := generateCmd.String("customers", "", "Comma-separated list of customers")

	terraformCmd := flag.NewFlagSet("terraform", flag.ExitOnError)
	tfCommand := terraformCmd.String("command", "", "Terraform command to execute (init, plan, apply, build, print)")
	tfCompany := terraformCmd.String("company", "", "Company name (required)")
	tfProduct := terraformCmd.String("product", "", "Product name (required)")
	tfProvider := terraformCmd.String("provider", "", "Provider name (required)")
	tfInfratype := terraformCmd.String("infratype", "", "Infrastructure type (prod or nonprod)")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'generate' or 'terraform' subcommands")
		os.Exit(1)
	}

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

func handleGenerateCommand(company, product, provider, modules, customers string) {
	if company == "" || product == "" || provider == "" {
		fmt.Println("Error: --company, --product, and --provider are required")
		os.Exit(1)
	}

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

	if err := handlers.GenerateTerraform(&req); err != nil {
		fmt.Printf("Error generating Terraform code: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Terraform code generated successfully")
}

func handleTerraformCommand(command, company, product, provider, infratype string) {
	if command == "" || company == "" || product == "" || provider == "" {
		fmt.Println("Error: --command, --company, --product, and --provider are required")
		os.Exit(1)
	}

	// 'init', 'build', and 'print' commands require '--infratype'
	if (command == "init" || command == "build" || command == "print") && infratype == "" {
		fmt.Println("Error: --infratype is required for 'init', 'build', and 'print' commands")
		os.Exit(1)
	}

	if command == "print" {
		printTerraformCommands(command, company, product, provider, infratype)
	} else {
		if err := runTerraformCommand(command, company, product, provider, infratype); err != nil {
			log.Fatalf("Error executing Terraform command: %v\n", err)
		}
	}
}

func runTerraformCommand(command, company, product, provider, infratype string) error {
	terraformDir := filepath.Join("output", "terraform", company, product)
	if _, err := os.Stat(terraformDir); os.IsNotExist(err) {
		return fmt.Errorf("Terraform directory %s does not exist", terraformDir)
	}

	if err := os.Chdir(terraformDir); err != nil {
		return fmt.Errorf("error changing directory to %s: %v", terraformDir, err)
	}

	switch command {
	case "init":
		backendConfig := fmt.Sprintf("backend/%s_%s.tfvars", product, strings.ToLower(infratype))
		args := []string{"init", "-no-color", "-get=true", "-force-copy", fmt.Sprintf("-backend-config=%s", backendConfig)}
		return executeCommand("terraform", args)

	case "plan":
		args := []string{"plan", "-no-color", "-input=false", "-lock=true", "-refresh=true", "-var-file=./vars.tfvars"}
		return executeCommand("terraform", args)

	case "apply":
		args := []string{"apply", "-no-color", "-input=false", "-auto-approve=true", "-lock=true", "-lock-timeout=7200s", "-refresh=true", "-var-file=./vars.tfvars"}
		return executeCommand("terraform", args)

	case "build":
		backendConfig := fmt.Sprintf("backend/%s_%s.tfvars", product, strings.ToLower(infratype))
		buildCommands := [][]string{
			{"init", "-no-color", "-get=true", "-force-copy", fmt.Sprintf("-backend-config=%s", backendConfig)},
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

func printTerraformCommands(command, company, product, provider, infratype string) {
	terraformDir := filepath.Join("output", "terraform", company, product)
	fmt.Printf("Working directory: %s\n", terraformDir)

	backendConfig := fmt.Sprintf("backend/%s_%s.tfvars", product, strings.ToLower(infratype))
	fmt.Printf("terraform init -no-color -get=true -force-copy -backend-config=%s\n", backendConfig)
	fmt.Println("terraform plan -no-color -input=false -lock=true -refresh=true -var-file=./vars.tfvars")
	fmt.Println("terraform apply -no-color -input=false -auto-approve=true -lock=true -lock-timeout=7200s -refresh=true -var-file=./vars.tfvars")
}

func executeCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("Running command: %s %s\n", command, strings.Join(args, " "))
	return cmd.Run()
}
