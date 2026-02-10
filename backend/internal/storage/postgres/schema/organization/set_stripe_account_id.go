package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *OrganizationRepository) SetStripeAccountID(ctx context.Context, orgID uuid.UUID, stripeAccountID string) (*models.Organization, error) {
	query, err := schema.ReadSQLBaseScript("organization/sql/set_stripe_account_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}
	
	row := r.db.QueryRow(ctx, query, stripeAccountID, orgID)
	
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
			errr := errs.NotFound("Organization", "id", orgID.String())
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to set stripe account: ", err.Error())
		return nil, &errr
	}
	
	return &updatedOrganization, nil
}