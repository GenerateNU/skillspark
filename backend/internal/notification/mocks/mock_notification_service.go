package notificationmocks

import (
	"context"
	"skillspark/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockNotificationService struct {
	mock.Mock
}

func (m *MockNotificationService) SendNotification(ctx context.Context, input *models.SendNotificationInput) error {
	args := m.Called(ctx, input)
	return args.Error(0)
}

func (m *MockNotificationService) ScheduleNotification(ctx context.Context, input *models.CreateScheduledNotificationInput) (*models.Notification, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Notification), args.Error(1)
}
