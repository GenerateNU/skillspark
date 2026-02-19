# Build Lambda binary
resource "null_resource" "lambda_build" {
  triggers = {
    source_hash = sha256(join("", [
      for f in fileset("${var.lambda_source_path}", "**/*.go") : filesha256("${var.lambda_source_path}/${f}")
    ]))
  }

  provisioner "local-exec" {
    command = <<-EOT
      mkdir -p ${path.module}/lambda && \
      cd ${var.lambda_source_path} && \
      GOOS=linux GOARCH=amd64 go build -o ${path.module}/lambda/bootstrap .
    EOT
  }
}

# Archive Lambda function binary
data "archive_file" "lambda_zip" {
  depends_on  = [null_resource.lambda_build]
  type        = "zip"
  source_file = "${path.module}/lambda/bootstrap"
  output_path = "${path.module}/lambda_function.zip"
}

# IAM role for Lambda function
resource "aws_iam_role" "lambda_role" {
  name = "${var.project_name}-${var.environment}-notification-lambda-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name        = "${var.project_name}-${var.environment}-notification-lambda-role"
    Environment = var.environment
    Project     = var.project_name
  }
}

# IAM policy for Lambda to read from SQS
resource "aws_iam_role_policy" "lambda_sqs_policy" {
  name = "${var.project_name}-${var.environment}-lambda-sqs-policy"
  role = aws_iam_role.lambda_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "sqs:ReceiveMessage",
          "sqs:DeleteMessage",
          "sqs:GetQueueAttributes"
        ]
        Resource = [
          aws_sqs_queue.notification_queue.arn
        ]
      },
      {
        Effect = "Allow"
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Resource = "arn:aws:logs:*:*:*"
      }
    ]
  })
}

# Lambda function
resource "aws_lambda_function" "notification_processor" {
  filename                       = data.archive_file.lambda_zip.output_path
  function_name                  = "${var.project_name}-${var.environment}-notification-processor"
  role                           = aws_iam_role.lambda_role.arn
  handler                        = "bootstrap"
  runtime                        = var.lambda_runtime
  timeout                        = var.lambda_timeout
  memory_size                    = var.lambda_memory_size
  reserved_concurrent_executions = var.lambda_reserved_concurrency

  source_code_hash = data.archive_file.lambda_zip.output_base64sha256

  depends_on = [
    null_resource.lambda_build,
    data.archive_file.lambda_zip,
    aws_iam_role_policy.lambda_sqs_policy
  ]

  environment {
    variables = {
      RESEND_API_KEY    = var.resend_api_key
      EXPO_ACCESS_TOKEN = var.expo_access_token
    }
  }

  tags = {
    Name        = "${var.project_name}-${var.environment}-notification-processor"
    Environment = var.environment
    Project     = var.project_name
  }
}

# CloudWatch log group for Lambda function
# resource "aws_cloudwatch_log_group" "lambda_logs" {
#   name              = "/aws/lambda/${var.project_name}-${var.environment}-notification-processor"
#   retention_in_days = 14

#   tags = {
#     Name        = "${var.project_name}-${var.environment}-notification-processor-logs"
#     Environment = var.environment
#     Project     = var.project_name
#   }
# }

# Event source mapping: SQS to Lambda
resource "aws_lambda_event_source_mapping" "sqs_trigger" {
  event_source_arn                   = aws_sqs_queue.notification_queue.arn
  function_name                      = aws_lambda_function.notification_processor.arn
  batch_size                         = var.sqs_batch_size
  maximum_batching_window_in_seconds = var.sqs_batch_window

  # Enable partial batch response
  function_response_types = ["ReportBatchItemFailures"]
}

