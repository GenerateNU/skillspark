package emergencycontact

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *EmergencyContactRepository) GetEmergencyContactByID(ctx context.Context, id uuid.UUID) (*models.EmergencyContact, error) {
	query, err := schema.ReadSQLBaseScript("get_by_id.sql", SqlSavedFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	var ec models.EmergencyContact

	err = r.db.QueryRow(ctx, query, id).Scan(
		&ec.ID,
		&ec.Name,
		&ec.GuardianID,
		&ec.PhoneNumber,
		&ec.CreatedAt,
		&ec.UpdatedAt,
	)
	if err != nil {
		err := errs.NotFound("EmergencyContact", "id", id)
		return nil, &err
	}

	return &ec, nil
}
