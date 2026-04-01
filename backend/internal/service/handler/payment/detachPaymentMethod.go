package payment

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) DetachGuardianPaymentMethod(ctx context.Context, input *models.DetachPaymentMethodInput) (*struct{}, error) {
	err := h.StripeClient.DetachPaymentMethod(ctx, input.Body.PaymentMethodID)
	if err != nil {
		return nil, err
	}
	return &struct{}{}, nil
}
