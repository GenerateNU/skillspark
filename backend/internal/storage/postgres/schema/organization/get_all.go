package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/jackc/pgx/v5"
)

func (r *OrganizationRepository) GetAllOrganizations(ctx context.Context, pagination utils.Pagination) ([]models.Organization, error) {

	baseQuery, err := schema.ReadSQLBaseScript("organization/sql/get_all.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, baseQuery, pagination.Limit, pagination.GetOffset())
	if err != nil {
		errr := errs.InternalServerError("Failed to get organizations: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	orgs, err := pgx.CollectRows(rows, scanOrganization)
	if err != nil {
		errr := errs.InternalServerError("Failed to collect organizations: ", err.Error())
		return nil, &errr
	}

	return orgs, nil
}

func scanOrganization(row pgx.CollectableRow) (models.Organization, error) {
	var org models.Organization
	err := row.Scan(
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
	return org, err
}
