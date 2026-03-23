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
	var region, accessKey, secretKey, queueURL string

	if sqsConfig.UseLocalStack {
		region = sqsConfig.LocalStackRegion
		accessKey = sqsConfig.LocalStackAccessKey
		secretKey = sqsConfig.LocalStackSecretKey
		queueURL = sqsConfig.LocalStackQueueURL
	} else {
		region = sqsConfig.Region
		accessKey = sqsConfig.AccessKey
		secretKey = sqsConfig.SecretKey
		queueURL = sqsConfig.QueueURL
	}

	opts := []func(*config.LoadOptions) error{
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	}
	if sqsConfig.UseLocalStack {
		opts = append(opts, config.WithBaseEndpoint(sqsConfig.LocalStackEndpoint))
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	return &Client{
		SQS:      sqs.NewFromConfig(cfg),
		QueueURL: queueURL,
	}, nil
}
