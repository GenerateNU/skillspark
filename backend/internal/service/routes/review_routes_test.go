package routes_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	translatemocks "skillspark/internal/translation/mocks"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupReviewTestAPI(
	reviewRepo *repomocks.MockReviewRepository,
	regRepo *repomocks.MockRegistrationRepository,
	guardianRepo *repomocks.MockGuardianRepository,
	eventRepo *repomocks.MockEventRepository,
	translateClient *translatemocks.TranslateMock,
) (*fiber.App, huma.API) {

	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		Review:       reviewRepo,
		Registration: regRepo,
		Guardian:     guardianRepo,
		Event:        eventRepo,
	}

	routes.SetUpReviewRoutes(api, repo, translateClient)

	return app, api
}

func TestCreateReview_Success(t *testing.T) {
	t.Parallel()

	regID := uuid.New()
	guardianID := uuid.New()
	reviewID := uuid.New()

	translateRepo := new(translatemocks.TranslateMock)
	translated := "งานดี"
	translateRepo.
		On("GetTranslation", mock.Anything, "Great event").
		Return(&translated, nil)

	regRepo := new(repomocks.MockRegistrationRepository)
	regRepo.
		On("GetRegistrationByID", mock.Anything, mock.Anything, mock.Anything).
		Return(&models.GetRegistrationByIDOutput{
			Body: models.Registration{
				ID: regID,
			},
		}, nil)

	guardianRepo := new(repomocks.MockGuardianRepository)
	guardianRepo.
		On("GetGuardianByID", mock.Anything, guardianID).
		Return(&models.Guardian{ID: guardianID}, nil)

	reviewRepo := new(repomocks.MockReviewRepository)
	reviewRepo.
		On("CreateReview", mock.Anything, mock.Anything).
		Return(&models.Review{ID: reviewID}, nil)

	app, _ := setupReviewTestAPI(
		reviewRepo,
		regRepo,
		guardianRepo,
		new(repomocks.MockEventRepository),
		translateRepo,
	)

	payload := map[string]interface{}{
		"registration_id": regID.String(),
		"guardian_id":     guardianID.String(),
		"description":     "Great event",
		"categories":      []string{"fun"},
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/api/v1/review",
		bytes.NewBuffer(body),
	)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			slog.Error("Failed to close transaction: " + closeErr.Error())
		}
	}()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	regRepo.AssertExpectations(t)
	guardianRepo.AssertExpectations(t)
	reviewRepo.AssertExpectations(t)
	translateRepo.AssertExpectations(t)
}

func TestGetReviewsByEventID(t *testing.T) {
	t.Parallel()

	eventID := uuid.New()

	eventRepo := new(repomocks.MockEventRepository)
	eventRepo.
		On("GetEventByID", mock.Anything, eventID).
		Return(&models.Event{ID: eventID}, nil)

	reviewRepo := new(repomocks.MockReviewRepository)
	reviewRepo.
		On(
			"GetReviewsByEventID",
			mock.Anything, // context
			eventID,       // uuid
			"en-US",       // AcceptLanguage
			mock.Anything, // utils.Pagination
		).
		Return([]models.Review{}, nil)

	app, _ := setupReviewTestAPI(
		reviewRepo,
		new(repomocks.MockRegistrationRepository),
		new(repomocks.MockGuardianRepository),
		eventRepo,
		new(translatemocks.TranslateMock),
	)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/api/v1/review/event/"+eventID.String(),
		nil,
	)
	req.Header.Set("Accept-Language", "en-US")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			slog.Error("Failed to close transaction: " + closeErr.Error())
		}
	}()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	eventRepo.AssertExpectations(t)
	reviewRepo.AssertExpectations(t)
}

func TestGetReviewsByGuardianID(t *testing.T) {
	t.Parallel()

	guardianID := uuid.New()

	guardianRepo := new(repomocks.MockGuardianRepository)
	guardianRepo.
		On("GetGuardianByID", mock.Anything, guardianID).
		Return(&models.Guardian{ID: guardianID}, nil)

	reviewRepo := new(repomocks.MockReviewRepository)
	reviewRepo.
		On(
			"GetReviewsByGuardianID",
			mock.Anything, // context
			guardianID,    // uuid
			"en-US",       // AcceptLanguage
			mock.Anything, // utils.Pagination
		).
		Return([]models.Review{}, nil)

	app, _ := setupReviewTestAPI(
		reviewRepo,
		new(repomocks.MockRegistrationRepository),
		guardianRepo,
		new(repomocks.MockEventRepository),
		new(translatemocks.TranslateMock),
	)

	req, _ := http.NewRequest(
		http.MethodGet,
		"/api/v1/review/guardian/"+guardianID.String(),
		nil,
	)
	req.Header.Set("Accept-Language", "en-US")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			slog.Error("Failed to close transaction: " + closeErr.Error())
		}
	}()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	guardianRepo.AssertExpectations(t)
	reviewRepo.AssertExpectations(t)
}

func TestDeleteReview(t *testing.T) {
	t.Parallel()

	reviewID := uuid.New()

	reviewRepo := new(repomocks.MockReviewRepository)
	reviewRepo.
		On("DeleteReview", mock.Anything, reviewID).
		Return(nil)

	app, _ := setupReviewTestAPI(
		reviewRepo,
		new(repomocks.MockRegistrationRepository),
		new(repomocks.MockGuardianRepository),
		new(repomocks.MockEventRepository),
		new(translatemocks.TranslateMock),
	)

	req, _ := http.NewRequest(
		http.MethodDelete,
		"/api/v1/review/"+reviewID.String(),
		nil,
	)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			slog.Error("Failed to close transaction: " + closeErr.Error())
		}
	}()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	reviewRepo.AssertExpectations(t)
}
