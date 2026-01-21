package createorganization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Execute(ctx context.Context, db *pgxpool.Pool, org *models.Organization) *errs.HTTPError {
	query, err := schema.ReadSQLBaseScript("organization/createorganization/baseQuery.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return &errr
	}

	_, err = db.Exec(ctx, query,
		org.ID,
		org.Name,
		org.Active,
		org.PfpS3Key,
		org.LocationID,
		org.CreatedAt,
		org.UpdatedAt,
	)

	if err != nil {
		errr := errs.InternalServerError("Failed to create organization: ", err.Error())
		return &errr
	}

	return nil
}