package registration

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) GetRegistrationByPaymentIntentID(ctx context.Context, input *models.GetRegistrationByPaymentIntentIDInput) (*models.GetRegistrationByIDOutput, error) {
	registration, err := h.RegistrationRepository.GetRegistrationByPaymentIntentID(ctx, input.PaymentIntentID)
	if err != nil {
		return nil, err
	}

	return &models.GetRegistrationByIDOutput{
		Body: *registration,
	}, nil
}