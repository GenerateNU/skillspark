# AWS Lambda Function Setup for Notification Processing

## Overview

The Lambda function is responsible for processing notification messages from AWS SQS and sending them via Resend (email) and Expo Push Notifications (push).

## Lambda Function Requirements

### 1. Trigger Configuration

- **Event Source**: AWS SQS Queue
- **Queue**: The SQS queue URL configured in the Go application
- **Batch Size**: Configure to achieve 2 messages per second rate limit
- **Batch Window**: Adjust based on batch size to maintain 2/sec rate
- **Reserved Concurrency**: Set to limit concurrent executions (e.g., 2 for 2/sec rate)

### 2. Message Structure

The Lambda function will receive SQS messages with the following JSON structure (matching `NotificationMessage` from Go models):

```json
{
  "notification_type": "email" | "push" | "both",
  "recipient_email": "user@example.com" (optional),
  "recipient_push_token": "ExponentPushToken[...]" (optional),
  "subject": "Email subject" (optional, for email notifications),
  "body": "Notification body text",
  "metadata": {} (optional JSON object for additional data)
}
```

### 3. Processing Logic

The Lambda function should:

1. Parse the SQS message body (JSON)
2. Validate the notification structure
3. Based on `notification_type`:
   - **"email"**: Send email via Resend API
   - **"push"**: Send push notification via Expo Push Notification API
   - **"both"**: Send both email and push notification
4. Handle errors appropriately:
   - Log errors for debugging
   - Use SQS dead-letter queue for failed messages
   - Return appropriate status codes

### 4. Environment Variables

The Lambda function should have the following environment variables:

- `RESEND_API_KEY`: API key for Resend email service
- `EXPO_ACCESS_TOKEN`: (Optional) Expo access token if using authenticated API

### 5. Rate Limiting

To achieve 2 messages per second:

- Configure SQS batch size appropriately
- Set Lambda reserved concurrency to limit concurrent executions
- Use Lambda batch window to control processing rate
- Example: Batch size of 10 with 5-second window = 2/sec

### 6. Error Handling

- Failed messages should be sent to a dead-letter queue (DLQ)
- Log all errors with appropriate context
- Retry logic should be handled by SQS (configure max receive count)

### 7. Dependencies

The Lambda function will need:

- HTTP client for Resend API calls
- HTTP client for Expo Push Notification API calls
- JSON parsing library
- Error handling and logging utilities

## Resend API Integration

- **Endpoint**: `https://api.resend.com/emails`
- **Method**: POST
- **Headers**: `Authorization: Bearer {RESEND_API_KEY}`, `Content-Type: application/json`
- **Rate Limit**: 2 requests per second (enforced by Lambda concurrency)

## Expo Push Notification API Integration

- **Endpoint**: `https://exp.host/--/api/v2/push/send`
- **Method**: POST
- **Headers**: `Content-Type: application/json`, optionally `Authorization: Bearer {EXPO_ACCESS_TOKEN}`
- **Rate Limit**: Handled by Expo (typically much higher than 2/sec)

## Example Lambda Function Structure

```python
# Pseudo-code structure
def lambda_handler(event, context):
    for record in event['Records']:
        try:
            message_body = json.loads(record['body'])
            notification_type = message_body['notification_type']
            
            if notification_type in ['email', 'both']:
                send_email_via_resend(message_body)
            
            if notification_type in ['push', 'both']:
                send_push_via_expo(message_body)
                
        except Exception as e:
            # Log error and let SQS handle retry/DLQ
            logger.error(f"Failed to process notification: {e}")
            raise
```

## SQS Dead-Letter Queue Setup

1. Create a DLQ for failed notifications
2. Configure the main SQS queue to send failed messages to DLQ after max receive count (e.g., 3)
3. Monitor DLQ for patterns indicating issues

## Monitoring

- Set up CloudWatch alarms for:
  - Lambda errors
  - DLQ message count
  - Lambda duration
  - SQS queue depth

