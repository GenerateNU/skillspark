package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetRegistrationByID(ctx context.Context, input *models.GetRegistrationByIDInput) (*models.GetRegistrationByIDOutput, error) {

	if input.AcceptLanguage != "en-US" && input.AcceptLanguage != "th-TH" {
		e := errs.BadRequest("Invalid AcceptLanguage parameter: language does not exist")
		return nil, &e
	}

	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	registration, httpErr := h.RegistrationRepository.GetRegistrationByID(ctx, &models.GetRegistrationByIDInput{ID: id, AcceptLanguage: input.AcceptLanguage}, nil)
	if httpErr != nil {
		return nil, httpErr
	}

	return registration, nil
}
