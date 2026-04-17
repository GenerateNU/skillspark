package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *OrganizationRepository) UpdateOrganization(ctx context.Context, input *models.UpdateOrganizationDBInput, PfpS3Key *string) (*models.Organization, error) {
	query, err := schema.ReadSQLBaseScript("update.sql", SqlOrganizationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	existing, httpErr := r.GetOrganizationByID(ctx, input.ID, input.AcceptLanguage)
	if httpErr != nil {
		return nil, httpErr
	}

	existingAboutEN, existingAboutTH, err := r.getAboutColumns(ctx, input.ID)
	if err != nil {
		errr := errs.InternalServerError("Failed to load existing about columns: ", err.Error())
		return nil, &errr
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
	if input.Body.Links != nil {
		existing.Links = *input.Body.Links
	}
	if input.Body.AboutEN != nil {
		existingAboutEN = input.Body.AboutEN
	}
	if input.Body.AboutTH != nil {
		existingAboutTH = input.Body.AboutTH
	}

	jsonLinks, err := byteSliceLinks(existing.Links)
	if err != nil {
		errr := errs.InternalServerError("Failed to serialize links: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		existing.Name,
		existing.Active,
		existing.PfpS3Key,
		existing.LocationID,
		jsonLinks,
		existingAboutEN,
		existingAboutTH,
		input.ID,
	)

	var updatedOrganization models.Organization
	var aboutEN, aboutTH *string

	var rawLinks []byte
	err = row.Scan(
		&updatedOrganization.ID,
		&updatedOrganization.Name,
		&updatedOrganization.Active,
		&aboutEN,
		&aboutTH,
		&updatedOrganization.PfpS3Key,
		&updatedOrganization.LocationID,
		&rawLinks,
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

	updatedOrganization.About = pickAbout(input.AcceptLanguage, aboutEN, aboutTH)

	updatedOrganization.Links, err = scanLinks(rawLinks)
	if err != nil {
		errr := errs.InternalServerError("Failed to deserialize links: ", err.Error())
		return nil, &errr
	}

	return &updatedOrganization, nil
}
