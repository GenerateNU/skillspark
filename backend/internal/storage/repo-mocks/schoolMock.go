package repomocks

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/stretchr/testify/mock"
)

type MockSchoolRepository struct {
	mock.Mock
}

func (m *MockSchoolRepository) CreateSchool(ctx context.Context, input *models.CreateSchoolInput) (*models.School, error) {

	args := m.Called(ctx, input)

	var school *models.School
	if v := args.Get(0); v != nil {
		school = v.(*models.School)
	}

	var err error
	if v := args.Get(1); v != nil {
		err = v.(error)
	}

	return school, err
}

func (m *MockSchoolRepository) GetAllSchools(ctx context.Context, pagination utils.Pagination) ([]models.School, error) {
	args := m.Called(ctx, pagination)

	// Handle nil slice and/or error safely
	var schools []models.School
	if v := args.Get(0); v != nil {
		schools = v.([]models.School)
	}

	var err error
	if v := args.Get(1); v != nil {
		err = v.(error)
	}

	return schools, err
}
