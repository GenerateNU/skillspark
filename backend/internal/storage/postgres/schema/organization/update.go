package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *OrganizationRepository) UpdateOrganization(ctx context.Context, input *models.UpdateOrganizationInput, PfpS3Key *string) (*models.Organization, error) {
	query, err := schema.ReadSQLBaseScript("organization/sql/update.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	existing, httpErr := r.GetOrganizationByID(ctx, input.ID)
	if httpErr != nil {
		return nil, httpErr
	}

	if input.Body.Name != nil {
		existing.Name = *input.Body.Name
	}
	if input.Body.Active != nil {
		existing.Active = *input.Body.Active
	}
	if PfpS3Key != nil {
		existing.PfpS3Key = PfpS3Key
	}
	if input.Body.LocationID != nil {
		existing.LocationID = input.Body.LocationID
	}

	row := r.db.QueryRow(ctx, query,
		existing.Name,
		existing.Active,
		existing.PfpS3Key,
		existing.LocationID,
		input.ID,
	)

	var updatedOrganization models.Organization

	err = row.Scan(
		&updatedOrganization.ID,
		&updatedOrganization.Name,
		&updatedOrganization.Active,
		&updatedOrganization.PfpS3Key,
		&updatedOrganization.LocationID,
		&updatedOrganization.StripeAccountID,
		&updatedOrganization.StripeAccountActivated,
		&updatedOrganization.CreatedAt,
		&updatedOrganization.UpdatedAt,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			errr := errs.NotFound("Organization", "id", input.ID.String())
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to update organization: ", err.Error())
		return nil, &errr
	}

	return &updatedOrganization, nil
}
