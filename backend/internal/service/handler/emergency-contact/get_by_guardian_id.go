package emergencycontact

import (
	"context"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetEmergencyContactByGuardianID(ctx context.Context, guardian_id uuid.UUID) (*models.GetEmergencyContactByGuardianIDOutput, error) {
	return nil, nil
}
