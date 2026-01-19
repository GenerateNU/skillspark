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

func (r *ChildRepository) GetChildrenByParentID(ctx context.Context, parentID uuid.UUID) ([]models.Child, *errs.HTTPError) {

	query, err := schema.ReadSQLBaseScript("child/sql/get_my_parent_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, query, parentID)

	children, err := pgx.CollectRows(
		rows,
		func(row pgx.CollectableRow) (models.Child, error) {
			var child models.Child
			err = row.Scan(&child.ID, &child.SchoolID, &child.SchoolName, &child.BirthMonth, &child.BirthYear, &child.Interests, &child.GuardianID, &child.CreatedAt, &child.UpdatedAt)
			return child, err
		},
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			httpErr := errs.NotFound("Child", "Parent has no children", parentID)
			return nil, &httpErr
		}
		err := errs.InternalServerError("Failed to fetch parent by id: ", err.Error())
		return nil, &err
	}
	defer rows.Close()

	return children, nil
}
