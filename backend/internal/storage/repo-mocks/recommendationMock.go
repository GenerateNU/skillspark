package repomocks

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockRecommendationRepository struct {
	mock.Mock
}

func (m *MockRecommendationRepository) GetRecommendationsByChildID(ctx context.Context, childInterests []string, childBirthYear int, acceptLanguage string, pagination utils.Pagination, minDate *time.Time, maxDate *time.Time) ([]models.Event, error) {
	args := m.Called(ctx, childInterests, childBirthYear, acceptLanguage, pagination, minDate, maxDate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Event), nil
}
