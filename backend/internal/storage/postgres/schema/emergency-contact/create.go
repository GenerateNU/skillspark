package emergencycontact

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *EmergencyContactRepository) CreateEmergencyContact(ctx context.Context, input *models.CreateEmergencyContactInput) (*models.CreateEmergencyContactOutput, error) {
	query, err := schema.ReadSQLBaseScript("create.sql", SqlSavedFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, input.Body.Name, input.Body.GuardianID, input.Body.PhoneNumber)

	var createdEmergencyContact models.EmergencyContact

	err = row.Scan(
		&createdEmergencyContact.ID,
		&createdEmergencyContact.Name,
		&createdEmergencyContact.GuardianID,
		&createdEmergencyContact.PhoneNumber,
		&createdEmergencyContact.CreatedAt,
		&createdEmergencyContact.UpdatedAt,
	)

	if err != nil {
		err := errs.InternalServerError("Failed to create emergency contact: ", err.Error())
		return nil, &err
	}

	output := &models.CreateEmergencyContactOutput{}
	output.Body = &createdEmergencyContact

	return output, nil

}
