package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

)

func (r *OrganizationRepository) CreateOrganization(ctx context.Context, org *models.Organization) *errs.HTTPError {
	query, err := schema.ReadSQLBaseScript("organization/sql/create.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return &errr
	}

	_, err = r.db.Exec(ctx, query,
		org.ID,
		org.Name,
		org.Active,
		org.PfpS3Key,
		org.LocationID,
		org.CreatedAt,
		org.UpdatedAt,
	)

	if err != nil {
		errr := errs.InternalServerError("Failed to create organization: ", err.Error())
		return &errr
	}

	return nil
}