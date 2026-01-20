package child

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *ChildRepository) UpdateChildByID(ctx context.Context, child *models.UpdateChildInput) (*models.Child, *errs.HTTPError) {

	query, err := schema.ReadSQLBaseScript("child/sql/update_by_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(
		ctx,
		query,
		child.Body.Name,
		child.Body.SchoolID,
		child.Body.BirthMonth,
		child.Body.BirthYear,
		child.Body.Interests,
		child.Body.GuardianID,
		child.Body.ID)

	var updatedChild models.Child
	err = row.Scan(&updatedChild.ID, &updatedChild.SchoolID, &updatedChild.SchoolName, &updatedChild.BirthMonth, &updatedChild.BirthYear, &updatedChild.Interests, &updatedChild.GuardianID, &updatedChild.CreatedAt, &updatedChild.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Child", "id", child.Body.ID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch child by id: ", err.Error())
		return nil, &err
	}

	return &updatedChild, nil
}
