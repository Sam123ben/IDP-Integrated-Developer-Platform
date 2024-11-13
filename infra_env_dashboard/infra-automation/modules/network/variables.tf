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

variable "public_subnet_name" {
  description = "Name of the public subnet for OpenVPN"
  type        = string
  default     = "public-subnet"
}

variable "public_subnet_cidr" {
  description = "CIDR for the public subnet for OpenVPN"
  type        = string
  default     = "10.0.3.0/24"
}

variable "bastion_name" {
  description = "Name for the Bastion host"
}

variable "bastion_sku" {
  description = "SKU for the Bastion host"
  type        = string
  default     = "Basic"
}

variable "bastion_subnet_name" {
  description = "Name for the Bastion subnet"
}

variable "bastion_subnet_cidr" {
  description = "CIDR for the Bastion subnet"
  type        = string
}