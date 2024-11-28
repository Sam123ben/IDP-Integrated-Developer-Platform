terraform {
  required_providers {
    
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 4.10.0"
    }
    
  }
  required_version = "1.0.0"
}


# Configure the azurerm provider
provider "azurerm" {
  
  # Authentication variables
  
  client_id = "000000-0000-0000-0000-000000000000"
  
  client_secret = "000000-0000-0000-0000-000000000000"
  
  subscription_id = "000000-0000-0000-0000-000000000000"
  
  tenant_id = "000000-0000-0000-0000-000000000000"
  
  
}
