package s3_client

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *Client) DeleteImage(ctx context.Context, key string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	_, err := c.S3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("failed to delete image with key %q: %w", key, err)
	}

	return nil
}
