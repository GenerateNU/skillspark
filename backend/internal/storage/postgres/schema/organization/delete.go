package organization

import (
	"context"
	"github.com/google/uuid"
	"skillspark/internal/errs"
	"skillspark/internal/storage/postgres/schema"
)

func (r *OrganizationRepository) DeleteOrganization(ctx context.Context, id uuid.UUID) *errs.HTTPError {
	query, err := schema.ReadSQLBaseScript("organization/sql/delete.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return &errr
	}

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		errr := errs.InternalServerError("Failed to delete organization: ", err.Error())
		return &errr
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		errr := errs.NotFound("Organization", "id", id)
		return &errr
	}

	return nil
}
