package s3_client

import (
	"context"

	s3_config "skillspark/internal/config"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	S3     *s3.Client
	Bucket string
}

func NewClient(bucket s3_config.S3) *Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(bucket.Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(bucket.AccessKey,
			bucket.SecretKey, "")))
	// if err != nil {
	// 	return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	// }

	return &Client{
		S3:     s3.NewFromConfig(cfg),
		Bucket: bucket.Bucket,
	}
}
