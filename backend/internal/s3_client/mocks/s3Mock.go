package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

type S3ClientMock struct {
	mock.Mock
}

func (m *S3ClientMock) UploadImage(ctx context.Context, key *string, image_data []byte) (*string, error) {
	args := m.Called(ctx, key, image_data)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*string), nil
}

func (m *S3ClientMock) GeneratePresignedURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	args := m.Called(ctx, key, expiry)
	return args.String(0), args.Error(1)
}
