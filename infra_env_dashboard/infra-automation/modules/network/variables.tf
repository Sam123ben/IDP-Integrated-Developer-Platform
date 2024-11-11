variable "vnet_name" {}
variable "location" {}
variable "resource_group_name" {}
variable "address_space" {}

variable "app_subnet_name" {}
variable "app_subnet_cidr" {}

variable "db_subnet_name" {}
variable "db_subnet_cidr" {}
variable "tags" {}

variable "app_nsg_name" {
  description = "Name for the Application Network Security Group"
}

variable "db_nsg_name" {
  description = "Name for the Database Network Security Group"
}