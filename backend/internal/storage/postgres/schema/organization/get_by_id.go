package organization

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *OrganizationRepository) GetOrganizationByID(ctx context.Context, id uuid.UUID, AcceptLanguage string) (*models.Organization, error) {
	query, err := schema.ReadSQLBaseScript("get_by_id.sql", SqlOrganizationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, id)
	var org models.Organization
	var aboutEN, aboutTH *string
	var rawLinks []byte
	err = row.Scan(
		&org.ID,
		&org.Name,
		&org.Active,
		&aboutEN,
		&aboutTH,
		&org.PfpS3Key,
		&org.LocationID,
		&rawLinks,
		&org.StripeAccountID,
		&org.StripeAccountActivated,
		&org.CreatedAt,
		&org.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Organization", "id", id)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch organization by id: ", err.Error())
		return nil, &err
	}

	org.About = pickAbout(AcceptLanguage, aboutEN, aboutTH)

	org.Links, err = scanLinks(rawLinks)
	if err != nil {
		errr := errs.InternalServerError("Failed to deserialize links: ", err.Error())
		return nil, &errr
	}

	return &org, nil
}

// getAboutColumns returns the raw about_en and about_th values for merging during updates.
func (r *OrganizationRepository) getAboutColumns(ctx context.Context, id uuid.UUID) (*string, *string, error) {
	var aboutEN, aboutTH *string
	err := r.db.QueryRow(ctx, "SELECT about_en, about_th FROM organization WHERE id = $1", id).Scan(&aboutEN, &aboutTH)
	if err != nil {
		return nil, nil, err
	}
	return aboutEN, aboutTH, nil
}
