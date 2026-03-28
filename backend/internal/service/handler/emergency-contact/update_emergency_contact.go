package emergencycontact

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) UpdateEmergencyContact(ctx context.Context, input *models.UpdateEmergencyContactInput) (*models.UpdateEmergencyContactOutput, error) {
	updatedEmergencyContact, err := h.UpdateEmergencyContact(ctx, input)
	if err != nil {
		return nil, err
	}

	return updatedEmergencyContact, nil
}
