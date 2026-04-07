package routes_test

import (
	"io"
	"net/http"
	"testing"
	"time"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRecommendationTestAPI(
	childRepo *repomocks.MockChildRepository,
	recommendationRepo *repomocks.MockRecommendationRepository,
) (*fiber.App, huma.API) {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))
	repo := &storage.Repository{
		Child:          childRepo,
		Recommendation: recommendationRepo,
	}
	routes.SetupRecommendationRoutes(api, repo)
	return app, api
}

func TestGetRecommendationsByChildID(t *testing.T) {
	t.Parallel()

	childID := uuid.MustParse("10000000-0000-0000-0000-000000000001")
	eight := 8
	twelve := 12

	child := &models.Child{
		ID:         childID,
		Name:       "Test Child",
		BirthYear:  2015,
		BirthMonth: 3,
		Interests:  []string{"science", "technology"},
	}

	events := []models.Event{
		{
			ID:          uuid.MustParse("60000000-0000-0000-0000-000000000001"),
			Title:       "Junior Robotics Workshop",
			Description: "Learn robotics!",
			AgeRangeMin: &eight,
			AgeRangeMax: &twelve,
			Category:    []string{"science", "technology"},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	defaultPagination := utils.Pagination{Page: 1, Limit: 10}

	tests := []struct {
		name       string
		childID    string
		mockSetup  func(*repomocks.MockChildRepository, *repomocks.MockRecommendationRepository)
		statusCode int
	}{
		{
			name:    "success",
			childID: childID.String(),
			mockSetup: func(c *repomocks.MockChildRepository, r *repomocks.MockRecommendationRepository) {
				c.On("GetChildByID", mock.Anything, childID).Return(child, nil)
				r.On("GetRecommendationsByChildID", mock.Anything, child.Interests, child.BirthYear, "en-US", defaultPagination, (*time.Time)(nil), (*time.Time)(nil)).Return(events, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:    "child not found",
			childID: childID.String(),
			mockSetup: func(c *repomocks.MockChildRepository, r *repomocks.MockRecommendationRepository) {
				notFound := errs.NotFound("Child", "id", childID)
				c.On("GetChildByID", mock.Anything, childID).Return(nil, &notFound)
			},
			statusCode: http.StatusNotFound,
		},
		{
			name:    "invalid child_id",
			childID: "not-a-uuid",
			mockSetup: func(c *repomocks.MockChildRepository, r *repomocks.MockRecommendationRepository) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name:    "repo error",
			childID: childID.String(),
			mockSetup: func(c *repomocks.MockChildRepository, r *repomocks.MockRecommendationRepository) {
				c.On("GetChildByID", mock.Anything, childID).Return(child, nil)
				repoErr := errs.InternalServerError("db error", "")
				r.On("GetRecommendationsByChildID", mock.Anything, child.Interests, child.BirthYear, "en-US", defaultPagination, (*time.Time)(nil), (*time.Time)(nil)).Return(nil, &repoErr)
			},
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockChild := new(repomocks.MockChildRepository)
			mockRec := new(repomocks.MockRecommendationRepository)
			tt.mockSetup(mockChild, mockRec)

			app, _ := setupRecommendationTestAPI(mockChild, mockRec)

			req, err := http.NewRequest(http.MethodGet, "/api/v1/recommendations/"+tt.childID, nil)
			assert.NoError(t, err)
			req.Header.Set("Accept-Language", "en-US")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			if tt.statusCode != resp.StatusCode {
				body, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(body))
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockChild.AssertExpectations(t)
			mockRec.AssertExpectations(t)
		})
	}
}
