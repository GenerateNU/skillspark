package s3_client

import (
	"context"
	"fmt"

	s3_config "skillspark/internal/config"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	S3     *s3.Client
	Bucket string
}

func NewClient(bucket s3_config.S3) (*Client, error) {
	opts := []func(*config.LoadOptions) error{
		config.WithRegion(bucket.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(bucket.AccessKey, bucket.SecretKey, "")),
	}
	if bucket.UseLocalStack {
		opts = append(opts, config.WithBaseEndpoint(bucket.LocalStackEndpoint))
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		if bucket.UseLocalStack {
			o.UsePathStyle = true
		}
	})

	return &Client{
		S3:     s3Client,
		Bucket: bucket.Bucket,
	}, nil
}
