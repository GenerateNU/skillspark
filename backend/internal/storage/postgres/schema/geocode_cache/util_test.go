package geocode_cache

import (
	"context"
	"skillspark/internal/models"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func createTestGeocodeCacheEntry(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.GeocodeCache {
	t.Helper()

	repo := NewGeocodeCacheRepository(db)

	entry, err := repo.CreateGeocodeCache(ctx,
		"123 sukhumvit road, bangkok, thailand",
		"123 Sukhumvit Road, Bangkok, Thailand",
		13.7300,
		100.5697,
	)

	require.NoError(t, err)
	require.NotNil(t, entry)

	return entry
}
