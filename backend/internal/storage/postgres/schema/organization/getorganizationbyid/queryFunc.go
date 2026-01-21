package getorganizationbyid

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Execute(ctx context.Context, db *pgxpool.Pool, id uuid.UUID) (*models.Organization, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("organization/getorganizationbyid/baseQuery.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := db.QueryRow(ctx, query, id)
	var org models.Organization
	err = row.Scan(
		&org.ID,
		&org.Name,
		&org.Active,
		&org.PfpS3Key,
		&org.LocationID,
		&org.CreatedAt,
		&org.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Organization", "id", id)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch organization by id: ", err.Error())
		return nil, &err
	}

	return &org, nil
}