package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/google/uuid"
)

func (r *SavedRepository) GetByGuardianID(ctx context.Context, user_id uuid.UUID, pagination utils.Pagination, AcceptLanguage string) ([]models.Saved, error) {

	query, err := schema.ReadSQLBaseScript("get_all_for_user.sql", SqlSavedFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, err
	}

	rows, err := r.db.Query(
		ctx,
		query,
		user_id,
		pagination.Limit,
		pagination.GetOffset(),
	)

	if err != nil {
		err := errs.InternalServerError("Failed to fetch all saved: ", err.Error())
		return nil, &err
	}
	defer rows.Close()

	saved := []models.Saved{}

	for rows.Next() {
		var s models.Saved

		var titleEN, descriptionEN string
		var titleTH, descriptionTH *string

		err := rows.Scan(
			&s.ID,
			&s.GuardianID,
			&s.CreatedAt,
			&s.UpdatedAt,

			&s.Event.ID,
			&titleEN,
			&titleTH,
			&descriptionEN,
			&descriptionTH,
			&s.Event.OrganizationID,
			&s.Event.AgeRangeMin,
			&s.Event.AgeRangeMax,
			&s.Event.Category,
			&s.Event.HeaderImageS3Key,
			&s.Event.CreatedAt,
			&s.Event.UpdatedAt,
		)
		if err != nil {
			err := errs.InternalServerError("Failed to scan all event occurrences: ", err.Error())
			return nil, &err
		}

		switch AcceptLanguage {
		case "th-TH":
			if titleTH != nil {
				s.Event.Title = *titleTH
			}
			if descriptionTH != nil {
				s.Event.Description = *descriptionTH
			}
		case "en-US":
			s.Event.Title = titleEN
			s.Event.Description = descriptionEN
		}

		saved = append(saved, s)
	}

	if err := rows.Err(); err != nil {
		err := errs.InternalServerError("Rows iteration error: ", err.Error())
		return nil, &err
	}

	return saved, nil
}
