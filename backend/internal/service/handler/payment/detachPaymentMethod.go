package payment

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) DetachGuardianPaymentMethod(ctx context.Context, input *models.DetachPaymentMethodInput) (*struct{}, error) {

	pms, err := h.GetPaymentMethodsByGuardianID(ctx, &models.GetPaymentMethodsByGuardianIDInput{GuardianID: input.Body.GuardianID})
	if err != nil {
		return nil, err
	}

	for _, pm := range pms.Body.PaymentMethods {
		if pm.ID == input.Body.PaymentMethodID {
			err := h.StripeClient.DetachPaymentMethod(ctx, input.Body.PaymentMethodID)
			if err != nil {
				return nil, err
			}
			return &struct{}{}, nil
		}
	}

	return nil, errs.NotFound("PaymentMethod", "id", input.Body.PaymentMethodID)
}
