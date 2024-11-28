// backend/main.go

package main

import (
	"backend/handlers"
	"backend/models"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	company := generateCmd.String("company", "", "Company name (required)")
	product := generateCmd.String("product", "", "Product name (required)")
	provider := generateCmd.String("provider", "", "Provider name (required)")
	modules := generateCmd.String("modules", "", "Comma-separated list of modules")
	customers := generateCmd.String("customers", "", "Comma-separated list of customers")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'generate' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		generateCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

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
}
