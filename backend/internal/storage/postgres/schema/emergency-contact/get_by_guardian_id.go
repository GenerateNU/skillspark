package emergencycontact

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *EmergencyContactRepository) GetEmergencyContactByGuardianID(ctx context.Context, guardian_id uuid.UUID) ([]*models.EmergencyContact, error) {
	query, err := schema.ReadSQLBaseScript("get_by_guardian_id.sql", SqlSavedFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, err
	}

	rows, err := r.db.Query(
		ctx,
		query,
		guardian_id,
	)

	if err != nil {
		err := errs.InternalServerError("Failed to fetch all emergency contacts: ", err.Error())
		return nil, &err
	}
	defer rows.Close()

	emergencyContacts, err := pgx.CollectRows(
		rows,
		func(row pgx.CollectableRow) (*models.EmergencyContact, error) {
			var emergencyContact models.EmergencyContact
			err := row.Scan(
				&emergencyContact.ID,
				&emergencyContact.Name,
				&emergencyContact.GuardianID,
				&emergencyContact.PhoneNumber,
				&emergencyContact.CreatedAt,
				&emergencyContact.UpdatedAt,
			)
			return &emergencyContact, err
		},
	)

	if err != nil {
		err := errs.InternalServerError("Failed to scan emergency contact: ", err.Error())
		return nil, &err
	}

	if err := rows.Err(); err != nil {
		err := errs.InternalServerError("Rows iteration error: ", err.Error())
		return nil, &err
	}

	return emergencyContacts, nil
}
