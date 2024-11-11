variable "location" {
  description = "Azure region where resources will be created"
  default     = "Australia East"
}

variable "admin_password" {
  description = "Admin password for PostgreSQL server"
  sensitive   = true
}

variable "storage_endpoint" {
  description = "Storage endpoint for threat detection policy"
}