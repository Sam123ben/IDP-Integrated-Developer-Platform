# App Service Plan for both frontend and backend apps
resource "azurerm_app_service_plan" "app_service_plan" {
  name                = var.app_service_plan_name
  location            = var.location
  resource_group_name = var.resource_group_name
  sku {
    tier = "Standard"
    size = "S1"
  }
}

# Backend App Service with database connection via private endpoint
resource "azurerm_app_service" "backend_app" {
  name                = var.backend_app_name
  location            = var.location
  resource_group_name = var.resource_group_name
  app_service_plan_id = azurerm_app_service_plan.app_service_plan.id

  app_settings = {
    WEBSITES_PORT = "8080"
    DATABASE_URL  = var.database_url  # Assuming database_url is private endpoint accessible URL
  }

  site_config {
    linux_fx_version = "DOCKER|sam123ben/infra-dashboard-backend:latest"
  }
}

# Frontend App Service listening on port 3000
resource "azurerm_app_service" "frontend_app" {
  name                = var.frontend_app_name
  location            = var.location
  resource_group_name = var.resource_group_name
  app_service_plan_id = azurerm_app_service_plan.app_service_plan.id

  app_settings = {
    WEBSITES_PORT = "3000"
  }

  site_config {
    linux_fx_version = "DOCKER|sam123ben/infra-dashboard:latest"
  }
}