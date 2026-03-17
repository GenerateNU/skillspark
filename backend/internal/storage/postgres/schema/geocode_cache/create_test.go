package geocode_cache

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
)

func TestGeocodeCacheRepository_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGeocodeCacheRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	normalized := "123 sukhumvit road, bangkok, thailand"
	raw := "123 Sukhumvit Road, Bangkok, Thailand"

	entry, err := repo.CreateGeocodeCache(ctx, normalized, raw, 13.7300, 100.5697)

	assert.Nil(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, normalized, entry.Address)
	assert.Equal(t, raw, entry.RawAddress)
	assert.Equal(t, 13.7300, entry.Latitude)
	assert.Equal(t, 100.5697, entry.Longitude)
	assert.False(t, entry.CreatedAt.IsZero())
}

func TestGeocodeCacheRepository_Create_Upsert(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGeocodeCacheRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	normalized := "silom road, bangkok, thailand"

	_, err := repo.CreateGeocodeCache(ctx, normalized, "Silom Road, Bangkok, Thailand", 13.7245, 100.5287)
	assert.Nil(t, err)

	// Insert again with updated coordinates — should upsert cleanly
	updated, err := repo.CreateGeocodeCache(ctx, normalized, "Silom Road, Bangkok, Thailand", 13.7250, 100.5290)

	assert.Nil(t, err)
	assert.NotNil(t, updated)
	assert.Equal(t, 13.7250, updated.Latitude)
	assert.Equal(t, 100.5290, updated.Longitude)
}

func TestGeocodeCacheRepository_Create_PreservesRawAddress(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGeocodeCacheRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	normalized := "  wat pho, phra nakhon, bangkok  "
	raw := "  Wat Pho, Phra Nakhon, Bangkok  "

	entry, err := repo.CreateGeocodeCache(ctx, normalized, raw, 13.7465, 100.4927)

	assert.Nil(t, err)
	assert.NotNil(t, entry)
	// Raw address is stored as-is (normalization is the caller's responsibility)
	assert.Equal(t, raw, entry.RawAddress)
	assert.Equal(t, normalized, entry.Address)
}

func TestGeocodeCacheRepository_Create_NormalizedAddressIsKey(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGeocodeCacheRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	normalized := "lumpini park, bangkok, thailand"

	// Insert with two different raw representations of the same normalized address
	_, err := repo.CreateGeocodeCache(ctx, normalized, "Lumpini Park, Bangkok, Thailand", 13.7306, 100.5418)
	assert.Nil(t, err)

	updated, err := repo.CreateGeocodeCache(ctx, normalized, "LUMPINI PARK, BANGKOK, THAILAND", 13.7306, 100.5418)
	assert.Nil(t, err)
	assert.NotNil(t, updated)

	// Only one row should exist — the normalized address is the primary key
	fetched, err := repo.GetGeocodeCache(ctx, normalized)
	assert.Nil(t, err)
	assert.NotNil(t, fetched)
	assert.Equal(t, normalized, fetched.Address)
	assert.Equal(t, "LUMPINI PARK, BANGKOK, THAILAND", fetched.RawAddress)
}
