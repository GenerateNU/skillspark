package repomocks

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockSchoolRepository struct {
	mock.Mock
}

func (m *MockSchoolRepository) GetAllSchools(ctx context.Context) ([]models.School, *errs.HTTPError) {
	args := m.Called(ctx)

	// Handle nil slice and/or error safely
	var schools []models.School
	if v := args.Get(0); v != nil {
		schools = v.([]models.School)
	}

	var httpErr *errs.HTTPError
	if v := args.Get(1); v != nil {
		httpErr = v.(*errs.HTTPError)
	}

	return schools, httpErr
}

