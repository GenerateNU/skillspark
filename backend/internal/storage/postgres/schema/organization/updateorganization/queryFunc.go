package updateorganization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Execute(ctx context.Context, db *pgxpool.Pool, org *models.Organization) *errs.HTTPError {
	query, err := schema.ReadSQLBaseScript("organization/updateorganization/baseQuery.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return &errr
	}

	result, err := db.Exec(ctx, query,
		org.Name,
		org.Active,
		org.PfpS3Key,
		org.LocationID,
		org.UpdatedAt,
		org.ID,
	)

	if err != nil {
		errr := errs.InternalServerError("Failed to update organization: ", err.Error())
		return &errr
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		errr := errs.NotFound("Organization", "id", org.ID)
		return &errr
	}

	return nil
}