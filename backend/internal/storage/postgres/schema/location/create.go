package location

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *LocationRepository) CreateLocation(ctx context.Context, location *models.CreateLocationInput) (*models.Location, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("location/sql/create.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, location.Body.Latitude, location.Body.Longitude, location.Body.Address, location.Body.City, location.Body.State, location.Body.ZipCode, location.Body.Country)

	var createdLocation models.Location

	err = row.Scan(&createdLocation.ID, &createdLocation.Latitude, &createdLocation.Longitude, &createdLocation.Address, &createdLocation.City, &createdLocation.State, &createdLocation.ZipCode, &createdLocation.Country, &createdLocation.CreatedAt, &createdLocation.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to create location: ", err.Error())
		return nil, &err
	}

	return &createdLocation, nil
}
