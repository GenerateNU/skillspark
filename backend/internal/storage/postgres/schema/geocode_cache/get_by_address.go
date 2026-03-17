package geocode_cache

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *GeocodeCacheRepository) GetGeocodeCache(ctx context.Context, address string) (*models.GeocodeCache, error) {
	query, err := schema.ReadSQLBaseScript("get_by_address.sql", SqlGeocodeCacheFiles)
	if err != nil {
		e := errs.InternalServerError("failed to read base query: ", err.Error())
		return nil, &e
	}

	row := r.db.QueryRow(ctx, query, address)
	var entry models.GeocodeCache
	err = row.Scan(&entry.Address, &entry.RawAddress, &entry.Latitude, &entry.Longitude, &entry.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			e := errs.NotFound("GeocodeCache", "address", address)
			return nil, &e
		}
		e := errs.InternalServerError("failed to fetch geocode cache: ", err.Error())
		return nil, &e
	}

	return &entry, nil
}
