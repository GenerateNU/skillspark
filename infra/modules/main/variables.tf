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

variable "opensearch_domain_name" {
  description = "name for opensearch instance"
  type        = string
  default     = "skillsparkdomain"

}

variable "opensearch_engine_version" {
  description = "engine version for opensearch"
  type = string
  default = "OpenSearch_2.17"

}
variable "opensearch_instance" {
  description = "instance to use for opensearch"
  type        = string
  default     = "r8g.large.search"
}

variable "master_user" {
  description = "master username for opensearch"
  type        = string
  default     = ""
}

variable "master_pass" {
  description = "master password for opensearch"
  type        = string
  sensitive   = true 
  default     = "" 
}

variable "event_shards" {
  description = "shards to use for events index"
  type        = integer
  default     = 2 
}

variable "event_replicas" {
  description = "replicas to use for events index"
  type        = integer
  default     = 3
}


