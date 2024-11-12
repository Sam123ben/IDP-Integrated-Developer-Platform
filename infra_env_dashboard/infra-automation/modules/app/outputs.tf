output "backend_app_url" {
  value = azurerm_linux_web_app.backend_app.default_hostname
  description = "URL of the backend app service"
}

output "frontend_app_url" {
  value = azurerm_linux_web_app.frontend_app.default_hostname
  description = "URL of the frontend app service"
}