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

# Private Endpoint for PostgreSQL in the database subnet
resource "azurerm_private_endpoint" "postgres_private_endpoint" {
  name                = "${var.db_server_name}-private-endpoint"
  location            = var.location
  resource_group_name = var.resource_group_name
  subnet_id           = var.db_subnet_id  # Subnet ID where private endpoint will be created

  private_service_connection {
    name                           = "${var.db_server_name}-connection"
    private_connection_resource_id = azurerm_postgresql_server.db_server.id
    is_manual_connection           = false
    subresource_names              = ["postgresqlServer"]
  }
}

# Private DNS Zone for the PostgreSQL server (replace with actual private DNS zone name)
resource "azurerm_private_dns_zone" "db_private_dns_zone" {
  name                = "privatelink.postgres.database.azure.com"
  resource_group_name = var.resource_group_name
}

# DNS Zone Association with VNet
resource "azurerm_private_dns_zone_virtual_network_link" "db_dns_zone_link" {
  name                  = "${var.db_server_name}-dns-link"
  resource_group_name   = var.resource_group_name
  private_dns_zone_name = azurerm_private_dns_zone.db_private_dns_zone.name
  virtual_network_id    = var.vnet_id
}

# DNS A Record for the PostgreSQL private endpoint
resource "azurerm_private_dns_a_record" "db_private_a_record" {
  name                = "${azurerm_postgresql_server.db_server.name}.postgres.database.azure.com"
  zone_name           = azurerm_private_dns_zone.db_private_dns_zone.name
  resource_group_name = var.resource_group_name
  ttl                 = 300
  records             = [azurerm_private_endpoint.postgres_private_endpoint.private_ip_address]
}