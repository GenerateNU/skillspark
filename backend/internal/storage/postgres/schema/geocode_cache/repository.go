package geocode_cache

import "github.com/jackc/pgx/v5/pgxpool"

type GeocodeCacheRepository struct {
	db *pgxpool.Pool
}

func NewGeocodeCacheRepository(db *pgxpool.Pool) *GeocodeCacheRepository {
	return &GeocodeCacheRepository{db: db}
}
