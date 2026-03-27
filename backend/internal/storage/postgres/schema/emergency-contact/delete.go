package emergencycontact

import (
	"context"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (r *EmergencyContactRepository) DeleteEmergencyContact(ctx context.Context, guardian_id uuid.UUID) (*models.DeleteEmergencyContactOutput, error) {
	return nil, nil
}
