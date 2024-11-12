# Output for database name
output "db_name" {
  description = "Name of the PostgreSQL database"
  value       = azurerm_postgresql_database.database.name
}

# Output for PostgreSQL server FQDN
output "db_server_fqdn" {
  description = "FQDN of the PostgreSQL server"
  value       = azurerm_postgresql_server.db_server.fqdn
}

# Output for PostgreSQL server name
output "db_server_name" {
  description = "Name of the PostgreSQL server"
  value       = azurerm_postgresql_server.db_server.name
}