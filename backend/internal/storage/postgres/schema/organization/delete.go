package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *OrganizationRepository) DeleteOrganization(ctx context.Context, id uuid.UUID) (*models.Organization, error) {
	query, err := schema.ReadSQLBaseScript("delete.sql", SqlOrganizationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, id)

	var deletedOrganization models.Organization

	err = row.Scan(
		&deletedOrganization.ID,
		&deletedOrganization.Name,
		&deletedOrganization.Active,
		&deletedOrganization.PfpS3Key,
		&deletedOrganization.LocationID,
		&deletedOrganization.StripeAccountID,
		&deletedOrganization.StripeAccountActivated,
		&deletedOrganization.CreatedAt,
		&deletedOrganization.UpdatedAt,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			errr := errs.NotFound("Organization", "id", id.String())
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to delete organization: ", err.Error())
		return nil, &errr
	}

	return &deletedOrganization, nil
}
