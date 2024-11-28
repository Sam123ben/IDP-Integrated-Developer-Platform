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
	tfCommand := terraformCmd.String("command", "", "Terraform command to execute (init, plan, apply, build)")
	tfCompany := terraformCmd.String("company", "", "Company name (required)")
	tfProduct := terraformCmd.String("product", "", "Product name (required)")
	tfProvider := terraformCmd.String("provider", "", "Provider name (required)")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'generate' or 'terraform' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		generateCmd.Parse(os.Args[2:])
		if generateCmd.Parsed() {
			if *company == "" || *product == "" || *provider == "" {
				fmt.Println("Error: --company, --product, and --provider are required")
				generateCmd.Usage()
				os.Exit(1)
			}

			// Prepare the request data
			req := models.GenerateRequest{
				OrganisationName: *company,
				ProductName:      *product,
				Provider:         *provider,
				Modules:          []string{},
			}

			// Handle modules
			if *modules != "" {
				moduleNames := strings.Split(*modules, ",")
				for _, moduleName := range moduleNames {
					moduleName = strings.TrimSpace(moduleName)
					req.Modules = append(req.Modules, moduleName)
				}
			}

			// Handle customers
			if *customers != "" {
				req.Customers = strings.Split(*customers, ",")
				for i, customer := range req.Customers {
					req.Customers[i] = strings.TrimSpace(customer)
				}
			}

			// Call the generation logic
			err := handlers.GenerateTerraform(&req)
			if err != nil {
				fmt.Printf("Error generating Terraform code: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Terraform code generated successfully")
		}

	case "terraform":
		terraformCmd.Parse(os.Args[2:])
		if terraformCmd.Parsed() {
			if *tfCommand == "" || *tfCompany == "" || *tfProduct == "" || *tfProvider == "" {
				fmt.Println("Error: --command, --company, --product, and --provider are required")
				terraformCmd.Usage()
				os.Exit(1)
			}

			// Run the Terraform command
			err := runTerraformCommand(*tfCommand, *tfCompany, *tfProduct, *tfProvider)
			if err != nil {
				log.Fatalf("Error executing Terraform command: %v\n", err)
			}
		}

	default:
		fmt.Println("Expected 'generate' or 'terraform' subcommands")
		os.Exit(1)
	}
}

// runTerraformCommand executes Terraform commands (init, plan, apply, build)
func runTerraformCommand(command, company, product, provider string) error {
	// Determine the working directory
	terraformDir := filepath.Join("output", "terraform", company, product)
	if _, err := os.Stat(terraformDir); os.IsNotExist(err) {
		return fmt.Errorf("Terraform directory %s does not exist", terraformDir)
	}

	// Change to the Terraform directory
	if err := os.Chdir(terraformDir); err != nil {
		return fmt.Errorf("error changing directory to %s: %v", terraformDir, err)
	}

	// Map of commands to Terraform actions
	tfCommands := map[string][]string{
		"init":  {"init"},
		"plan":  {"plan"},
		"apply": {"apply", "-auto-approve"},
		"build": {"init", "plan", "apply", "-auto-approve"},
	}

	// Get the commands for the specified action
	actions, ok := tfCommands[command]
	if !ok {
		return fmt.Errorf("unsupported Terraform command: %s", command)
	}

	// Execute the commands in sequence
	for i := 0; i < len(actions); i++ {
		cmd := exec.Command("terraform", actions[i])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Printf("Running Terraform command: terraform %s\n", actions[i])

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("error running Terraform command '%s': %v", actions[i], err)
		}
	}

	return nil
}
