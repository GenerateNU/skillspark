package sqs_client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// SendMessage sends a notification message to the SQS queue
func (c *Client) SendMessage(ctx context.Context, messageBody interface{}) error {
	// Serialize message body to JSON
	bodyBytes, err := json.Marshal(messageBody)
	if err != nil {
		return fmt.Errorf("failed to marshal message body: %w", err)
	}

	body := string(bodyBytes)

	_, err = c.SQS.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(c.QueueURL),
		MessageBody: aws.String(body),
	})

	if err != nil {
		return fmt.Errorf("failed to send message to SQS: %w", err)
	}

	return nil
}

