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

func NewClient(s3Config s3_config.S3) (*Client, error) {
	var region, accessKey, secretKey, bucket string

	if s3Config.UseLocalStack {
		region = s3Config.LocalStackRegion
		accessKey = s3Config.LocalStackAccessKey
		secretKey = s3Config.LocalStackSecretKey
		bucket = s3Config.LocalStackBucket
	} else {
		region = s3Config.Region
		accessKey = s3Config.AccessKey
		secretKey = s3Config.SecretKey
		bucket = s3Config.Bucket
	}

	opts := []func(*config.LoadOptions) error{
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	}
	if s3Config.UseLocalStack {
		opts = append(opts, config.WithBaseEndpoint(s3Config.LocalStackEndpoint))
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		if s3Config.UseLocalStack {
			o.UsePathStyle = true
		}
	})

	return &Client{
		S3:     s3Client,
		Bucket: bucket,
	}, nil
}
