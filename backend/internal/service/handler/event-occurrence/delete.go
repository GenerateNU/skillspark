package eventoccurrence

import (
	"context"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) CancelEventOccurrence(ctx context.Context, id uuid.UUID) (string, error) {
	registrations, err := h.RegistrationRepository.GetRegistrationsByEventOccurrenceID(ctx, &models.GetRegistrationsByEventOccurrenceIDInput{
		EventOccurrenceID: id,
	})

	if err != nil {
		return "", err
	}

	for _, reg := range registrations.Body.Registrations {
		switch reg.PaymentIntentStatus {
		case "succeeded":
			refundInput := &models.RefundPaymentInput{
				PaymentIntentID: reg.StripePaymentIntentID,
			}
			if _, err := h.StripeClient.RefundPayment(ctx, refundInput); err != nil {
				return "", err
			}
		case "requires_capture":
			cancelInput := &models.CancelPaymentIntentInput{
				PaymentIntentID: reg.StripePaymentIntentID,
			}
			if _, err := h.StripeClient.CancelPaymentIntent(ctx, cancelInput); err != nil {
				return "", err
			}
		}
	}

	if err := h.EventOccurrenceRepository.CancelEventOccurrence(ctx, id); err != nil {
		return "", err
	}

	return "Event occurrence successfully cancelled.", nil
}
