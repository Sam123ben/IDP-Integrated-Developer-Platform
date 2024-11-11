output "db_server_fqdn" {
  value = azurerm_postgresql_server.db_server.fqdn
}

output "db_name" {
  value = azurerm_postgresql_database.database.name
}