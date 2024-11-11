# environments/dev/main.tf
module "network" {
  source = "../../modules/network"
  # Set values specific to the dev environment
}

module "database" {
  source = "../../modules/database"
  # Set values specific to the dev environment
}

module "app" {
  source = "../../modules/app"
  # Set values specific to the dev environment
}