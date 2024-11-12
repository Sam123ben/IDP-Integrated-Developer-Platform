output "backend_app_url" {
  value = azurerm_linux_web_app.backend_app.default_site_hostname
}

output "frontend_app_url" {
  value = azurerm_linux_web_app.frontend_app.default_site_hostname
}