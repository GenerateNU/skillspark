package s3_client

import (
	"context"
	"time"
)

type S3Interface interface {
	GeneratePresignedURL(ctx context.Context, key string, expiry time.Duration) (string, error)
	UploadImage(ctx context.Context, key *string, image_data []byte) (*string, error)
}
