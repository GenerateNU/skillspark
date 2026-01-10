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

func (r *LocationRepository) GetLocationByID(ctx context.Context, id uuid.UUID) (*models.Location, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("location/sql/get_by_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, id)
	var location models.Location
	err = row.Scan(&location.ID, &location.Latitude, &location.Longitude, &location.Address, &location.City, &location.State, &location.ZipCode, &location.Country, &location.CreatedAt, &location.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Location", "id", id)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch location by id: ", err.Error())
		return nil, &err
	}

	return &location, nil
}
