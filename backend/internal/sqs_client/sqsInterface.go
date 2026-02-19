package sqs_client

import (
	"context"
)

type SQSInterface interface {
	SendMessage(ctx context.Context, messageBody interface{}) error
}

