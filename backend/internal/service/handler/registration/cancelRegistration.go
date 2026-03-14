package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"time"
)

func (h *Handler) CancelRegistration(ctx context.Context, input *models.CancelRegistrationInput) (*models.CancelRegistrationOutput, error) {

	getInput := &models.GetRegistrationByIDInput{
		ID: input.ID,
	}

	registration, err := h.RegistrationRepository.GetRegistrationByID(ctx, getInput, nil)
	if err != nil {
		return nil, err
	}

	if registration.Body.Status == models.RegistrationStatusCancelled {
		return nil, errs.BadRequest("Registration is already cancelled")
	}

	var refundStatus string
	switch registration.Body.PaymentIntentStatus {
	case "succeeded":
		eventoccurrence, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, registration.Body.EventOccurrenceID, "en-US")
		if err != nil {
			return nil, err
		}
		if eventoccurrence.StartTime.Before(time.Now().AddDate(0, 0, 1)) {
			refundStatus = "no_refund_needed"
		} else {
			refundInput := &models.RefundPaymentInput{
				PaymentIntentID: registration.Body.StripePaymentIntentID,
			}
			refundOutput, err := h.StripeClient.RefundPayment(ctx, refundInput)
			if err != nil {
				return nil, errs.InternalServerError("Failed to refund payment: ", err.Error())
			}
			refundStatus = refundOutput.Body.Status
		}

	case "requires_capture":
		cancelInput := &models.CancelPaymentIntentInput{
			PaymentIntentID: registration.Body.StripePaymentIntentID,
		}

		_, err := h.StripeClient.CancelPaymentIntent(ctx, cancelInput)
		if err != nil {
			return nil, errs.InternalServerError("Failed to cancel payment intent: ", err.Error())
		}
		refundStatus = "cancelled"
	default:
		refundStatus = "no_refund_needed"
	}

	cancelledRegistration, err := h.RegistrationRepository.CancelRegistration(ctx, input)
	if err != nil {
		return nil, err
	}

	cancelledRegistration.Body.Message = "Registration cancelled successfully"
	cancelledRegistration.Body.RefundStatus = refundStatus

	return cancelledRegistration, nil
}
