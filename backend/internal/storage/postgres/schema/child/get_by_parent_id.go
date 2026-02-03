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

func (r *ChildRepository) GetChildrenByParentID(
	ctx context.Context,
	parentID uuid.UUID,
) ([]models.Child, error) {

	// if the guardian does not exist, want to error
	existsQuery, err := schema.ReadSQLBaseScript("child/sql/guardian_exists.sql")
	if err != nil {
		httpErr := errs.InternalServerError("Failed to read guardian exists query")
		return nil, &httpErr
	}

	var exists int
	err = r.db.QueryRow(ctx, existsQuery, parentID).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			httpErr := errs.NotFound("Guardian", "id", parentID)
			return nil, &httpErr
		}
		httpErr := errs.InternalServerError("Failed to check guardian existence")
		return nil, &httpErr
	}

	query, err := schema.ReadSQLBaseScript("child/sql/get_by_parent_id.sql")
	if err != nil {
		httpErr := errs.InternalServerError("Failed to read base query")
		return nil, &httpErr
	}

	rows, err := r.db.Query(ctx, query, parentID)
	if err != nil {
		httpErr := errs.InternalServerError("Failed to fetch children")
		return nil, &httpErr
	}
	defer rows.Close()

	children, err := pgx.CollectRows(
		rows,
		func(row pgx.CollectableRow) (models.Child, error) {
			var child models.Child
			err := row.Scan(
				&child.ID,
				&child.Name,
				&child.SchoolID,
				&child.SchoolName,
				&child.BirthMonth,
				&child.BirthYear,
				&child.Interests,
				&child.GuardianID,
				&child.CreatedAt,
				&child.UpdatedAt,
			)
			return child, err
		},
	)

	if err != nil {
		httpErr := errs.InternalServerError("Failed to collect children")
		return nil, &httpErr
	}

	return children, nil
}
