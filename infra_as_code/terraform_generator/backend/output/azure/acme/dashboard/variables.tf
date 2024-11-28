
variable "aws_oidc_token_file" {
  description = "AWS OIDC Token File"
  type        = string
  
}

variable "aws_region" {
  description = "AWS Region"
  type        = string
  
}

variable "aws_role_arn" {
  description = "AWS Role ARN"
  type        = string
  
}

variable "azure_client_id" {
  description = "Azure Client ID"
  type        = string
  
}

variable "azure_client_secret" {
  description = "Azure Client Secret"
  type        = string
  
  sensitive   = true
  
}

variable "azure_subscription_id" {
  description = "Azure Subscription ID"
  type        = string
  
}

variable "azure_tenant_id" {
  description = "Azure Tenant ID"
  type        = string
  
}

variable "environment" {
  description = "The deployment environment"
  type        = string
  
}

variable "gcp_credentials_file" {
  description = "GCP Credentials File"
  type        = string
  
  sensitive   = true
  
}

variable "gcp_project_id" {
  description = "GCP Project ID"
  type        = string
  
}

variable "gcp_region" {
  description = "GCP Region"
  type        = string
  
}

variable "region" {
  description = "The deployment region"
  type        = string
  
}

