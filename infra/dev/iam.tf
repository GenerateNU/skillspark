locals {
  # S3 bucket names — these live in the infra/buckets Terraform state
  s3_bucket_name = "${var.project_name}-${var.environment}"
}

# IAM user for the backend application
resource "aws_iam_user" "backend" {
  name = "${var.project_name}-${var.environment}-backend"

  tags = {
    Name        = "${var.project_name}-${var.environment}-backend"
    Environment = var.environment
    Project     = var.project_name
  }
}

# Access key for the backend user (secret stored in Terraform state)
resource "aws_iam_access_key" "backend" {
  user = aws_iam_user.backend.name
}

# Policy granting access to the notification queue and S3 bucket
resource "aws_iam_user_policy" "backend" {
  name = "${var.project_name}-${var.environment}-backend-policy"
  user = aws_iam_user.backend.name

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "SQSSendMessage"
        Effect = "Allow"
        Action = ["sqs:SendMessage"]
        Resource = [
          module.notifications.sqs_queue_arn,
        ]
      },
      {
        Sid    = "S3BucketAccess"
        Effect = "Allow"
        Action = [
          "s3:PutObject",
          "s3:GetObject",
          "s3:DeleteObject",
        ]
        Resource = "arn:aws:s3:::${local.s3_bucket_name}/*"
      }
    ]
  })
}
