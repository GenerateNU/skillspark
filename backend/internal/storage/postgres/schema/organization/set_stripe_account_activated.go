package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *OrganizationRepository) SetStripeAccountActivated(ctx context.Context, stripeAccountID string, activated bool) (*models.Organization, error) {
	query, err := schema.ReadSQLBaseScript("organization/sql/set_stripe_account_activated.sql", SqlOrganizationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}
	
	row := r.db.QueryRow(ctx, query, activated, stripeAccountID)
	
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
			errr := errs.NotFound("Organization", "stripe_account_id", stripeAccountID)
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to update stripe activation: ", err.Error())
		return nil, &errr
	}
	
	return &updatedOrganization, nil
}