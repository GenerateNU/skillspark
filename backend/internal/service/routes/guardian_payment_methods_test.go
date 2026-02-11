package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"skillspark/internal/models"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	stripemocks "skillspark/internal/stripeClient/mocks"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupGuardianPaymentMethodTestAPI(
	guardianRepo *repomocks.MockGuardianRepository,
	paymentMethodRepo *repomocks.MockGuardianPaymentMethodRepository,
	stripeClient *stripemocks.MockStripeClient,
) (*fiber.App, huma.API) {

	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		Guardian:              guardianRepo,
		GuardianPaymentMethod: paymentMethodRepo,
	}

	routes.SetupGuardianPaymentMethodRoutes(api, repo, stripeClient)

	return app, api
}

func TestHumaValidation_GetGuardianPaymentMethods(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		guardianID string
		mockSetup  func(*repomocks.MockGuardianRepository, *repomocks.MockGuardianPaymentMethodRepository)
		statusCode int
	}{
		{
			name:       "valid guardian with payment methods",
			guardianID: "88888888-8888-8888-8888-888888888888",
			mockSetup: func(gr *repomocks.MockGuardianRepository, pmr *repomocks.MockGuardianPaymentMethodRepository) {
				stripeCustomerID := "cus_test123"
				gr.On("GetGuardianByID", mock.Anything, uuid.MustParse("88888888-8888-8888-8888-888888888888")).
					Return(&models.Guardian{
						ID:               uuid.MustParse("88888888-8888-8888-8888-888888888888"),
						StripeCustomerID: &stripeCustomerID,
					}, nil)

				cardBrand := "visa"
				cardLast4 := "4242"
				pmr.On("GetPaymentMethodsByGuardianID", mock.Anything, uuid.MustParse("88888888-8888-8888-8888-888888888888")).
					Return([]models.GuardianPaymentMethod{
						{
							ID:                    uuid.New(),
							GuardianID:            uuid.MustParse("88888888-8888-8888-8888-888888888888"),
							StripePaymentMethodID: "pm_test123",
							CardBrand:             &cardBrand,
							CardLast4:             &cardLast4,
							IsDefault:             true,
							CreatedAt:             time.Now(),
							UpdatedAt:             time.Now(),
						},
					}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			guardianID: "not-a-uuid",
			mockSetup:  func(*repomocks.MockGuardianRepository, *repomocks.MockGuardianPaymentMethodRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockPaymentMethodRepo := new(repomocks.MockGuardianPaymentMethodRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockGuardianRepo, mockPaymentMethodRepo)

			app, _ := setupGuardianPaymentMethodTestAPI(mockGuardianRepo, mockPaymentMethodRepo, mockStripeClient)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/guardians/"+tt.guardianID+"/payment-methods",
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockGuardianRepo.AssertExpectations(t)
			mockPaymentMethodRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_CreateGuardianPaymentMethod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockGuardianRepository, *repomocks.MockGuardianPaymentMethodRepository)
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"guardian_id":             "88888888-8888-8888-8888-888888888888",
				"stripe_payment_method_id": "pm_test123",
				"card_brand":              "visa",
				"card_last4":              "4242",
				"card_exp_month":          12,
				"card_exp_year":           2027,
				"is_default":              true,
			},
			mockSetup: func(gr *repomocks.MockGuardianRepository, pmr *repomocks.MockGuardianPaymentMethodRepository) {
				stripeCustomerID := "cus_test123"
				gr.On("GetGuardianByID", mock.Anything, uuid.MustParse("88888888-8888-8888-8888-888888888888")).
					Return(&models.Guardian{
						ID:               uuid.MustParse("88888888-8888-8888-8888-888888888888"),
						StripeCustomerID: &stripeCustomerID,
					}, nil)

				pmr.On("CreateGuardianPaymentMethod", mock.Anything, mock.AnythingOfType("*models.CreateGuardianPaymentMethodInput")).
					Return(&models.GuardianPaymentMethod{
						ID:                    uuid.New(),
						GuardianID:            uuid.MustParse("88888888-8888-8888-8888-888888888888"),
						StripePaymentMethodID: "pm_test123",
						IsDefault:             true,
						CreatedAt:             time.Now(),
						UpdatedAt:             time.Now(),
					}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "missing required fields",
			payload: map[string]interface{}{
				"guardian_id": "88888888-8888-8888-8888-888888888888",
			},
			mockSetup:  func(*repomocks.MockGuardianRepository, *repomocks.MockGuardianPaymentMethodRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockPaymentMethodRepo := new(repomocks.MockGuardianPaymentMethodRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockGuardianRepo, mockPaymentMethodRepo)

			app, _ := setupGuardianPaymentMethodTestAPI(mockGuardianRepo, mockPaymentMethodRepo, mockStripeClient)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPost,
				"/api/v1/guardians/payment-methods",
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockGuardianRepo.AssertExpectations(t)
			mockPaymentMethodRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_SetDefaultPaymentMethod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		guardianID      string
		paymentMethodID string
		mockSetup       func(*repomocks.MockGuardianRepository, *repomocks.MockGuardianPaymentMethodRepository)
		statusCode      int
	}{
		{
			name:            "valid IDs",
			guardianID:      "88888888-8888-8888-8888-888888888888",
			paymentMethodID: "11111111-1111-1111-1111-111111111111",
			mockSetup: func(gr *repomocks.MockGuardianRepository, pmr *repomocks.MockGuardianPaymentMethodRepository) {
				stripeCustomerID := "cus_test123"
				gr.On("GetGuardianByID", mock.Anything, uuid.MustParse("88888888-8888-8888-8888-888888888888")).
					Return(&models.Guardian{
						ID:               uuid.MustParse("88888888-8888-8888-8888-888888888888"),
						StripeCustomerID: &stripeCustomerID,
					}, nil)

				pmr.On("UpdateGuardianPaymentMethod", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111111"), true).
					Return(&models.GuardianPaymentMethod{
						ID:         uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						GuardianID: uuid.MustParse("88888888-8888-8888-8888-888888888888"),
						IsDefault:  true,
					}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:            "invalid guardian UUID",
			guardianID:      "not-a-uuid",
			paymentMethodID: "11111111-1111-1111-1111-111111111111",
			mockSetup:       func(*repomocks.MockGuardianRepository, *repomocks.MockGuardianPaymentMethodRepository) {},
			statusCode:      http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockPaymentMethodRepo := new(repomocks.MockGuardianPaymentMethodRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockGuardianRepo, mockPaymentMethodRepo)

			app, _ := setupGuardianPaymentMethodTestAPI(mockGuardianRepo, mockPaymentMethodRepo, mockStripeClient)

			req, err := http.NewRequest(
				http.MethodPatch,
				"/api/v1/guardians/"+tt.guardianID+"/payment-methods/"+tt.paymentMethodID+"/default",
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockGuardianRepo.AssertExpectations(t)
			mockPaymentMethodRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_DeleteGuardianPaymentMethod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		pmID       string
		mockSetup  func(*repomocks.MockGuardianPaymentMethodRepository, *stripemocks.MockStripeClient)
		statusCode int
	}{
		{
			name: "valid UUID",
			pmID: "11111111-1111-1111-1111-111111111111",
			mockSetup: func(pmr *repomocks.MockGuardianPaymentMethodRepository, sc *stripemocks.MockStripeClient) {
				pmr.On("DeleteGuardianPaymentMethod", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(&models.GuardianPaymentMethod{
						ID:                    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						StripePaymentMethodID: "pm_test123",
					}, nil)

				sc.On("DetachPaymentMethod", mock.Anything, "pm_test123").Return(nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			pmID:       "not-a-uuid",
			mockSetup:  func(*repomocks.MockGuardianPaymentMethodRepository, *stripemocks.MockStripeClient) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockPaymentMethodRepo := new(repomocks.MockGuardianPaymentMethodRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockPaymentMethodRepo, mockStripeClient)

			app, _ := setupGuardianPaymentMethodTestAPI(mockGuardianRepo, mockPaymentMethodRepo, mockStripeClient)

			req, err := http.NewRequest(
				http.MethodDelete,
				"/api/v1/guardians/payment-methods/"+tt.pmID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockPaymentMethodRepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
		})
	}
}