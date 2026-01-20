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

func (r *ChildRepository) DeleteChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, error) {

	query, err := schema.ReadSQLBaseScript("child/sql/delete_by_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, childID)

	var deletedChild models.Child
	err = row.Scan(&deletedChild.ID, &deletedChild.Name, &deletedChild.SchoolID, &deletedChild.SchoolName, &deletedChild.BirthMonth, &deletedChild.BirthYear, &deletedChild.Interests, &deletedChild.GuardianID, &deletedChild.CreatedAt, &deletedChild.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Child", "id", childID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to delete child by id: ", err.Error())
		return nil, &err
	}

	return &deletedChild, nil
}
