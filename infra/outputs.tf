output "sqs_queue_url" {
  description = "URL of the SQS notification queue"
  value       = aws_sqs_queue.notification_queue.url
}

output "sqs_queue_arn" {
  description = "ARN of the SQS notification queue"
  value       = aws_sqs_queue.notification_queue.arn
}

output "sqs_dlq_url" {
  description = "URL of the SQS dead-letter queue"
  value       = aws_sqs_queue.notification_dlq.url
}

output "sqs_dlq_arn" {
  description = "ARN of the SQS dead-letter queue"
  value       = aws_sqs_queue.notification_dlq.arn
}

output "lambda_function_name" {
  description = "Name of the Lambda function"
  value       = aws_lambda_function.notification_processor.function_name
}

output "lambda_function_arn" {
  description = "ARN of the Lambda function"
  value       = aws_lambda_function.notification_processor.arn
}

output "lambda_role_arn" {
  description = "ARN of the Lambda execution role"
  value       = aws_iam_role.lambda_role.arn
}

