package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) GetRegistrationByPaymentIntentID(ctx context.Context, input *models.GetRegistrationByPaymentIntentIDInput) (*models.GetRegistrationByIDOutput, error) {

	if input.AcceptLanguage != "en-US" && input.AcceptLanguage != "th-TH" {
		e := errs.BadRequest("Invalid AcceptLanguage parameter: language does not exist")
		return nil, &e
	}

	registration, err := h.RegistrationRepository.GetRegistrationByPaymentIntentID(ctx, input.PaymentIntentID, input.AcceptLanguage)
	if err != nil {
		return nil, err
	}

	return &models.GetRegistrationByIDOutput{
		Body: *registration,
	}, nil
}
