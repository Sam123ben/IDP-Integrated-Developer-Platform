module "network" {
  source              = "./modules/network"
  vnet_name           = "aks-vnet"
  location            = var.location
  resource_group_name = azurerm_resource_group.rg.name
  address_space       = "10.0.0.0/16"
  
  app_subnet_name     = "app-subnet"
  app_subnet_cidr     = "10.0.1.0/24"
  db_subnet_name      = "db-subnet"
  db_subnet_cidr      = "10.0.2.0/24"
  
  app_nsg_name        = "app-nsg"
  db_nsg_name         = "db-nsg"
}

module "database" {
  source              = "./modules/database"
  db_server_name      = "postgres-server"
  location            = var.location
  resource_group_name = azurerm_resource_group.rg.name
  sku_name            = "B_Gen5_1"
  admin_username      = "myadmin"
  admin_password      = var.admin_password
  db_name             = "mydatabase"
  subnet_id           = module.network.db_subnet_id  # Pass subnet ID from the network module
}

module "app" {
  source                  = "./modules/app"
  app_service_plan_name   = "myAppServicePlan"
  location                = var.location
  resource_group_name     = azurerm_resource_group.rg.name
  backend_app_name        = "backend-app"
  frontend_app_name       = "frontend-app"
  database_url            = module.database.db_server_fqdn
}