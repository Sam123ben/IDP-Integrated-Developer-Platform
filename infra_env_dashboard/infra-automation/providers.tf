terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "~> 4.0"
    }
  }
  required_version = "1.9.8"  # Specify your required Terraform version
}

# Configure the Azure provider
provider "azurerm" {
  features {}

  # Use variables to hold authentication details
  client_id       = var.client_id
  client_secret   = var.client_secret
  subscription_id = var.subscription_id
  tenant_id       = var.tenant_id
}