package deleteorganization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Execute(ctx context.Context, db *pgxpool.Pool, id uuid.UUID) *errs.HTTPError {
	query, err := schema.ReadSQLBaseScript("organization/deleteorganization/baseQuery.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return &errr
	}

	result, err := db.Exec(ctx, query, id)
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