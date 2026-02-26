package eventoccurrence

import (
	"context"
	"log"
	"skillspark/internal/models"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) CancelEventOccurrence(ctx context.Context, id uuid.UUID) (string, error) {
	registrations, err := h.RegistrationRepository.GetRegistrationsByEventOccurrenceID(ctx, &models.GetRegistrationsByEventOccurrenceIDInput{
		EventOccurrenceID: id,
	})
	if err != nil {
		return "", err
	}

	eventOccurrence, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, id)
	if err != nil {
		return "", err
	}

	if err := h.EventOccurrenceRepository.CancelEventOccurrence(ctx, id); err != nil {
		return "", err
	}

	for _, reg := range registrations.Body.Registrations {
		switch reg.PaymentIntentStatus {
		case "succeeded":
			if eventOccurrence.StartTime.Before(time.Now().AddDate(0, 0, 1)) {
			} else {
				refundInput := &models.RefundPaymentInput{
					PaymentIntentID: reg.StripePaymentIntentID,
				}
				if _, err := h.StripeClient.RefundPayment(ctx, refundInput); err != nil {
					log.Printf("Failed to refund registration %s: %v", reg.ID, err)
				}
			}
		case "requires_capture":
			cancelInput := &models.CancelPaymentIntentInput{
				PaymentIntentID: reg.StripePaymentIntentID,
			}
			if _, err := h.StripeClient.CancelPaymentIntent(ctx, cancelInput); err != nil {
				log.Printf("Failed to cancel registration %s: %v", reg.ID, err)
			}
		}
	}

	return "Event occurrence successfully cancelled.", nil
}