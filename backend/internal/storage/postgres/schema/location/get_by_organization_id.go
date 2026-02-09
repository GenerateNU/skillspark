package location

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *LocationRepository) GetLocationByOrganizationID(ctx context.Context, orgID uuid.UUID) (*models.Location, error) {
	query, err := schema.ReadSQLBaseScript("location/sql/get_by_organization_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, orgID)
	var location models.Location
	err = row.Scan(
		&location.ID, 
		&location.Latitude, 
		&location.Longitude, 
		&location.AddressLine1, 
		&location.AddressLine2, 
		&location.Subdistrict, 
		&location.District, 
		&location.Province, 
		&location.PostalCode, 
		&location.Country, 
		&location.CreatedAt, 
		&location.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Location", "organization_id", orgID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch location by organization id: ", err.Error())
		return nil, &err
	}

	return &location, nil
}