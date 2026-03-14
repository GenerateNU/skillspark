#!/bin/bash
# Initializes LocalStack resources to match the cloud infrastructure in infra/modules/main/.
# Runs automatically inside the LocalStack container when it is ready.
set -e

REGION="us-east-1"
ACCOUNT_ID="000000000000"
PROJECT="skillspark"

echo "Initializing LocalStack resources..."

# ── SQS ────────────────────────────────────────────────────────────────────────

echo "Creating SQS dead-letter queue..."
awslocal sqs create-queue \
  --queue-name "${PROJECT}-notification-dlq" \
  --region "${REGION}"

DLQ_ARN="arn:aws:sqs:${REGION}:${ACCOUNT_ID}:${PROJECT}-notification-dlq"

echo "Creating SQS notification queue..."
awslocal sqs create-queue \
  --queue-name "${PROJECT}-notification-queue" \
  --region "${REGION}" \
  --attributes "{
    \"VisibilityTimeout\": \"180\",
    \"MessageRetentionPeriod\": \"345600\",
    \"RedrivePolicy\": \"{\\\"deadLetterTargetArn\\\":\\\"${DLQ_ARN}\\\",\\\"maxReceiveCount\\\":\\\"3\\\"}\"
  }"

QUEUE_ARN="arn:aws:sqs:${REGION}:${ACCOUNT_ID}:${PROJECT}-notification-queue"

# ── S3 ─────────────────────────────────────────────────────────────────────────

echo "Creating S3 bucket..."
awslocal s3 mb "s3://${PROJECT}-bucket" --region "${REGION}"

# ── IAM ────────────────────────────────────────────────────────────────────────

echo "Creating IAM role for Lambda..."
awslocal iam create-role \
  --role-name "${PROJECT}-lambda-role" \
  --assume-role-policy-document '{
    "Version": "2012-10-17",
    "Statement": [{
      "Effect": "Allow",
      "Principal": { "Service": "lambda.amazonaws.com" },
      "Action": "sts:AssumeRole"
    }]
  }'

awslocal iam put-role-policy \
  --role-name "${PROJECT}-lambda-role" \
  --policy-name "${PROJECT}-lambda-sqs-policy" \
  --policy-document "{
    \"Version\": \"2012-10-17\",
    \"Statement\": [
      {
        \"Effect\": \"Allow\",
        \"Action\": [\"sqs:ReceiveMessage\",\"sqs:DeleteMessage\",\"sqs:GetQueueAttributes\"],
        \"Resource\": [\"${QUEUE_ARN}\"]
      },
      {
        \"Effect\": \"Allow\",
        \"Action\": [\"logs:CreateLogGroup\",\"logs:CreateLogStream\",\"logs:PutLogEvents\"],
        \"Resource\": \"arn:aws:logs:*:*:*\"
      }
    ]
  }"

LAMBDA_ROLE_ARN="arn:aws:iam::${ACCOUNT_ID}:role/${PROJECT}-lambda-role"

# ── Lambda ─────────────────────────────────────────────────────────────────────

echo "Creating Lambda function..."
awslocal lambda create-function \
  --function-name "${PROJECT}-notification-processor" \
  --runtime provided.al2 \
  --role "${LAMBDA_ROLE_ARN}" \
  --handler bootstrap \
  --zip-file fileb:///tmp/lambda-zip/lambda_function.zip \
  --timeout 30 \
  --memory-size 256 \
  --environment "{\"Variables\":{\"RESEND_API_KEY\":\"${RESEND_API_KEY:-localstack-test}\",\"RESEND_FROM_EMAIL\":\"${RESEND_FROM_EMAIL:-onboarding@resend.dev}\",\"EXPO_ACCESS_TOKEN\":\"${EXPO_ACCESS_TOKEN:-}\"}}" \
  --region "${REGION}"

echo "Waiting for Lambda to be active..."
awslocal lambda wait function-active \
  --function-name "${PROJECT}-notification-processor" \
  --region "${REGION}"

# ── SQS → Lambda trigger ───────────────────────────────────────────────────────

echo "Creating SQS → Lambda event source mapping..."
awslocal lambda create-event-source-mapping \
  --function-name "${PROJECT}-notification-processor" \
  --event-source-arn "${QUEUE_ARN}" \
  --batch-size 10 \
  --maximum-batching-window-in-seconds 5 \
  --function-response-types ReportBatchItemFailures \
  --region "${REGION}"

# ── Summary ────────────────────────────────────────────────────────────────────

echo ""
echo "LocalStack initialization complete."
echo "  S3 bucket : ${PROJECT}-bucket"
echo "  SQS queue : http://sqs.${REGION}.localhost.localstack.cloud:4566/${ACCOUNT_ID}/${PROJECT}-notification-queue"
echo "  SQS DLQ   : http://sqs.${REGION}.localhost.localstack.cloud:4566/${ACCOUNT_ID}/${PROJECT}-notification-dlq"
echo "  Lambda    : ${PROJECT}-notification-processor"
