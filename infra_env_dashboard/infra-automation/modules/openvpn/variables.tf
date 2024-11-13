variable "resource_group_name" {
  description = "Name of the resource group."
  type        = string
}

variable "location" {
  description = "Azure location for the VM."
  type        = string
}

variable "vm_admin_username" {
  description = "Admin username for the VM."
  type        = string
}

variable "vm_admin_password" {
  description = "Admin password for the VM."
  type        = string
}

variable "subnet_id" {
  description = "Subnet ID where the VM will be deployed."
  type        = string
}

variable "tags" {
  description = "Tags to apply to resources."
  type        = map(string)
  default     = {}
}

variable "public_subnet_id" {
  description = "Subnet ID for the public subnet."
  type        = string
}