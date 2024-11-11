output "db_name" {
  value = azurerm_postgresql_database.database.name
}

output "db_server_fqdn" {
  description = "FQDN of the PostgreSQL server"
  value       = azurerm_postgresql_server.db_server.fqdn  # Adjust to match your PostgreSQL server resource name
}

output "db_server_name" {
  value = azurerm_postgresql_server.db_server.name
}