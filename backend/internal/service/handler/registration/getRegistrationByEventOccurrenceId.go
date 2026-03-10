package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetRegistrationsByEventOccurrenceID(ctx context.Context, input *models.GetRegistrationsByEventOccurrenceIDInput) (*models.GetRegistrationsByEventOccurrenceIDOutput, error) {

	if input.AcceptLanguage != "en-US" && input.AcceptLanguage != "th-TH" {
		e := errs.BadRequest("Invalid AcceptLanguage parameter: language does not exist")
		return nil, &e
	}

	eventOccurrenceID, err := uuid.Parse(input.EventOccurrenceID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid child ID format")
	}

	registrations, httpErr := h.RegistrationRepository.GetRegistrationsByEventOccurrenceID(ctx, &models.GetRegistrationsByEventOccurrenceIDInput{EventOccurrenceID: eventOccurrenceID, AcceptLanguage: input.AcceptLanguage})
	if httpErr != nil {
		return nil, httpErr
	}

	return registrations, nil
}
