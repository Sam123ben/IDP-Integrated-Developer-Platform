# App Service Plan for both frontend and backend apps
resource "azurerm_service_plan" "app_service_plan" {
  name                = var.app_service_plan_name
  location            = var.location
  resource_group_name = var.resource_group_name
  os_type             = "Linux"  # Explicitly set to Linux for Docker containers

  sku_name = "B3"  # Basic tier with 1.75 GB memory and 1 vCPU

  tags = var.tags
}

# Backend App Service with database connection via private endpoint
resource "azurerm_linux_web_app" "backend_app" {
  name                = var.backend_app_name
  location            = var.location
  resource_group_name = var.resource_group_name
  service_plan_id     = azurerm_service_plan.app_service_plan.id

  app_settings = {
    WEBSITES_PORT = "8080"
    DATABASE_URL  = var.database_url  # Database URL for private endpoint access
  }

  site_config {
    # Application stack specifying the Docker image
    application_stack {
      docker_image_name   = "sam123ben/infra-dashboard-backend:latest"
      docker_registry_url = "https://index.docker.io/v1/"  # URL for Docker Hub or private registry
    }
  }

  depends_on = [azurerm_service_plan.app_service_plan]
  tags       = var.tags
}

# Frontend App Service listening on port 3000
resource "azurerm_linux_web_app" "frontend_app" {
  name                = var.frontend_app_name
  location            = var.location
  resource_group_name = var.resource_group_name
  service_plan_id     = azurerm_service_plan.app_service_plan.id

  app_settings = {
    WEBSITES_PORT = "3000"
  }

  site_config {
    # Application stack specifying the Docker image
    application_stack {
      docker_image_name   = "sam123ben/infra-dashboard:latest"
      docker_registry_url = "https://index.docker.io/v1/"  # URL for Docker Hub or private registry
    }
  }

  depends_on = [azurerm_service_plan.app_service_plan]
  tags       = var.tags
}