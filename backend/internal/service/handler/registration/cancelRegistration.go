package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
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
		refundInput := &models.CancelPaymentIntentInput{
			PaymentIntentID: registration.Body.StripePaymentIntentID,
			StripeAccountID: registration.Body.OrgStripeAccountID,
		}
		
		refundOutput, err := h.StripeClient.CancelPaymentIntent(ctx, refundInput)
		if err != nil {
			return nil, errs.InternalServerError("Failed to refund payment: ", err.Error())
		}
		refundStatus = refundOutput.Body.Status
	case "requires_capture":
		cancelInput := &models.CancelPaymentIntentInput{
			PaymentIntentID: registration.Body.StripePaymentIntentID,
			StripeAccountID: registration.Body.OrgStripeAccountID,
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