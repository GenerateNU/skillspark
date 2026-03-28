package emergencycontact

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetEmergencyContactByGuardianID(ctx context.Context, guardian_id uuid.UUID) ([]*models.EmergencyContact, error) {
	guardian_id, err := uuid.Parse(guardian_id.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	emergencyContacts, httpErr := h.EmergencyContactRepository.GetEmergencyContactByGuardianID(ctx, guardian_id)
	if httpErr != nil {
		return nil, httpErr
	}
	return emergencyContacts, nil
}
