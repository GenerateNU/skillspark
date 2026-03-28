package emergencycontact

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) DeleteEmergencyContact(ctx context.Context, id uuid.UUID) (*models.DeleteEmergencyContactOutput, error) {
	id, err := uuid.Parse(id.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	emergencyContact, httpErr := h.EmergencyContactRepository.DeleteEmergencyContact(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}
	return emergencyContact, nil
}
