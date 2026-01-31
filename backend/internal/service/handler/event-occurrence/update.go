package eventoccurrence

import (
	"cmp"
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) UpdateEventOccurrence(ctx context.Context, input *models.UpdateEventOccurrenceInput) (*models.EventOccurrence, error) {
	// check that event occurrence exists already in database
	ogEventOccurrence, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	// check foreign keys
	var managerErr error
	var eventErr error
	var locationErr error

	managerId := input.Body.ManagerId
	if managerId != nil {
		_, managerErr = h.ManagerRepository.GetManagerByID(ctx, *managerId)
	}

	eventId := input.Body.EventId
	if eventId != nil {
		_, eventErr = h.EventRepository.GetEventByID(ctx, *eventId)
	}

	locationId := input.Body.LocationId
	if locationId != nil {
		_, locationErr = h.LocationRepository.GetLocationByID(ctx, *locationId)
	}

	if managerErr != nil || eventErr != nil || locationErr != nil {
		return nil, cmp.Or(managerErr, eventErr, locationErr)
	}

	// check that new currently enrolled number does not exceed the new or old max attendees
	newCurrEnrolled := input.Body.CurrEnrolled
	newMaxAttendees := input.Body.MaxAttendees
	ogMaxAttendees := ogEventOccurrence.MaxAttendees // cannot be null
	if newCurrEnrolled != nil {
		if (newMaxAttendees != nil) && (*newCurrEnrolled > *newMaxAttendees) {
			return nil, errs.BadRequest("Current enrolled cannot exceed max attendees")
		} else {
			if *newCurrEnrolled > ogMaxAttendees {
				return nil, errs.BadRequest("Current enrolled cannot exceed max attendees")
			}
		}
	}

	eventOccurrence, err := h.EventOccurrenceRepository.UpdateEventOccurrence(ctx, input)
	if err != nil {
		return nil, err
	}
	return eventOccurrence, nil
}
