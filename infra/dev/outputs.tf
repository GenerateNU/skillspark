output "sqs_queue_url" {
  description = "URL of the SQS notification queue (use in backend AWS_SQS_QUEUE_URL)"
  value       = module.notifications.sqs_queue_url
}

output "sqs_queue_arn" {
  description = "ARN of the SQS notification queue"
  value       = module.notifications.sqs_queue_arn
}

output "sqs_dlq_url" {
  description = "URL of the SQS dead-letter queue"
  value       = module.notifications.sqs_dlq_url
}

output "sqs_dlq_arn" {
  description = "ARN of the SQS dead-letter queue"
  value       = module.notifications.sqs_dlq_arn
}

output "lambda_function_name" {
  description = "Name of the Lambda function"
  value       = module.notifications.lambda_function_name
}

output "lambda_function_arn" {
  description = "ARN of the Lambda function"
  value       = module.notifications.lambda_function_arn
}

output "lambda_role_arn" {
  description = "ARN of the Lambda execution role"
  value       = module.notifications.lambda_role_arn
}
