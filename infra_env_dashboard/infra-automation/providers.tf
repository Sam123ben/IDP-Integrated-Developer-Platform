# Specify the required provider for Azure
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0"  # Specify the version you need
    }
  }
  required_version = ">= 1.0"  # Specify your required Terraform version
}

# Configure the Azure provider
provider "azurerm" {
  features {}  # Required block for azurerm provider configuration
  # Uncomment and replace with your values for inline authentication (only for dev use)
  client_id       = "your-client-id"
  client_secret   = "your-client-secret"
  subscription_id = "your-subscription-id"
  tenant_id       = "your-tenant-id"
}