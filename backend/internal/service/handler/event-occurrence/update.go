package eventoccurrence

import (
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
	// 
	
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