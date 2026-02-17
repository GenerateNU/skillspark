package child

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *ChildRepository) UpdateChildByID(ctx context.Context, childID uuid.UUID, child *models.UpdateChildInput) (*models.Child, error) {

	query, err := schema.ReadSQLBaseScript("update_by_id.sql", SqlChildFiles)
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
		childID)

	var updatedChild models.Child
	err = row.Scan(
		&updatedChild.ID,
		&updatedChild.Name,
		&updatedChild.SchoolID,
		&updatedChild.SchoolName,
		&updatedChild.BirthMonth,
		&updatedChild.BirthYear,
		&updatedChild.Interests,
		&updatedChild.GuardianID,
		&updatedChild.CreatedAt,
		&updatedChild.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Child", "id", childID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch child by id: ", err.Error())
		return nil, &err
	}

	return &updatedChild, nil
}
