terraform {
  backend "azurerm" {
    resource_group_name  = "dashboard-state-rg"
    storage_account_name = "dashboardstatestg"
    container_name       = "dashboard-tfstate"  # Replace <customer name> with your actual value
    key                  = "terraform.tfstate"
  }
}