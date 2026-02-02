package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *OrganizationRepository) CreateOrganization(ctx context.Context, input *models.CreateOrganizationInput) (*models.Organization, *errs.HTTPError) {

	query, err := schema.ReadSQLBaseScript("organization/sql/create.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		input.Body.Name,
		input.Body.Active,
		input.Body.PfpS3Key,
		input.Body.LocationID,
	)

	var createdOrganization models.Organization

	err = row.Scan(
		&createdOrganization.ID,
		&createdOrganization.Name,
		&createdOrganization.Active,
		&createdOrganization.PfpS3Key,
		&createdOrganization.LocationID,
		&createdOrganization.CreatedAt,
		&createdOrganization.UpdatedAt,
	)
	if err != nil {
		errr := errs.InternalServerError("Failed to create organization: ", err.Error())
		return nil, &errr
	}

	return &createdOrganization, nil
}
