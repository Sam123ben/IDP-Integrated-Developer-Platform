variable "resource_group_name" {
  description = "The name of the resource group"
  type        = string
}

variable "location" {
  description = "The Azure location where the resource group will be created"
  type        = string
}

variable "tags" {}