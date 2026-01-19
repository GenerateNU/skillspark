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

func (r *ChildRepository) GetChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("child/sql/get_by_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, childID)
	var child models.Child
	err = row.Scan(&child.ID, &child.SchoolID, &child.SchoolName, &child.BirthMonth, &child.BirthYear, &child.Interests, &child.GuardianID, &child.CreatedAt, &child.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Child", "id", childID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch child by id: ", err.Error())
		return nil, &err
	}

	return &child, nil

}
