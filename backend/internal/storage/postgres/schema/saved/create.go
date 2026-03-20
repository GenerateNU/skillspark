package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

const language = "en-US"

func (r *SavedRepository) CreateSaved(ctx context.Context, saved *models.CreateSavedInput) (*models.Saved, error) {

	query, err := schema.ReadSQLBaseScript("create.sql", SqlSavedFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, saved.Body.GuardianID, saved.Body.EventID)

	var createdSaved models.Saved
	var titleEN, descriptionEN string
	var titleTH, descriptionTH *string

	err = row.Scan(
		&createdSaved.ID,
		&createdSaved.GuardianID,
		&createdSaved.CreatedAt,
		&createdSaved.UpdatedAt,
		&createdSaved.Event.ID,
		&titleEN,
		&titleTH,
		&descriptionEN,
		&descriptionTH,
		&createdSaved.Event.OrganizationID,
		&createdSaved.Event.AgeRangeMin,
		&createdSaved.Event.AgeRangeMax,
		&createdSaved.Event.Category,
		&createdSaved.Event.HeaderImageS3Key,
		&createdSaved.Event.CreatedAt,
		&createdSaved.Event.UpdatedAt,
	)

	if err != nil {
		err := errs.InternalServerError("Failed to create saved event: ", err.Error())
		return nil, &err
	}

	switch language {
	case "th-TH":
		if titleTH != nil {
			createdSaved.Event.Title = *titleTH
		}
		if descriptionTH != nil {
			createdSaved.Event.Description = *descriptionTH
		}
	case "en-US":
		createdSaved.Event.Title = titleEN
		createdSaved.Event.Description = descriptionEN
	}

	return &createdSaved, nil

}
