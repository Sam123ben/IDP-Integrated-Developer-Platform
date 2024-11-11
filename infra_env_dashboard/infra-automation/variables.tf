variable "location" {
  description = "Azure region where resources will be created"
  default     = "Australia East"
}

variable "resource_group_name" {
  description = "The name of the resource group to be created"
  type        = string
}

variable "admin_password" {
  description = "Admin password for PostgreSQL server"
  sensitive   = true
}

variable "client_id" {
  type        = string
  description = "Azure Client ID for authentication"
}

variable "client_secret" {
  type        = string
  description = "Azure Client Secret for authentication"
  sensitive   = true
}

variable "subscription_id" {
  type        = string
  description = "Azure Subscription ID"
}

variable "tenant_id" {
  type        = string
  description = "Azure Tenant ID"
}