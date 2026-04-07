package eventoccurrence

import (
	"cmp"
	"context"
	"skillspark/internal/models"
)

func (h *Handler) CreateEventOccurrence(ctx context.Context, input *models.CreateEventOccurrenceInput) (*models.EventOccurrence, error) {

	// check that foreign keys exist in the database
	managerId := input.Body.ManagerId
	eventId := input.Body.EventId

	var managerErr error
	if managerId != nil {
		_, managerErr = h.ManagerRepository.GetManagerByID(ctx, *managerId)
	}

	_, eventErr := h.EventRepository.GetEventByID(ctx, eventId, input.AcceptLanguage)

	if managerErr != nil || eventErr != nil {
		return nil, cmp.Or(managerErr, eventErr)
	}

	eventOccurrence, err := h.EventOccurrenceRepository.CreateEventOccurrence(ctx, input)
	if err != nil {
		return nil, err
	}
	return eventOccurrence, nil
}
