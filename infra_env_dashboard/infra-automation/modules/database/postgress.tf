# PostgreSQL Flexible Server with Private Access Mode
resource "azurerm_postgresql_flexible_server" "db_server" {
  name                = var.db_server_name
  location            = var.location
  resource_group_name = var.resource_group_name
  delegated_subnet_id = var.subnet_id # Ensure this is a delegated subnet for PostgreSQL

  # Configure the PostgreSQL version and administrator credentials
  version                = "13"
  administrator_login    = var.admin_username
  administrator_password = var.admin_password

  # Set up storage
  storage_mb = 32768
  sku_name   = "B_Standard_B1ms"
  zone       = "1"

  # Enable Private Access Mode (if available in your region)
  public_network_access_enabled = false
  private_dns_zone_id           = azurerm_private_dns_zone.db_private_dns_zone.id

  tags = var.tags
  depends_on = [azurerm_private_dns_zone_virtual_network_link.db_dns_zone_link]
}

# PostgreSQL Database in the Flexible Server
resource "azurerm_postgresql_flexible_server_database" "database" {
  name      = var.db_name
  server_id = azurerm_postgresql_flexible_server.db_server.id
  collation = "en_US.utf8"
  charset   = "UTF8"

  # Prevent accidental data loss by disabling deletion
  lifecycle {
    prevent_destroy = true
  }
  depends_on = [ azurerm_postgresql_flexible_server.db_server ]
}

# Private DNS Zone for the PostgreSQL server
resource "azurerm_private_dns_zone" "db_private_dns_zone" {
  name                = "privatelink.postgres.database.azure.com"
  resource_group_name = var.resource_group_name
  tags                = var.tags # Apply tags here
}

# DNS Zone Association with VNet
resource "azurerm_private_dns_zone_virtual_network_link" "db_dns_zone_link" {
  name                  = "${var.db_server_name}-dns-link"
  resource_group_name   = var.resource_group_name
  private_dns_zone_name = azurerm_private_dns_zone.db_private_dns_zone.name
  virtual_network_id    = var.vnet_id
  tags                  = var.tags # Apply tags here
}