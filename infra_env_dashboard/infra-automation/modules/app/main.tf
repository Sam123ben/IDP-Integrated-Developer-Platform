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
    WEBSITES_PORT               = "8080"
    DATABASE_URL                = var.database_url  # Database URL for private endpoint access
    DOCKER_REGISTRY_SERVER_URL  = "https://index.docker.io/v1"
    SWAGGER_HOST                = "${var.backend_app_name}.azurewebsites.net"
    PORT                        = "443"
  }

  site_config {
    # Application stack specifying the Docker image
    application_stack {
      docker_image_name = "sam123ben/infra-dashboard-backend:latest"
    }
    cors {
      allowed_origins = ["*"]
    }
  }

  connection_string {
    name  = "DATABASE_URL"
    type  = "PostgreSQL"
    value = "postgres://${var.db_admin_user}:${var.db_admin_password}@${var.db_server_name}:5432/${var.db_name}"
  }

  client_affinity_enabled          = false
  https_only                       = true
  identity {
    type = "SystemAssigned"
  }

  tags       = var.tags
  depends_on = [azurerm_service_plan.app_service_plan]
}

# Frontend App Service listening on port 3000
resource "azurerm_linux_web_app" "frontend_app" {
  name                = var.frontend_app_name
  location            = var.location
  resource_group_name = var.resource_group_name
  service_plan_id     = azurerm_service_plan.app_service_plan.id

  app_settings = {
    WEBSITES_PORT       = "3000"
    REACT_APP_API_HOST  = "https://${azurerm_linux_web_app.backend_app.default_hostname}"
    REACT_APP_API_PORT  = "443"
  }

  site_config {
    # Application stack specifying the Docker image
    application_stack {
      docker_image_name = "sam123ben/infra-dashboard:latest"
    }
  }

  client_affinity_enabled = false
  https_only              = true
  identity {
    type = "SystemAssigned"
  }

  tags       = var.tags
  depends_on = [azurerm_service_plan.app_service_plan]
}