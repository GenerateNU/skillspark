# Terraform Infrastructure for SkillSpark Notifications

This directory contains Terraform configuration to provision AWS infrastructure for the SkillSpark notification system, including:

- SQS Queue for notification messages
- SQS Dead-Letter Queue for failed messages
- Lambda function to process notifications
- IAM roles and policies
- Event source mapping connecting SQS to Lambda

## Prerequisites

1. **AWS CLI** configured with appropriate credentials
2. **Terraform** >= 1.0 installed
3. **Go** >= 1.21 installed (for building Lambda function)
4. **Lambda function code** in the `lambda/` directory (within this infra directory)

## Configuration

### Variables

Key variables can be set via:
- `terraform.tfvars` file (recommended for sensitive values)
- Command-line flags: `-var="key=value"`
- Environment variables: `TF_VAR_key=value`

Important variables:
- `aws_region`: AWS region (default: `us-east-1`)
- `project_name`: Project name for resource naming (default: `skillspark`)
- `environment`: Environment name (default: `dev`)
- `resend_api_key`: Resend API key (required, sensitive)
- `expo_access_token`: Expo access token (optional, sensitive)
- `lambda_source_path`: Path to Lambda function code (default: `./lambda`)

### Concurrency Configuration

The infrastructure is configured to allow only one Lambda execution at a time:
- `lambda_reserved_concurrency`: 1 (ensures only one execution at a time)
- `sqs_batch_size`: 10 (messages per batch)
- `sqs_batch_window`: 5 (seconds to gather records)

This configuration ensures sequential processing of notifications, preventing concurrent executions.

## Usage

### 1. Build Lambda Function

Before deploying, build the Lambda function binary:

```bash
cd infra
./build-lambda.sh
```

Or manually:

```bash
cd infra/lambda
GOOS=linux GOARCH=amd64 go build -o ../lambda/bootstrap .
```

**Note**: Terraform will automatically build the Lambda function during `terraform apply` using a `null_resource` provisioner. However, you can pre-build it for faster deployments or to test the build process.

### 2. Initialize Terraform

```bash
terraform init
```

### 3. Create `terraform.tfvars`

Create a `terraform.tfvars` file with your configuration:

```hcl
aws_region = "us-east-1"
environment = "dev"
project_name = "skillspark"

resend_api_key = "your-resend-api-key"
expo_access_token = "your-expo-token"  # Optional

lambda_source_path = "./lambda"
```

### 4. Plan the deployment

```bash
terraform plan
```

### 5. Apply the configuration

```bash
terraform apply
```

Terraform will automatically:
- Build the Lambda function binary (if not pre-built)
- Create a zip file with the binary
- Deploy the Lambda function

### 6. Get outputs

After deployment, get the SQS queue URL:

```bash
terraform output sqs_queue_url
```

Use this URL in your backend application's `AWS_SQS_QUEUE_URL` environment variable.

## Lambda Function

The Lambda function is written in Go and located in the `lambda/` directory within this infra directory.

### Structure

- `main.go`: Lambda entry point
- `handler.go`: SQS event handler with partial batch failure support
- `notification.go`: Notification processing logic
- `resend.go`: Resend email API client
- `expo.go`: Expo push notification API client
- `models.go`: Notification message models

### Features

- Processes SQS messages with the `NotificationMessage` structure
- Sends emails via Resend API
- Sends push notifications via Expo API
- Supports partial batch failures (failed messages are retried, successful ones are deleted)
- Handles errors appropriately with logging
- Uses Go runtime (`provided.al2`)

### Building

The Lambda function must be built for Linux/AMD64 architecture:

```bash
cd infra/lambda
GOOS=linux GOARCH=amd64 go build -o ../lambda/bootstrap .
```

The binary must be named `bootstrap` for the `provided.al2` runtime.

### Environment Variables

The Lambda function requires:
- `RESEND_API_KEY`: Resend API key for email notifications (required)
- `EXPO_ACCESS_TOKEN`: Expo access token for push notifications (optional)
- `RESEND_FROM_EMAIL`: From email address for Resend (optional, defaults to `notifications@skillspark.app`)

See `../backend/internal/notification/LAMBDA_SETUP.md` for detailed Lambda function requirements.

## Resource Naming

Resources are named using the pattern: `{project_name}-{environment}-{resource_type}`

Example: `skillspark-dev-notification-queue`

## Outputs

After applying, Terraform will output:
- `sqs_queue_url`: Use this in your backend application
- `sqs_queue_arn`: Queue ARN for IAM policies
- `sqs_dlq_url`: Dead-letter queue URL for monitoring
- `lambda_function_name`: Lambda function name
- `lambda_function_arn`: Lambda function ARN

## Monitoring

Consider setting up CloudWatch alarms for:
- Lambda errors
- DLQ message count
- Lambda duration
- SQS queue depth

## Cleanup

To destroy all resources:

```bash
terraform destroy
```

**Warning**: This will delete all resources created by this Terraform configuration.

