package routes

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"skillspark/internal/models"
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

func setupRegistrationTestAPI(
	registrationRepo *repomocks.MockRegistrationRepository,
	childRepo *repomocks.MockChildRepository,
	guardianRepo *repomocks.MockGuardianRepository,
	eventOccurrenceRepo *repomocks.MockEventOccurrenceRepository,
	organizationRepo *repomocks.MockOrganizationRepository,
	stripeClient *stripemocks.MockStripeClient,
) (*fiber.App, huma.API) {

	app := fiber.New()

	api := humafiber.New(app, huma.DefaultConfig("Test API", "1.0.0"))

	repo := &storage.Repository{
		Registration:    registrationRepo,
		Child:           childRepo,
		Guardian:        guardianRepo,
		EventOccurrence: eventOccurrenceRepo,
		Organization:    organizationRepo,
	}

	SetupRegistrationRoutes(api, repo, stripeClient)

	return app, api
}

func TestHumaValidation_GetRegistrationByID(t *testing.T) {
	t.Parallel()

	childIdEx := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianIdEx := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	tests := []struct {
		name       string
		ID         string
		mockSetup  func(*repomocks.MockRegistrationRepository)
		statusCode int
	}{
		{
			name: "valid UUID",
			ID:   "80000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				m.On(
					"GetRegistrationByID",
					mock.Anything,
					mock.AnythingOfType("*models.GetRegistrationByIDInput"),
				).Return(&models.GetRegistrationByIDOutput{
					Body: models.Registration{
						ID:                     uuid.MustParse("80000000-0000-0000-0000-000000000001"),
						ChildID:                childIdEx,
						GuardianID:             guardianIdEx,
						EventOccurrenceID:      uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                 models.RegistrationStatusRegistered,
						EventName:              "STEM Club",
						OccurrenceStartTime:    time.Now(),
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
						StripePaymentIntentID:  "pi_test_123",
						StripeCustomerID:       "cus_test_123",
						OrgStripeAccountID:     "acct_test_123",
						StripePaymentMethodID:  "pm_test_123",
						TotalAmount:            10000,
						ProviderAmount:         8500,
						PlatformFeeAmount:      1500,
						Currency:               "thb",
						PaymentIntentStatus:    "requires_capture",
					},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			ID:         "not-a-uuid",
			mockSetup:  func(*repomocks.MockRegistrationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/registrations/"+tt.ID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetRegistrationsByChildID(t *testing.T) {
	t.Parallel()

	childIdEx := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianIdEx := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	tests := []struct {
		name       string
		childID    string
		mockSetup  func(*repomocks.MockRegistrationRepository)
		statusCode int
	}{
		{
			name:    "valid UUID",
			childID: "30000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByChildIDOutput{}
				output.Body.Registrations = []models.Registration{
					{
						ID:                     uuid.New(),
						ChildID:                childIdEx,
						GuardianID:             guardianIdEx,
						EventOccurrenceID:      uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                 models.RegistrationStatusRegistered,
						EventName:              "STEM Club",
						OccurrenceStartTime:    time.Now(),
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
						StripePaymentIntentID:  "pi_test_123",
						StripeCustomerID:       "cus_test_123",
						OrgStripeAccountID:     "acct_test_123",
						StripePaymentMethodID:  "pm_test_123",
						TotalAmount:            10000,
						ProviderAmount:         8500,
						PlatformFeeAmount:      1500,
						Currency:               "thb",
						PaymentIntentStatus:    "requires_capture",
					},
				}
				m.On("GetRegistrationsByChildID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByChildIDInput")).Return(output, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			childID:    "not-a-uuid",
			mockSetup:  func(*repomocks.MockRegistrationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/registrations/child/"+tt.childID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetRegistrationsByGuardianID(t *testing.T) {
	t.Parallel()

	childIdEx := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianIdEx := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	tests := []struct {
		name       string
		guardianID string
		mockSetup  func(*repomocks.MockRegistrationRepository)
		statusCode int
	}{
		{
			name:       "valid UUID",
			guardianID: "11111111-1111-1111-1111-111111111111",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByGuardianIDOutput{}
				output.Body.Registrations = []models.Registration{
					{
						ID:                     uuid.New(),
						ChildID:                childIdEx,
						GuardianID:             guardianIdEx,
						EventOccurrenceID:      uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                 models.RegistrationStatusRegistered,
						EventName:              "STEM Club",
						OccurrenceStartTime:    time.Now(),
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
						StripePaymentIntentID:  "pi_test_123",
						StripeCustomerID:       "cus_test_123",
						OrgStripeAccountID:     "acct_test_123",
						StripePaymentMethodID:  "pm_test_123",
						TotalAmount:            10000,
						ProviderAmount:         8500,
						PlatformFeeAmount:      1500,
						Currency:               "thb",
						PaymentIntentStatus:    "requires_capture",
					},
				}
				m.On("GetRegistrationsByGuardianID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByGuardianIDInput")).Return(output, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:       "invalid UUID",
			guardianID: "not-a-uuid",
			mockSetup:  func(*repomocks.MockRegistrationRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/registrations/guardian/"+tt.guardianID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_GetRegistrationsByEventOccurrenceID(t *testing.T) {
	t.Parallel()

	childIdEx := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianIdEx := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	tests := []struct {
		name              string
		eventOccurrenceID string
		mockSetup         func(*repomocks.MockRegistrationRepository)
		statusCode        int
	}{
		{
			name:              "valid UUID",
			eventOccurrenceID: "70000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockRegistrationRepository) {
				output := &models.GetRegistrationsByEventOccurrenceIDOutput{}
				output.Body.Registrations = []models.Registration{
					{
						ID:                     uuid.New(),
						ChildID:                childIdEx,
						GuardianID:             guardianIdEx,
						EventOccurrenceID:      uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                 models.RegistrationStatusRegistered,
						EventName:              "STEM Club",
						OccurrenceStartTime:    time.Now(),
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
						StripePaymentIntentID:  "pi_test_123",
						StripeCustomerID:       "cus_test_123",
						OrgStripeAccountID:     "acct_test_123",
						StripePaymentMethodID:  "pm_test_123",
						TotalAmount:            10000,
						ProviderAmount:         8500,
						PlatformFeeAmount:      1500,
						Currency:               "thb",
						PaymentIntentStatus:    "requires_capture",
					},
				}
				m.On("GetRegistrationsByEventOccurrenceID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationsByEventOccurrenceIDInput")).Return(output, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name:              "invalid UUID",
			eventOccurrenceID: "not-a-uuid",
			mockSetup:         func(*repomocks.MockRegistrationRepository) {},
			statusCode:        http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

			req, err := http.NewRequest(
				http.MethodGet,
				"/api/v1/registrations/event_occurrence/"+tt.eventOccurrenceID,
				nil,
			)
			assert.NoError(t, err)

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_CreateRegistration(t *testing.T) {
	t.Parallel()

	childID := "30000000-0000-0000-0000-000000000001"
	guardianID := "11111111-1111-1111-1111-111111111111"
	eventOccurrenceID := "70000000-0000-0000-0000-000000000001"
	organizationID := uuid.MustParse("10000000-0000-0000-0000-000000000001")

	childIDEx := uuid.MustParse(childID)
	guardianIDEx := uuid.MustParse(guardianID)
	eventOccurrenceIDEx := uuid.MustParse(eventOccurrenceID)

	stripeCustomerID := "cus_test_123"
	stripeAccountID := "acct_test_123"
	paymentMethodID := "pm_test_123"

	validEventOccurrence := &models.EventOccurrence{
		ID:        eventOccurrenceIDEx,
		Price:     10000,
		StartTime: time.Now().Add(25 * time.Hour),
		Event: models.Event{
			ID:             uuid.New(),
			OrganizationID: organizationID,
			Title:          "STEM Club",
		},
	}

	tests := []struct {
		name       string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient)
		statusCode int
	}{
		{
			name: "valid payload",
			payload: map[string]interface{}{
				"child_id":            childID,
				"guardian_id":         guardianID,
				"event_occurrence_id": eventOccurrenceID,
				"payment_method_id":   paymentMethodID,
				"currency":            "thb",
				"status":              "registered",
				"amount":              10000,
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceIDEx).
					Return(validEventOccurrence, nil)

				guardianRepo.On("GetGuardianByID", mock.Anything, guardianIDEx).
					Return(&models.Guardian{
						ID:               guardianIDEx,
						StripeCustomerID: &stripeCustomerID,
					}, nil)

				childRepo.On("GetChildByID", mock.Anything, childIDEx).
					Return(&models.Child{
						ID:         childIDEx,
						GuardianID: guardianIDEx, // must match guardian
					}, nil)

				orgRepo.On("GetOrganizationByID", mock.Anything, organizationID).
					Return(&models.Organization{
						ID:              organizationID,
						StripeAccountID: &stripeAccountID,
					}, nil)

				sc.On("CreatePaymentIntent", mock.Anything, mock.AnythingOfType("*models.CreatePaymentIntentInput")).
					Return(&models.CreatePaymentIntentOutput{
						Body: struct {
							PaymentIntentID   string `json:"payment_intent_id" doc:"Stripe payment intent ID"`
							ClientSecret      string `json:"client_secret" doc:"Client secret for frontend to confirm payment"`
							Status            string `json:"status" doc:"Payment intent status"`
							TotalAmount       int    `json:"total_amount" doc:"Total amount in cents"`
							ProviderAmount    int    `json:"provider_amount" doc:"Amount provider receives in cents"`
							PlatformFeeAmount int    `json:"platform_fee_amount" doc:"Platform fee in cents"`
							Currency          string `json:"currency" doc:"Currency code"`
						}{
							PaymentIntentID:   "pi_test_123",
							TotalAmount:       10000,
							ProviderAmount:    8500,
							PlatformFeeAmount: 1500,
							Currency:          "thb",
							Status:            "requires_capture",
						},
					}, nil)

				regRepo.On("CreateRegistration", mock.Anything, mock.AnythingOfType("*models.CreateRegistrationWithPaymentData")).
					Return(&models.CreateRegistrationOutput{
						Body: models.Registration{
							ID:                    uuid.New(),
							ChildID:               childIDEx,
							GuardianID:            guardianIDEx,
							EventOccurrenceID:     eventOccurrenceIDEx,
							Status:                models.RegistrationStatusRegistered,
							EventName:             "STEM Club",
							OccurrenceStartTime:   time.Now(),
							CreatedAt:             time.Now(),
							UpdatedAt:             time.Now(),
							StripePaymentIntentID: "pi_test_123",
							StripeCustomerID:      stripeCustomerID,
							OrgStripeAccountID:    stripeAccountID,
							StripePaymentMethodID: paymentMethodID,
							TotalAmount:           10000,
							ProviderAmount:        8500,
							PlatformFeeAmount:     1500,
							Currency:              "thb",
							PaymentIntentStatus:   "requires_capture",
						},
					}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "missing child_id",
			payload: map[string]interface{}{
				"guardian_id":         guardianID,
				"event_occurrence_id": eventOccurrenceID,
				"payment_method_id":   paymentMethodID,
				"currency":            "thb",
				"status":              "registered",
			},
			mockSetup:  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "missing guardian_id",
			payload: map[string]interface{}{
				"child_id":            childID,
				"event_occurrence_id": eventOccurrenceID,
				"payment_method_id":   paymentMethodID,
				"currency":            "thb",
				"status":              "registered",
			},
			mockSetup:  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "missing event_occurrence_id",
			payload: map[string]interface{}{
				"child_id":          childID,
				"guardian_id":       guardianID,
				"payment_method_id": paymentMethodID,
				"currency":          "thb",
				"status":            "registered",
			},
			mockSetup:  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid child_id format",
			payload: map[string]interface{}{
				"child_id":            "not-a-uuid",
				"guardian_id":         guardianID,
				"event_occurrence_id": eventOccurrenceID,
				"payment_method_id":   paymentMethodID,
				"currency":            "thb",
				"status":              "registered",
			},
			mockSetup:  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid guardian_id format",
			payload: map[string]interface{}{
				"child_id":            childID,
				"guardian_id":         "not-a-uuid",
				"event_occurrence_id": eventOccurrenceID,
				"payment_method_id":   paymentMethodID,
				"currency":            "thb",
				"status":              "registered",
			},
			mockSetup:  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid status value",
			payload: map[string]interface{}{
				"child_id":            childID,
				"guardian_id":         guardianID,
				"event_occurrence_id": eventOccurrenceID,
				"payment_method_id":   paymentMethodID,
				"currency":            "thb",
				"status":              "not-a-valid-status",
			},
			mockSetup:  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/api/v1/registrations", bytes.NewBuffer(bodyBytes))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			if tt.statusCode != resp.StatusCode {
				body, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(body))
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
			mockChildRepo.AssertExpectations(t)
			mockGuardianRepo.AssertExpectations(t)
			mockEORepo.AssertExpectations(t)
			mockOrgRepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_UpdateRegistration(t *testing.T) {
	t.Parallel()

	childIdEx := uuid.MustParse("30000000-0000-0000-0000-000000000002")
	guardianIdEx := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	registrationID := "80000000-0000-0000-0000-000000000001"

	tests := []struct {
		name       string
		id         string
		payload    map[string]interface{}
		mockSetup  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository)
		statusCode int
	}{
		{
			name: "valid child_id update",
			id:   registrationID,
			payload: map[string]interface{}{
				"child_id": "30000000-0000-0000-0000-000000000002",
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository) {
				childRepo.On("GetChildByID", mock.Anything, childIdEx).Return(&models.Child{
					ID: childIdEx,
				}, nil)
				
				regRepo.On(
					"UpdateRegistration",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateRegistrationInput"),
				).Return(&models.UpdateRegistrationOutput{
					Body: models.Registration{
						ID:                     uuid.MustParse(registrationID),
						ChildID:                childIdEx,
						GuardianID:             guardianIdEx,
						EventOccurrenceID:      uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                 models.RegistrationStatusRegistered,
						EventName:              "STEM Club",
						OccurrenceStartTime:    time.Now(),
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
						StripePaymentIntentID:  "pi_test_123",
						StripeCustomerID:       "cus_test_123",
						OrgStripeAccountID:     "acct_test_123",
						StripePaymentMethodID:  "pm_test_123",
						TotalAmount:            10000,
						ProviderAmount:         8500,
						PlatformFeeAmount:      1500,
						Currency:               "thb",
						PaymentIntentStatus:    "requires_capture",
					},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "valid status update",
			id:   registrationID,
			payload: map[string]interface{}{
				"status": "cancelled",
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository) {
				regRepo.On(
					"UpdateRegistration",
					mock.Anything,
					mock.AnythingOfType("*models.UpdateRegistrationInput"),
				).Return(&models.UpdateRegistrationOutput{
					Body: models.Registration{
						ID:                     uuid.MustParse(registrationID),
						ChildID:                childIdEx,
						GuardianID:             guardianIdEx,
						EventOccurrenceID:      uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                 models.RegistrationStatusCancelled,
						EventName:              "STEM Club",
						OccurrenceStartTime:    time.Now(),
						CreatedAt:              time.Now(),
						UpdatedAt:              time.Now(),
						StripePaymentIntentID:  "pi_test_123",
						StripeCustomerID:       "cus_test_123",
						OrgStripeAccountID:     "acct_test_123",
						StripePaymentMethodID:  "pm_test_123",
						TotalAmount:            10000,
						ProviderAmount:         8500,
						PlatformFeeAmount:      1500,
						Currency:               "thb",
						PaymentIntentStatus:    "requires_capture",
					},
				}, nil)
			},
			statusCode: http.StatusOK,
		},
		{
			name: "invalid id format",
			id:   "not-a-uuid",
			payload: map[string]interface{}{
				"status": "cancelled",
			},
			mockSetup:  func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository) {},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRegRepo := new(repomocks.MockRegistrationRepository)
			mockChildRepo := new(repomocks.MockChildRepository)
			mockGuardianRepo := new(repomocks.MockGuardianRepository)
			mockEORepo := new(repomocks.MockEventOccurrenceRepository)
			mockOrgRepo := new(repomocks.MockOrganizationRepository)
			mockStripeClient := new(stripemocks.MockStripeClient)
			tt.mockSetup(mockRegRepo, mockChildRepo)

			app, _ := setupRegistrationTestAPI(mockRegRepo, mockChildRepo, mockGuardianRepo, mockEORepo, mockOrgRepo, mockStripeClient)

			bodyBytes, err := json.Marshal(tt.payload)
			assert.NoError(t, err)

			req, err := http.NewRequest(
				http.MethodPatch,
				"/api/v1/registrations/"+tt.id,
				bytes.NewBuffer(bodyBytes),
			)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			if tt.statusCode != resp.StatusCode {
				bodyBytes, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(bodyBytes))
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
		})
	}
}