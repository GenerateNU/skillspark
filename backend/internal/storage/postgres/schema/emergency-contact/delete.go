package emergencycontact

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *EmergencyContactRepository) DeleteEmergencyContact(ctx context.Context, id uuid.UUID) (*models.DeleteEmergencyContactOutput, error) {
	query, err := schema.ReadSQLBaseScript("delete.sql", SqlSavedFiles)

	emergencyContact, err := r.GetEmergencyContactByID(ctx, id)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	commandTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		err := errs.InternalServerError("Failed to delete saved: ", err.Error())
		return nil, &err
	}

	if commandTag.RowsAffected() == 0 {
		err := errs.NotFound("Saved", "id", id)
		return nil, &err
	}

	output := &models.DeleteEmergencyContactOutput{}
	output.Body = emergencyContact

	return output, nil
}
