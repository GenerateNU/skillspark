package s3_client

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *Client) GetImage(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}

	// First check if the image exists
	_, err := c.S3.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return "", fmt.Errorf("failed to get image with key %q: %w", key, err)
	}

	// Generate presigned URL for the existing image
	url, err := c.GeneratePresignedURL(ctx, key, time.Hour)

	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL for image %q: %w", key, err)
	}

	return url, nil
}
