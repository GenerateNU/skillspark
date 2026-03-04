variable "aws_region" {
  description = "AWS region for resources"
  type        = string
  default     = "us-east-2"
}

variable "project_name" {
  description = "Project name used for resource naming"
  type        = string
  default     = "skillspark"
}

variable "environment" {
  description = "Environment name (e.g., dev, staging, prod)"
  type        = string
  default     = "dev"
}

# Secrets (set via TF_VAR_* or terraform.tfvars, do not commit)
variable "resend_api_key" {
  description = "Resend API key for email notifications"
  type        = string
  sensitive   = true
}

variable "expo_access_token" {
  description = "Expo access token for push notifications (optional)"
  type        = string
  sensitive   = true
  default     = ""
}
