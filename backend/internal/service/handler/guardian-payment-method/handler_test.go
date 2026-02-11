package guardianpaymentmethod

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	stripemocks "skillspark/internal/stripeClient/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetPaymentMethodsByGuardianID(t *testing.T) {
	tests := []struct {
		name      string
		guardianID string
		mockSetup func(*repomocks.MockGuardianRepository, *repomocks.MockGuardianPaymentMethodRepository)
		wantErr   bool
		wantCount int
	}{
		{
			name:       "successfully retrieves payment methods",
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
			wantErr:   false,
			wantCount: 1,
		},
		{
			name:       "guardian has no stripe customer - returns empty",
			guardianID: "88888888-8888-8888-8888-888888888889",
			mockSetup: func(gr *repomocks.MockGuardianRepository, pmr *repomocks.MockGuardianPaymentMethodRepository) {
				gr.On("GetGuardianByID", mock.Anything, uuid.MustParse("88888888-8888-8888-8888-888888888889")).
					Return(&models.Guardian{
						ID:               uuid.MustParse("88888888-8888-8888-8888-888888888889"),
						StripeCustomerID: nil,
					}, nil)
			},
			wantErr:   false,
			wantCount: 0,
		},
		{
			name:       "guardian not found",
			guardianID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(gr *repomocks.MockGuardianRepository, pmr *repomocks.MockGuardianPaymentMethodRepository) {
				gr.On("GetGuardianByID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
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

			handler := NewHandler(mockGuardianRepo, mockPaymentMethodRepo, mockStripeClient)
			ctx := context.Background()

			input := &models.GetGuardianPaymentMethodsByGuardianIDInput{
				GuardianID: uuid.MustParse(tt.guardianID),
			}
			output, err := handler.GetPaymentMethodsByGuardianID(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				require.NotNil(t, output)
				assert.Equal(t, tt.wantCount, len(output.Body))
			}

			mockGuardianRepo.AssertExpectations(t)
			mockPaymentMethodRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateGuardianPaymentMethod(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateGuardianPaymentMethodInput
		mockSetup func(*repomocks.MockGuardianRepository, *repomocks.MockGuardianPaymentMethodRepository)
		wantErr   bool
	}{
		{
			name: "successfully creates payment method",
			input: func() *models.CreateGuardianPaymentMethodInput {
				input := &models.CreateGuardianPaymentMethodInput{}
				input.Body.GuardianID = uuid.MustParse("88888888-8888-8888-8888-888888888888")
				input.Body.IsDefault = true
				return input
			}(),
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
			wantErr: false,
		},
		{
			name: "fails when guardian has no stripe customer",
			input: func() *models.CreateGuardianPaymentMethodInput {
				input := &models.CreateGuardianPaymentMethodInput{}
				input.Body.GuardianID = uuid.MustParse("88888888-8888-8888-8888-888888888889")
				input.Body.IsDefault = false
				return input
			}(),
			mockSetup: func(gr *repomocks.MockGuardianRepository, pmr *repomocks.MockGuardianPaymentMethodRepository) {
				gr.On("GetGuardianByID", mock.Anything, uuid.MustParse("88888888-8888-8888-8888-888888888889")).
					Return(&models.Guardian{
						ID:               uuid.MustParse("88888888-8888-8888-8888-888888888889"),
						StripeCustomerID: nil,
					}, nil)
			},
			wantErr: true,
		},
		{
			name: "fails when guardian not found",
			input: func() *models.CreateGuardianPaymentMethodInput {
				input := &models.CreateGuardianPaymentMethodInput{}
				input.Body.GuardianID = uuid.MustParse("00000000-0000-0000-0000-000000000000")
				input.Body.IsDefault = false
				return input
			}(),
			mockSetup: func(gr *repomocks.MockGuardianRepository, pmr *repomocks.MockGuardianPaymentMethodRepository) {
				gr.On("GetGuardianByID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
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

			handler := NewHandler(mockGuardianRepo, mockPaymentMethodRepo, mockStripeClient)
			ctx := context.Background()

			paymentMethod, err := handler.CreateGuardianPaymentMethod(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, paymentMethod)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, paymentMethod)
			}

			mockGuardianRepo.AssertExpectations(t)
			mockPaymentMethodRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_DeleteGuardianPaymentMethod(t *testing.T) {
	tests := []struct {
		name      string
		pmID      string
		mockSetup func(*repomocks.MockGuardianPaymentMethodRepository, *stripemocks.MockStripeClient)
		wantErr   bool
	}{
		{
			name: "successfully deletes payment method",
			pmID: "11111111-1111-1111-1111-111111111111",
			mockSetup: func(pmr *repomocks.MockGuardianPaymentMethodRepository, sc *stripemocks.MockStripeClient) {
				pmr.On("DeleteGuardianPaymentMethod", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return(&models.GuardianPaymentMethod{
						ID:                    uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						StripePaymentMethodID: "pm_test123",
						IsDefault:             false,
					}, nil)

				sc.On("DetachPaymentMethod", mock.Anything, "pm_test123").Return(nil)
			},
			wantErr: false,
		},
		{
			name: "payment method not found",
			pmID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(pmr *repomocks.MockGuardianPaymentMethodRepository, sc *stripemocks.MockStripeClient) {
				pmr.On("DeleteGuardianPaymentMethod", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Guardian Payment Method", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
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

			handler := NewHandler(mockGuardianRepo, mockPaymentMethodRepo, mockStripeClient)
			ctx := context.Background()

			input := &models.DeleteGuardianPaymentMethodInput{
				ID: uuid.MustParse(tt.pmID),
			}
			output, err := handler.DeleteGuardianPaymentMethod(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, output)
				assert.Equal(t, "Payment method deleted successfully", output.Body.Message)
			}

			mockPaymentMethodRepo.AssertExpectations(t)
			mockStripeClient.AssertExpectations(t)
		})
	}
}

func TestHandler_UpdateGuardianPaymentMethod(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.SetDefaultPaymentMethodInput
		mockSetup func(*repomocks.MockGuardianRepository, *repomocks.MockGuardianPaymentMethodRepository)
		wantErr   bool
	}{
		{
			name: "successfully sets payment method as default",
			input: &models.SetDefaultPaymentMethodInput{
				GuardianID:      uuid.MustParse("88888888-8888-8888-8888-888888888888"),
				PaymentMethodID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			},
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
			wantErr: false,
		},
		{
			name: "fails when guardian has no stripe customer",
			input: &models.SetDefaultPaymentMethodInput{
				GuardianID:      uuid.MustParse("88888888-8888-8888-8888-888888888889"),
				PaymentMethodID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			},
			mockSetup: func(gr *repomocks.MockGuardianRepository, pmr *repomocks.MockGuardianPaymentMethodRepository) {
				gr.On("GetGuardianByID", mock.Anything, uuid.MustParse("88888888-8888-8888-8888-888888888889")).
					Return(&models.Guardian{
						ID:               uuid.MustParse("88888888-8888-8888-8888-888888888889"),
						StripeCustomerID: nil,
					}, nil)
			},
			wantErr: true,
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

			handler := NewHandler(mockGuardianRepo, mockPaymentMethodRepo, mockStripeClient)
			ctx := context.Background()

			output, err := handler.SetDefaultGuardianPaymentMethod(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, output)
			} else {
				assert.NoError(t, err)
				require.NotNil(t, output)
				assert.True(t, output.Body.IsDefault)
			}

			mockGuardianRepo.AssertExpectations(t)
			mockPaymentMethodRepo.AssertExpectations(t)
		})
	}
}