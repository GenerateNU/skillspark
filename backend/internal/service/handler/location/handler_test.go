package location

import (
	"net/http/httptest"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetLocationById(t *testing.T) {
	tests := []struct {
		id             string
		mockSetup      func(*repomocks.MockLocationRepository)
		expectedStatus int
		wantErr        bool
	}{
		{
			id: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("GetLocationByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(&models.Location{
					ID:        uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					Latitude:  40.7128,
					Longitude: -74.0060,
					Address:   "123 Broadway",
					City:      "New York",
					State:     "NY",
					ZipCode:   "10001",
					Country:   "USA",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			expectedStatus: 200,
			wantErr:        false,
		},
		{
			id: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("GetLocationByID", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Internal server error").Code,
					Message: "Internal server error",
				})
			},
			expectedStatus: 500,
			wantErr:        true,
		}}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			app := fiber.New(fiber.Config{
				ErrorHandler: errs.ErrorHandler,
			})
			mockRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			app.Get("/locations/:id", func(c *fiber.Ctx) error {
				location, err := handler.GetLocationById(c.Context(), &models.GetLocationByIDInput{ID: uuid.MustParse(c.Params("id"))})
				if err != nil {
					return err
				}
				return c.Status(fiber.StatusOK).JSON(location)
			})

			req := httptest.NewRequest("GET", "/locations/"+tt.id, nil)
			res, _ := app.Test(req, -1)

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
			mockRepo.AssertExpectations(t)
		})
	}
}
