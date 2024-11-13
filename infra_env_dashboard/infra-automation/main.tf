# Define the resource group using the resource_group module
module "resource_group" {
  source              = "./modules/resource_group"
  resource_group_name = var.resource_group_name
  location            = var.location
  tags                = var.tags # Pass tags to the module
}

# Define the network module
module "network" {
  source              = "./modules/network"
  vnet_name           = "dashboard-vnet"
  location            = module.resource_group.location
  resource_group_name = module.resource_group.name
  address_space       = "10.0.0.0/16"

  app_subnet_name = "app-subnet"
  app_subnet_cidr = "10.0.1.0/24"
  db_subnet_name  = "db-subnet"
  db_subnet_cidr  = "10.0.2.0/24"
  public_subnet_name  = "public-subnet"
  public_subnet_cidr  = "10.0.3.0/24"

  app_nsg_name = "app-nsg"
  db_nsg_name  = "db-nsg"

  tags = var.tags # Pass tags to the module

  depends_on = [module.resource_group]
}

# Define the database module
module "database" {
  source              = "./modules/database"
  db_server_name      = "dashboard-server"
  location            = module.resource_group.location
  resource_group_name = module.resource_group.name
  sku_name            = "B_Gen5_1"
  admin_username      = "psqladmin"
  admin_password      = var.admin_password
  vm_admin_password   = var.vm_admin_password
  db_name             = "mydatabase"
  subnet_id           = module.network.db_subnet_id # Pass subnet ID from the network module output
  app_subnet_id       = module.network.app_subnet_id # Pass app subnet ID from the network module output
  vnet_id             = module.network.vnet_id      # Pass VNet ID from the network module output

  tags = var.tags # Pass tags to the module

  depends_on = [module.network]
}

# Define the app module
module "app" {
  source                = "./modules/app"
  app_service_plan_name = "DashboardAppPlan"
  location              = module.resource_group.location
  resource_group_name   = module.resource_group.name
  backend_app_name      = "dashboard-backend-app"
  frontend_app_name     = "dashboard-frontend-app"
  database_url          = module.database.db_server_fqdn # Pass FQDN from database module output
  db_server_name        = module.database.db_server_name

  tags = var.tags # Pass tags to the module

  depends_on = [module.database]
}

module "openvpn" {
  source              = "./modules/openvpn"
  location            = var.location
  resource_group_name = var.resource_group_name
  public_subnet_id    = module.network.public_subnet_id
  vm_admin_username   = "openvpn"
  vm_admin_password   = var.vm_admin_password
  tags                = var.tags
}
