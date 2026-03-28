package emergencycontact

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *EmergencyContactRepository) UpdateEmergencyContact(ctx context.Context, input *models.UpdateEmergencyContactInput) (*models.UpdateEmergencyContactOutput, error) {
	query, err := schema.ReadSQLBaseScript("update.sql", SqlSavedFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, input.Body.Name, input.Body.GuardianID, input.Body.PhoneNumber)

	var updatedEmergencyContact models.EmergencyContact

	err = row.Scan(
		&updatedEmergencyContact.ID,
		&updatedEmergencyContact.Name,
		&updatedEmergencyContact.GuardianID,
		&updatedEmergencyContact.PhoneNumber,
		&updatedEmergencyContact.CreatedAt,
		&updatedEmergencyContact.UpdatedAt,
	)

	if err != nil {
		err := errs.InternalServerError("Failed to create emergency contact: ", err.Error())
		return nil, &err
	}

	output := &models.UpdateEmergencyContactOutput{}
	output.Body = &updatedEmergencyContact

	return output, nil
}
