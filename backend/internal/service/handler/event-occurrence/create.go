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
	locationId := input.Body.LocationId

	_, managerErr := h.ManagerRepository.GetManagerByID(ctx, *managerId)
	_, eventErr := h.EventRepository.GetEventByID(ctx, eventId)
	_, locationErr := h.LocationRepository.GetLocationByID(ctx, locationId)
	
	if (managerId != nil && managerErr != nil) || eventErr != nil || locationErr != nil {
		return nil, cmp.Or(eventErr, locationErr, managerErr)
	}

	eventOccurrence, err := h.EventOccurrenceRepository.CreateEventOccurrence(ctx, input)
	if err != nil {
		return nil, err
	}
	return eventOccurrence, nil
}