package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) UpdateRegistrationPaymentStatus(ctx context.Context, input *models.UpdateRegistrationPaymentStatusInput) (*models.UpdateRegistrationPaymentStatusOutput, error) {

	if input.AcceptLanguage != "en-US" && input.AcceptLanguage != "th-TH" {
		e := errs.BadRequest("Invalid AcceptLanguage parameter: language does not exist")
		return nil, &e
	}

	updated, err := h.RegistrationRepository.UpdateRegistrationPaymentStatus(ctx, input)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
