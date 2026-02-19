package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type TranslateMock struct {
	mock.Mock
}

func (m *TranslateMock) GetTranslation(ctx context.Context, input string, sl string, dl string) (*string, error) {
	args := m.Called(ctx, input, sl, dl)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*string), nil
}
