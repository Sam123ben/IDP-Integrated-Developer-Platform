# PostgreSQL Flexible Server
resource "azurerm_postgresql_flexible_server" "db_server" {
  name                = var.db_server_name
  location            = var.location
  resource_group_name = var.resource_group_name

  version                       = "12"
  administrator_login           = var.admin_username
  administrator_password        = var.admin_password
  delegated_subnet_id           = var.subnet_id  # Subnet ID for Flexible Server
  private_dns_zone_id           = azurerm_private_dns_zone.db_private_dns_zone.id
  public_network_access_enabled = false
  storage_mb                    = 32768  # Example size; adjust as needed
  storage_tier                  = "P30"
  sku_name                      = "GP_Standard_D4s_v3"  # Example SKU; adjust as needed

  zone = "1"  # Define the zone based on region availability
  tags = var.tags  # Apply tags here
}

# PostgreSQL Database in the Flexible Server
resource "azurerm_postgresql_flexible_database" "database" {
  name                = var.db_name
  resource_group_name = var.resource_group_name
  server_name         = azurerm_postgresql_flexible_server.db_server.name
  charset             = "UTF8"
  collation           = "English_United States.1252"
  tags = var.tags  # Apply tags here
}

# Private DNS Zone for the PostgreSQL server
resource "azurerm_private_dns_zone" "db_private_dns_zone" {
  name                = "privatelink.postgres.database.azure.com"
  resource_group_name = var.resource_group_name
  tags = var.tags  # Apply tags here
}

# DNS Zone Association with VNet
resource "azurerm_private_dns_zone_virtual_network_link" "db_dns_zone_link" {
  name                  = "${var.db_server_name}-dns-link"
  resource_group_name   = var.resource_group_name
  private_dns_zone_name = azurerm_private_dns_zone.db_private_dns_zone.name
  virtual_network_id    = var.vnet_id
  tags = var.tags  # Apply tags here
}

# Private Endpoint for PostgreSQL Flexible Server in the database subnet
resource "azurerm_private_endpoint" "postgres_private_endpoint" {
  name                = "${var.db_server_name}-private-endpoint"
  location            = var.location
  resource_group_name = var.resource_group_name
  subnet_id           = var.subnet_id

  private_service_connection {
    name                           = "${var.db_server_name}-connection"
    private_connection_resource_id = azurerm_postgresql_flexible_server.db_server.id
    is_manual_connection           = false
    subresource_names              = ["postgresqlServer"]
  }
  tags = var.tags  # Apply tags here
}

# DNS A Record for the PostgreSQL private endpoint
resource "azurerm_private_dns_a_record" "db_private_a_record" {
  name                = "${azurerm_postgresql_flexible_server.db_server.name}.postgres.database.azure.com"
  zone_name           = azurerm_private_dns_zone.db_private_dns_zone.name
  resource_group_name = var.resource_group_name
  ttl                 = 300
  records             = [azurerm_private_endpoint.postgres_private_endpoint.private_service_connection[0].private_ip_address]
  tags = var.tags  # Apply tags here
}