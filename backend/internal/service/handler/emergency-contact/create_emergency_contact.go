package emergencycontact

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) CreateEmergencyContact(ctx context.Context, input *models.CreateEmergencyContactInput) (*models.CreateEmergencyContactOutput, error) {
	createdEmergencyContact, err := h.EmergencyContactRepository.CreateEmergencyContact(ctx, input)
	if err != nil {
		return nil, err
	}

	return createdEmergencyContact, nil
}
