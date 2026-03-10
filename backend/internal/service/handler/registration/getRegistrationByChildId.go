package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetRegistrationsByChildID(ctx context.Context, input *models.GetRegistrationsByChildIDInput) (*models.GetRegistrationsByChildIDOutput, error) {

	if input.AcceptLanguage != "en-US" && input.AcceptLanguage != "th-TH" {
		e := errs.BadRequest("Invalid AcceptLanguage parameter: language does not exist")
		return nil, &e
	}

	childID, err := uuid.Parse(input.ChildID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid child ID format")
	}

	registrations, httpErr := h.RegistrationRepository.GetRegistrationsByChildID(ctx, &models.GetRegistrationsByChildIDInput{ChildID: childID, AcceptLanguage: input.AcceptLanguage})
	if httpErr != nil {
		return nil, httpErr
	}

	return registrations, nil
}
