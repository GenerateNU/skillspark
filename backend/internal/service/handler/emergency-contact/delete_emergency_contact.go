package emergencycontact

import (
	"context"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) DeleteEmergencyContact(ctx context.Context, id uuid.UUID) (*models.DeleteEmergencyContactOutput, error) {
	return nil, nil
}
