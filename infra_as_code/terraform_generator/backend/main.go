// backend/main.go

package main

import (
	"backend/handlers"
	"backend/models"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "generate" {
		// Command-line mode
		handleCLI()
	} else {
		// Web server mode
		http.HandleFunc("/generate", handlers.GenerateTerraformHandler)
		log.Println("Server is running on port 8080...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}
}

func handleCLI() {
	// Define flags
	company := flag.String("company", "", "Company name")
	product := flag.String("product", "", "Product name")
	customers := flag.String("customers", "", "Comma-separated list of customer names (optional)")
	provider := flag.String("provider", "", "Cloud provider (aws, azure, gcp)")
	modules := flag.String("modules", "", "Comma-separated list of modules")

	// Parse flags
	flag.CommandLine.Parse(os.Args[2:])

	// Validate required flags
	if *company == "" || *product == "" || *provider == "" {
		fmt.Println("Error: --company, --product, and --provider are required")
		flag.Usage()
		os.Exit(1)
	}

	// Prepare the request data
	req := &models.GenerateRequest{
		OrganisationName: *company,
		ProductName:      *product,
		Provider:         *provider,
		Modules:          []models.Module{},
	}

	// Handle modules
	if *modules != "" {
		moduleNames := strings.Split(*modules, ",")
		for _, moduleName := range moduleNames {
			moduleName = strings.TrimSpace(moduleName)
			module := models.Module{
				ModuleName: moduleName,
				Source:     fmt.Sprintf("./modules/%s", moduleName),
				Variables:  make(map[string]string),
			}
			req.Modules = append(req.Modules, module)
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
	err := handlers.GenerateTerraform(req)
	if err != nil {
		fmt.Printf("Error generating Terraform code: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Terraform code generated successfully")
}
