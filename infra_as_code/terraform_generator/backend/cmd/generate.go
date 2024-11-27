// backend/cmd/generate.go

package cmd

import (
	"fmt"
	"path/filepath"

	"backend/services/config"
	"backend/services/generator"
	"backend/services/structure"

	"github.com/spf13/cobra"
)

var (
	customer   string
	product    string
	configPath string
	outputDir  string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Terraform templates",
	Long:  "Generate Terraform provider.tf and backend.tf dynamically based on the configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		if customer == "" || product == "" {
			fmt.Println("Error: --customer and --product are required")
			return
		}

		// Load configuration
		configData, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Printf("Error loading configuration: %v\n", err)
			return
		}

		// Get provider details for Azure
		providerConfig, exists := configData.Providers["azure"]
		if !exists {
			fmt.Println("Error: Azure configuration not found in the config file")
			return
		}

		// Ensure folder structure
		baseOutputDir := filepath.Join(outputDir, "azure", customer, product)
		err = structure.EnsureFolderStructure(baseOutputDir)
		if err != nil {
			fmt.Printf("Error creating folder structure: %v\n", err)
			return
		}

		// Generate provider.tf
		err = generator.GenerateProviderFile(baseOutputDir, providerConfig)
		if err != nil {
			fmt.Printf("Error generating provider.tf: %v\n", err)
			return
		}

		// Generate backend.tf
		err = generator.GenerateBackendFile(baseOutputDir, providerConfig.Backend)
		if err != nil {
			fmt.Printf("Error generating backend.tf: %v\n", err)
			return
		}

		fmt.Printf("Terraform templates successfully generated at %s\n", baseOutputDir)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&customer, "customer", "u", "", "Customer name")
	generateCmd.Flags().StringVarP(&product, "product", "r", "", "Product name")
	generateCmd.Flags().StringVarP(&configPath, "config", "c", "configs/terraform-generator.yaml", "Path to configuration")
	generateCmd.Flags().StringVarP(&outputDir, "output", "o", "output", "Output directory")
}
