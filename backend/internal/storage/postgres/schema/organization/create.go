package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *OrganizationRepository) CreateOrganization(ctx context.Context, input *models.CreateOrganizationDBInput, PfpS3Key *string) (*models.Organization, error) {

	query, err := schema.ReadSQLBaseScript("create.sql", SqlOrganizationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	jsonLinks, err := byteSliceLinks(input.Body.Links)
	if err != nil {
		errr := errs.InternalServerError("Failed to serialize links: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		input.Body.Name,
		input.Body.Active,
		PfpS3Key,
		input.Body.LocationID,
		jsonLinks,
		input.Body.AboutEN,
		input.Body.AboutTH,
	)

	var createdOrganization models.Organization
	var aboutEN, aboutTH *string
	var rawLinks []byte
	err = row.Scan(
		&createdOrganization.ID,
		&createdOrganization.Name,
		&createdOrganization.Active,
		&aboutEN,
		&aboutTH,
		&createdOrganization.PfpS3Key,
		&createdOrganization.LocationID,
		&rawLinks,
		&createdOrganization.StripeAccountID,
		&createdOrganization.StripeAccountActivated,
		&createdOrganization.CreatedAt,
		&createdOrganization.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	createdOrganization.About = pickAbout(input.AcceptLanguage, aboutEN, aboutTH)

	createdOrganization.Links, err = scanLinks(rawLinks)
	if err != nil {
		errr := errs.InternalServerError("Failed to deserialize links: ", err.Error())
		return nil, &errr
	}

	return &createdOrganization, nil
}
