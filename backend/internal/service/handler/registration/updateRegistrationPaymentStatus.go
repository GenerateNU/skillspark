package registration

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) UpdateRegistrationPaymentStatus(ctx context.Context, input *models.UpdateRegistrationPaymentStatusInput) (*models.UpdateRegistrationPaymentStatusOutput, error) {
	updated, err := h.RegistrationRepository.UpdateRegistrationPaymentStatus(ctx, input)
	if err != nil {
		return nil, err
	}

	return updated, nil
}