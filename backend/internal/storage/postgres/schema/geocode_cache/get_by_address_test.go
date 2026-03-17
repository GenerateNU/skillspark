package geocode_cache

import (
	"context"
	"net/http"
	"testing"

	"skillspark/internal/errs"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
)

func TestGeocodeCacheRepository_GetByAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()
	t.Parallel()

	seeded := createTestGeocodeCacheEntry(t, ctx, testDB)

	repo := NewGeocodeCacheRepository(testDB)
	entry, err := repo.GetGeocodeCache(ctx, seeded.Address)

	assert.Nil(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, seeded.Address, entry.Address)
	assert.Equal(t, seeded.RawAddress, entry.RawAddress)
	assert.Equal(t, seeded.Latitude, entry.Latitude)
	assert.Equal(t, seeded.Longitude, entry.Longitude)
	assert.False(t, entry.CreatedAt.IsZero())
}

func TestGeocodeCacheRepository_GetByAddress_NormalizedKeyOnly(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGeocodeCacheRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	normalized := "siam square, bangkok, thailand"
	raw := "Siam Square, Bangkok, Thailand"

	_, err := repo.CreateGeocodeCache(ctx, normalized, raw, 13.7455, 100.5340)
	assert.Nil(t, err)

	// Hit using normalized address
	entry, err := repo.GetGeocodeCache(ctx, normalized)
	assert.Nil(t, err)
	assert.NotNil(t, entry)

	// Miss using raw address — cache key is the normalized form only
	miss, err := repo.GetGeocodeCache(ctx, raw)
	assert.NotNil(t, err)
	assert.Nil(t, miss)
	httpErr, ok := err.(*errs.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, http.StatusNotFound, httpErr.Code)
}

func TestGeocodeCacheRepository_GetByAddress_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGeocodeCacheRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	entry, err := repo.GetGeocodeCache(ctx, "this address does not exist")

	assert.Nil(t, entry)
	assert.NotNil(t, err)
	httpErr, ok := err.(*errs.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, http.StatusNotFound, httpErr.Code)
}
