output "db_name" {
  description = "Name of the PostgreSQL database"
  value       = azurerm_postgresql_flexible_server_database.database.name
}

output "db_server_fqdn" {
  description = "FQDN of the PostgreSQL flexible server"
  value       = azurerm_postgresql_flexible_server.db_server.fqdn
}

output "db_server_name" {
  description = "Name of the PostgreSQL flexible server"
  value       = azurerm_postgresql_flexible_server.db_server.name
}