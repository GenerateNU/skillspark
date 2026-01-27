package s3_client

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadImage uploads file content to S3 with the given key.
// The caller is responsible for closing the reader after this function returns.
func (c *Client) UploadImage(ctx context.Context, key string, image_data []byte) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}

	data := bytes.NewReader(image_data)

	_, err := c.S3.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(key),
		Body:   data,
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload image with key %q: %w", key, err)
	}

	return key, nil
}
