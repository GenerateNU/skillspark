variable "aws_region" {
  description = "AWS region for resources"
  type        = string
  default     = "us-east-1"
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

variable "lambda_runtime" {
  description = "Lambda function runtime"
  type        = string
  default     = "provided.al2"
}

variable "lambda_timeout" {
  description = "Lambda function timeout in seconds"
  type        = number
  default     = 30
}

variable "lambda_memory_size" {
  description = "Lambda function memory size in MB"
  type        = number
  default     = 256
}

variable "lambda_reserved_concurrency" {
  description = "Lambda reserved concurrency (for rate limiting: 2 messages/sec)"
  type        = number
  default     = 2
}

variable "sqs_batch_size" {
  description = "Maximum number of records to retrieve in a single batch"
  type        = number
  default     = 10
}

variable "sqs_batch_window" {
  description = "Maximum amount of time to gather records before invoking the function (seconds)"
  type        = number
  default     = 5
}

variable "sqs_max_receive_count" {
  description = "Maximum number of times a message can be received before moving to DLQ"
  type        = number
  default     = 3
}

variable "sqs_message_retention_seconds" {
  description = "Number of seconds to retain messages in the queue"
  type        = number
  default     = 345600 # 4 days
}

variable "sqs_visibility_timeout_seconds" {
  description = "Visibility timeout for messages in seconds (should be >= lambda_timeout * 6)"
  type        = number
  default     = 180 # 3 minutes (6x lambda timeout)
}

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

variable "lambda_source_path" {
  description = "Path to the Lambda function source code"
  type        = string
  default     = "./lambda"
}

