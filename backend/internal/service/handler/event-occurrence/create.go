package eventoccurrence

import (
	"cmp"
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateEventOccurrence(ctx context.Context, input *models.CreateEventOccurrenceInput) (*models.EventOccurrence, error) {

	if input.AcceptLanguage != "en-US" && input.AcceptLanguage != "th-TH" {
		e := errs.BadRequest("Invalid AcceptLanguage parameter: language does not exist")
		return nil, &e
	}

	// check that foreign keys exist in the database
	managerId := input.Body.ManagerId
	eventId := input.Body.EventId
	locationId := input.Body.LocationId

	var managerErr error
	if managerId != nil {
		_, managerErr = h.ManagerRepository.GetManagerByID(ctx, *managerId)
	}

	_, eventErr := h.EventRepository.GetEventByID(ctx, eventId, input.AcceptLanguage)
	_, locationErr := h.LocationRepository.GetLocationByID(ctx, locationId)

	if managerErr != nil || eventErr != nil || locationErr != nil {
		return nil, cmp.Or(managerErr, eventErr, locationErr)
	}

	eventOccurrence, err := h.EventOccurrenceRepository.CreateEventOccurrence(ctx, input)
	if err != nil {
		return nil, err
	}
	return eventOccurrence, nil
}
