package sqs_client

import (
	"context"
	"fmt"

	sqs_config "skillspark/internal/config"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Client struct {
	SQS      *sqs.Client
	QueueURL string
}

func NewClient(sqsConfig sqs_config.SQS) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(sqsConfig.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(sqsConfig.AccessKey,
			sqsConfig.SecretKey, "")))

	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	return &Client{
		SQS:      sqs.NewFromConfig(cfg),
		QueueURL: sqsConfig.QueueURL,
	}, nil
}

