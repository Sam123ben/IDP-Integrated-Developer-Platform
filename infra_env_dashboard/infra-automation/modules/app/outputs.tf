output "backend_app_url" {
  value = azurerm_app_service.backend_app.default_site_hostname
}

output "frontend_app_url" {
  value = azurerm_app_service.frontend_app.default_site_hostname
}