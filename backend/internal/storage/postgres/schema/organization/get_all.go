package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *OrganizationRepository) GetAllOrganizations(ctx context.Context, offset, pageSize int) ([]models.Organization, int, *errs.HTTPError) {
	
	baseQuery, err := schema.ReadSQLBaseScript("organization/sql/get_all.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, 0, &errr
	}

	countQuery, err := schema.ReadSQLBaseScript("organization/sql/count_all.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read count query: ", err.Error())
		return nil, 0, &errr
	}

	var totalCount int
	err = r.db.QueryRow(ctx, countQuery).Scan(&totalCount)
	if err != nil {
		errr := errs.InternalServerError("Failed to count organizations: ", err.Error())
		return nil, 0, &errr
	}

	rows, err := r.db.Query(ctx, baseQuery, pageSize, offset)
	if err != nil {
		errr := errs.InternalServerError("Failed to get organizations: ", err.Error())
		return nil, 0, &errr
	}
	defer rows.Close()

	orgs, err := pgx.CollectRows(rows, scanOrganization)
	if err != nil {
		errr := errs.InternalServerError("Failed to collect organizations: ", err.Error())
		return nil, 0, &errr
	}

	return orgs, totalCount, nil
}

func scanOrganization(row pgx.CollectableRow) (models.Organization, error) {
	var org models.Organization
	err := row.Scan(
		&org.ID,
		&org.Name,
		&org.Active,
		&org.PfpS3Key,
		&org.LocationID,
		&org.CreatedAt,
		&org.UpdatedAt,
	)
	return org, err
}