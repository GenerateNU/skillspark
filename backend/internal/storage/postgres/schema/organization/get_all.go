package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/jackc/pgx/v5"
)

func (r *OrganizationRepository) GetAllOrganizations(ctx context.Context, pagination utils.Pagination, AcceptLanguage string) ([]models.Organization, error) {

	baseQuery, err := schema.ReadSQLBaseScript("get_all.sql", SqlOrganizationFiles)
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

	orgs, err := pgx.CollectRows(rows, scanOrganizationFactory(AcceptLanguage))
	if err != nil {
		errr := errs.InternalServerError("Failed to collect organizations: ", err.Error())
		return nil, &errr
	}

	return orgs, nil
}

func scanOrganizationFactory(AcceptLanguage string) func(pgx.CollectableRow) (models.Organization, error) {
	return func(row pgx.CollectableRow) (models.Organization, error) {
		var org models.Organization
		var aboutEN, aboutTH *string
		var lat, lng *float64
		var district *string
		err := row.Scan(
			&org.ID,
			&org.Name,
			&org.Active,
			&aboutEN,
			&aboutTH,
			&org.PfpS3Key,
			&org.LocationID,
			&org.StripeAccountID,
			&org.StripeAccountActivated,
			&org.CreatedAt,
			&org.UpdatedAt,
			&lat,
			&lng,
			&district,
		)
		if err != nil {
			return org, err
		}
		org.About = pickAbout(AcceptLanguage, aboutEN, aboutTH)
		org.Latitude = lat
		org.Longitude = lng
		org.District = district
		return org, nil
	}
}
