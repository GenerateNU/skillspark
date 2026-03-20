package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *SavedRepository) CreateSaved(ctx context.Context, saved *models.CreateSavedInput) (*models.Saved, error) {

	query, err := schema.ReadSQLBaseScript("create.sql", SqlSavedFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, saved.Body.GuardianID, saved.Body.EventOccurrenceID)

	var createdSaved models.Saved

	err = row.Scan(&createdSaved.ID, &createdSaved.GuardianID, &createdSaved.EventOccurrenceID, &createdSaved.CreatedAt, &createdSaved.UpdatedAt)

	if err != nil {
		err := errs.InternalServerError("Failed to create saved event: ", err.Error())
		return nil, &err
	}

	return &createdSaved, nil

}
