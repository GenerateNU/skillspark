package organization

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateOrganization(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Create an organization first
	active := true
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Original Name"
		i.Body.Active = &active
		return i
	}()

	created, createErr := repo.CreateOrganization(ctx, input)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	// Update it
	created.Name = "Updated Name"
	created.Active = false

	updated, updateErr := repo.UpdateOrganization(ctx, created)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, "Updated Name", updated.Name)
	assert.False(t, updated.Active)

	// Verify update persisted
	fetched, getErr := repo.GetOrganizationByID(ctx, created.ID)
	require.Nil(t, getErr)
	assert.Equal(t, "Updated Name", fetched.Name)
	assert.False(t, fetched.Active)
}

func TestUpdateOrganization_WithLocation(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Create organization
	active := true
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Test Org"
		i.Body.Active = &active
		return i
	}()

	created, createErr := repo.CreateOrganization(ctx, input)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	// Update with location
	locationID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	created.LocationID = &locationID
	created.Name = "Test Org with Location"

	updated, updateErr := repo.UpdateOrganization(ctx, created)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, "Test Org with Location", updated.Name)
	assert.Equal(t, &locationID, updated.LocationID)
}

func TestUpdateOrganization_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Try to update non-existent organization
	org := &models.Organization{
		ID:     uuid.New(),
		Name:   "Does Not Exist",
		Active: true,
	}

	updated, err := repo.UpdateOrganization(ctx, org)

	require.NotNil(t, err)
	assert.Nil(t, updated)
}