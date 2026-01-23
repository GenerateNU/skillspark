package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *OrganizationRepository) UpdateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("organization/sql/update.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		org.Name,
		org.Active,
		org.PfpS3Key,
		org.LocationID,
		org.ID,
	)

	var updatedOrganization models.Organization

	err = row.Scan(
		&updatedOrganization.ID,
		&updatedOrganization.Name,
		&updatedOrganization.Active,
		&updatedOrganization.PfpS3Key,
		&updatedOrganization.LocationID,
		&updatedOrganization.CreatedAt,
		&updatedOrganization.UpdatedAt,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			errr := errs.NotFound("Organization", "id", org.ID.String())
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to update organization: ", err.Error())
		return nil, &errr
	}

	return &updatedOrganization, nil
}