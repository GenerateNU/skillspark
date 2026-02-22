package organization

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *OrganizationRepository) GetOrganizationByID(ctx context.Context, id uuid.UUID) (*models.Organization, error) {
	query, err := schema.ReadSQLBaseScript("get_by_id.sql", SqlOrganizationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, id)
	var org models.Organization
	err = row.Scan(
		&org.ID,
		&org.Name,
		&org.Active,
		&org.PfpS3Key,
		&org.LocationID,
		&org.StripeAccountID,
		&org.StripeAccountActivated,
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
