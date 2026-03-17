package geocode_cache

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *GeocodeCacheRepository) CreateGeocodeCache(ctx context.Context, normalizedAddress, rawAddress string, lat, lng float64) (*models.GeocodeCache, error) {
	query, err := schema.ReadSQLBaseScript("create.sql", SqlGeocodeCacheFiles)
	if err != nil {
		e := errs.InternalServerError("failed to read base query: ", err.Error())
		return nil, &e
	}

	row := r.db.QueryRow(ctx, query, normalizedAddress, rawAddress, lat, lng)
	var entry models.GeocodeCache
	err = row.Scan(&entry.Address, &entry.RawAddress, &entry.Latitude, &entry.Longitude, &entry.CreatedAt)
	if err != nil {
		e := errs.InternalServerError("failed to create geocode cache entry: ", err.Error())
		return nil, &e
	}

	return &entry, nil
}
