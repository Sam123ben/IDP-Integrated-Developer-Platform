resource "azurerm_postgresql_server" "db_server" {
  name                = var.db_server_name
  location            = var.location
  resource_group_name = var.resource_group_name

  sku_name   = var.sku_name
  version    = "11"
  storage_mb = 5120
  administrator_login          = var.admin_username
  administrator_login_password = var.admin_password
  auto_grow_enabled            = true

  threat_detection_policy {
    state                      = "Enabled"
    email_account_admins       = true
    storage_endpoint           = var.storage_endpoint
    retention_days             = 7
  }
}

resource "azurerm_postgresql_database" "database" {
  name                = var.db_name
  resource_group_name = var.resource_group_name
  server_name         = azurerm_postgresql_server.db_server.name
  charset             = "UTF8"
  collation           = "English_United States.1252"
}