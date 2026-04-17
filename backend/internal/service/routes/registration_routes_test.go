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
	SetupRegistrationRoutes(api, repo, stripeClient, nil)
	return app, api
}

func TestHumaValidation_GetRegistrationByID(t *testing.T) {
	t.Parallel()

	childIDEx := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianIDEx := uuid.MustParse("11111111-1111-1111-1111-111111111111")

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
				m.On("GetRegistrationByID", mock.Anything, mock.AnythingOfType("*models.GetRegistrationByIDInput")).
					Return(&models.GetRegistrationByIDOutput{
						Body: models.Registration{
							ID:                    uuid.MustParse("80000000-0000-0000-0000-000000000001"),
							ChildID:               childIDEx,
							GuardianID:            guardianIDEx,
							EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
							Status:                models.RegistrationStatusRegistered,
							EventName:             "STEM Club",
							OccurrenceStartTime:   time.Now(),
							CreatedAt:             time.Now(),
							UpdatedAt:             time.Now(),
							StripePaymentIntentID: "pi_test_123",
							StripeCustomerID:      "cus_test_123",
							OrgStripeAccountID:    "acct_test_123",
							StripePaymentMethodID: "pm_test_123",
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

			req, err := http.NewRequest(http.MethodGet, "/api/v1/registrations/"+tt.ID, nil)
			assert.NoError(t, err)
			req.Header.Set("Accept-Language", "en-US")

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

	childIDEx := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianIDEx := uuid.MustParse("11111111-1111-1111-1111-111111111111")

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
						ID:                    uuid.New(),
						ChildID:               childIDEx,
						GuardianID:            guardianIDEx,
						EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                models.RegistrationStatusRegistered,
						EventName:             "STEM Club",
						OccurrenceStartTime:   time.Now(),
						CreatedAt:             time.Now(),
						UpdatedAt:             time.Now(),
						StripePaymentIntentID: "pi_test_123",
						StripeCustomerID:      "cus_test_123",
						OrgStripeAccountID:    "acct_test_123",
						StripePaymentMethodID: "pm_test_123",
						TotalAmount:           10000,
						ProviderAmount:        8500,
						PlatformFeeAmount:     1500,
						Currency:              "thb",
						PaymentIntentStatus:   "requires_capture",
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

			req, err := http.NewRequest(http.MethodGet, "/api/v1/registrations/child/"+tt.childID, nil)
			assert.NoError(t, err)
			req.Header.Set("Accept-Language", "en-US")

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

	childIDEx := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianIDEx := uuid.MustParse("11111111-1111-1111-1111-111111111111")

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
						ID:                    uuid.New(),
						ChildID:               childIDEx,
						GuardianID:            guardianIDEx,
						EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                models.RegistrationStatusRegistered,
						EventName:             "STEM Club",
						OccurrenceStartTime:   time.Now(),
						CreatedAt:             time.Now(),
						UpdatedAt:             time.Now(),
						StripePaymentIntentID: "pi_test_123",
						StripeCustomerID:      "cus_test_123",
						OrgStripeAccountID:    "acct_test_123",
						StripePaymentMethodID: "pm_test_123",
						TotalAmount:           10000,
						ProviderAmount:        8500,
						PlatformFeeAmount:     1500,
						Currency:              "thb",
						PaymentIntentStatus:   "requires_capture",
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

			req, err := http.NewRequest(http.MethodGet, "/api/v1/registrations/guardian/"+tt.guardianID, nil)
			assert.NoError(t, err)
			req.Header.Set("Accept-Language", "en-US")

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

	childIDEx := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	guardianIDEx := uuid.MustParse("11111111-1111-1111-1111-111111111111")

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
						ID:                    uuid.New(),
						ChildID:               childIDEx,
						GuardianID:            guardianIDEx,
						EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
						Status:                models.RegistrationStatusRegistered,
						EventName:             "STEM Club",
						OccurrenceStartTime:   time.Now(),
						CreatedAt:             time.Now(),
						UpdatedAt:             time.Now(),
						StripePaymentIntentID: "pi_test_123",
						StripeCustomerID:      "cus_test_123",
						OrgStripeAccountID:    "acct_test_123",
						StripePaymentMethodID: "pm_test_123",
						TotalAmount:           10000,
						ProviderAmount:        8500,
						PlatformFeeAmount:     1500,
						Currency:              "thb",
						PaymentIntentStatus:   "requires_capture",
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

			req, err := http.NewRequest(http.MethodGet, "/api/v1/registrations/event_occurrence/"+tt.eventOccurrenceID, nil)
			assert.NoError(t, err)
			req.Header.Set("Accept-Language", "en-US")

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
	paymentMethodID := "pm_test_123"

	validEventOccurrence := &models.EventOccurrence{
		ID:           eventOccurrenceIDEx,
		Price:        10000,
		Currency:     "thb",
		StartTime:    time.Now().Add(25 * time.Hour),
		CurrEnrolled: 5,
		MaxAttendees: 15,
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
				"status":              "registered",
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository, guardianRepo *repomocks.MockGuardianRepository, eoRepo *repomocks.MockEventOccurrenceRepository, orgRepo *repomocks.MockOrganizationRepository, sc *stripemocks.MockStripeClient) {
				eoRepo.On("GetEventOccurrenceByID", mock.Anything, eventOccurrenceIDEx, mock.Anything).
					Return(validEventOccurrence, nil)

				childRepo.On("GetChildByID", mock.Anything, childIDEx).
					Return(&models.Child{
						ID:         childIDEx,
						GuardianID: guardianIDEx,
					}, nil)

				guardianRepo.On("GetGuardianByID", mock.Anything, guardianIDEx).
					Return(&models.Guardian{
						ID:               guardianIDEx,
						StripeCustomerID: &stripeCustomerID,
					}, nil)

				regRepo.On("CreateRegistration", mock.Anything, mock.AnythingOfType("*models.CreateRegistrationData")).
					Return(&models.CreateRegistrationOutput{
						Body: models.Registration{
							ID:                  uuid.New(),
							ChildID:             childIDEx,
							GuardianID:          guardianIDEx,
							EventOccurrenceID:   eventOccurrenceIDEx,
							Status:              models.RegistrationStatusRegistered,
							EventName:           "STEM Club",
							OccurrenceStartTime: time.Now(),
							CreatedAt:           time.Now(),
							UpdatedAt:           time.Now(),
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
				"status":              "registered",
			},
			mockSetup: func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "missing guardian_id",
			payload: map[string]interface{}{
				"child_id":            childID,
				"event_occurrence_id": eventOccurrenceID,
				"payment_method_id":   paymentMethodID,
				"status":              "registered",
			},
			mockSetup: func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "missing event_occurrence_id",
			payload: map[string]interface{}{
				"child_id":          childID,
				"guardian_id":       guardianID,
				"payment_method_id": paymentMethodID,
				"status":            "registered",
			},
			mockSetup: func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid child_id format",
			payload: map[string]interface{}{
				"child_id":            "not-a-uuid",
				"guardian_id":         guardianID,
				"event_occurrence_id": eventOccurrenceID,
				"payment_method_id":   paymentMethodID,
				"status":              "registered",
			},
			mockSetup: func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid guardian_id format",
			payload: map[string]interface{}{
				"child_id":            childID,
				"guardian_id":         "not-a-uuid",
				"event_occurrence_id": eventOccurrenceID,
				"payment_method_id":   paymentMethodID,
				"status":              "registered",
			},
			mockSetup: func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {
			},
			statusCode: http.StatusUnprocessableEntity,
		},
		{
			name: "invalid status value",
			payload: map[string]interface{}{
				"child_id":            childID,
				"guardian_id":         guardianID,
				"event_occurrence_id": eventOccurrenceID,
				"payment_method_id":   paymentMethodID,
				"status":              "not-a-valid-status",
			},
			mockSetup: func(*repomocks.MockRegistrationRepository, *repomocks.MockChildRepository, *repomocks.MockGuardianRepository, *repomocks.MockEventOccurrenceRepository, *repomocks.MockOrganizationRepository, *stripemocks.MockStripeClient) {
			},
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
			req.Header.Set("Accept-Language", "en-US")

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

	childIDEx := uuid.MustParse("30000000-0000-0000-0000-000000000002")
	guardianIDEx := uuid.MustParse("11111111-1111-1111-1111-111111111111")
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
				childRepo.On("GetChildByID", mock.Anything, childIDEx).Return(&models.Child{
					ID: childIDEx,
				}, nil)

				regRepo.On("UpdateRegistration", mock.Anything, mock.AnythingOfType("*models.UpdateRegistrationInput")).
					Return(&models.UpdateRegistrationOutput{
						Body: models.Registration{
							ID:                    uuid.MustParse(registrationID),
							ChildID:               childIDEx,
							GuardianID:            guardianIDEx,
							EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
							Status:                models.RegistrationStatusRegistered,
							EventName:             "STEM Club",
							OccurrenceStartTime:   time.Now(),
							CreatedAt:             time.Now(),
							UpdatedAt:             time.Now(),
							StripePaymentIntentID: "pi_test_123",
							StripeCustomerID:      "cus_test_123",
							OrgStripeAccountID:    "acct_test_123",
							StripePaymentMethodID: "pm_test_123",
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
			name: "valid status update",
			id:   registrationID,
			payload: map[string]interface{}{
				"status": "cancelled",
			},
			mockSetup: func(regRepo *repomocks.MockRegistrationRepository, childRepo *repomocks.MockChildRepository) {
				regRepo.On("UpdateRegistration", mock.Anything, mock.AnythingOfType("*models.UpdateRegistrationInput")).
					Return(&models.UpdateRegistrationOutput{
						Body: models.Registration{
							ID:                    uuid.MustParse(registrationID),
							ChildID:               childIDEx,
							GuardianID:            guardianIDEx,
							EventOccurrenceID:     uuid.MustParse("70000000-0000-0000-0000-000000000001"),
							Status:                models.RegistrationStatusCancelled,
							EventName:             "STEM Club",
							OccurrenceStartTime:   time.Now(),
							CreatedAt:             time.Now(),
							UpdatedAt:             time.Now(),
							StripePaymentIntentID: "pi_test_123",
							StripeCustomerID:      "cus_test_123",
							OrgStripeAccountID:    "acct_test_123",
							StripePaymentMethodID: "pm_test_123",
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

			req, err := http.NewRequest(http.MethodPatch, "/api/v1/registrations/"+tt.id, bytes.NewBuffer(bodyBytes))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept-Language", "en-US")

			resp, err := app.Test(req)
			assert.NoError(t, err)
			defer func() { _ = resp.Body.Close() }()

			if tt.statusCode != resp.StatusCode {
				body, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(body))
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)
			mockRegRepo.AssertExpectations(t)
		})
	}
}

func TestHumaValidation_InvalidAcceptLanguage(t *testing.T) {
	t.Parallel()

	invalidLangs := []struct {
		name string
		lang string
	}{
		{name: "unsupported locale fr-FR", lang: "fr-FR"},
		{name: "lowercase en-us", lang: "en-us"},
		{name: "random string", lang: "invalid"},
	}

	type endpoint struct {
		name   string
		method string
		path   string
		body   []byte
	}

	minimalCreateBody, _ := json.Marshal(map[string]interface{}{
		"child_id":            "30000000-0000-0000-0000-000000000001",
		"guardian_id":         "11111111-1111-1111-1111-111111111111",
		"event_occurrence_id": "70000000-0000-0000-0000-000000000001",
		"status":              "registered",
	})
	minimalUpdateBody, _ := json.Marshal(map[string]interface{}{"status": "cancelled"})
	minimalPaymentStatusBody, _ := json.Marshal(map[string]interface{}{"payment_intent_status": "succeeded"})

	endpoints := []endpoint{
		{name: "GetRegistrationByID", method: http.MethodGet, path: "/api/v1/registrations/80000000-0000-0000-0000-000000000001"},
		{name: "GetRegistrationsByChildID", method: http.MethodGet, path: "/api/v1/registrations/child/30000000-0000-0000-0000-000000000001"},
		{name: "GetRegistrationsByGuardianID", method: http.MethodGet, path: "/api/v1/registrations/guardian/11111111-1111-1111-1111-111111111111"},
		{name: "GetRegistrationsByEventOccurrenceID", method: http.MethodGet, path: "/api/v1/registrations/event_occurrence/70000000-0000-0000-0000-000000000001"},
		{name: "CreateRegistration", method: http.MethodPost, path: "/api/v1/registrations", body: minimalCreateBody},
		{name: "UpdateRegistration", method: http.MethodPatch, path: "/api/v1/registrations/80000000-0000-0000-0000-000000000001", body: minimalUpdateBody},
		{name: "CancelRegistration", method: http.MethodPost, path: "/api/v1/registrations/80000000-0000-0000-0000-000000000001/cancel"},
		{name: "UpdateRegistrationPaymentStatus", method: http.MethodPatch, path: "/api/v1/registrations/80000000-0000-0000-0000-000000000001/payment-status", body: minimalPaymentStatusBody},
	}

	for _, ep := range endpoints {
		ep := ep
		for _, tt := range invalidLangs {
			tt := tt
			t.Run(ep.name+"/"+tt.name, func(t *testing.T) {
				t.Parallel()

				app, _ := setupRegistrationTestAPI(
					new(repomocks.MockRegistrationRepository),
					new(repomocks.MockChildRepository),
					new(repomocks.MockGuardianRepository),
					new(repomocks.MockEventOccurrenceRepository),
					new(repomocks.MockOrganizationRepository),
					new(stripemocks.MockStripeClient),
				)

				var reqBody io.Reader
				if ep.body != nil {
					reqBody = bytes.NewBuffer(ep.body)
				}

				req, err := http.NewRequest(ep.method, ep.path, reqBody)
				assert.NoError(t, err)
				if ep.body != nil {
					req.Header.Set("Content-Type", "application/json")
				}
				req.Header.Set("Accept-Language", tt.lang)

				resp, err := app.Test(req)
				assert.NoError(t, err)
				defer func() { _ = resp.Body.Close() }()

				assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode,
					"expected 422 for invalid Accept-Language %q on %s %s", tt.lang, ep.method, ep.path)
			})
		}
	}
}
