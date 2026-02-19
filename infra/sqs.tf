# Dead Letter Queue for failed messages
resource "aws_sqs_queue" "notification_dlq" {
  name                      = "${var.project_name}-${var.environment}-notification-dlq"
  message_retention_seconds = var.sqs_message_retention_seconds

  tags = {
    Name        = "${var.project_name}-${var.environment}-notification-dlq"
    Environment = var.environment
    Project     = var.project_name
  }
}

# Main SQS Queue for notifications
resource "aws_sqs_queue" "notification_queue" {
  name                       = "${var.project_name}-${var.environment}-notification-queue"
  message_retention_seconds  = var.sqs_message_retention_seconds
  visibility_timeout_seconds = var.sqs_visibility_timeout_seconds

  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.notification_dlq.arn
    maxReceiveCount     = var.sqs_max_receive_count
  })

  tags = {
    Name        = "${var.project_name}-${var.environment}-notification-queue"
    Environment = var.environment
    Project     = var.project_name
  }
}

