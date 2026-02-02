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

func TestDeleteOrganization(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Create an organization to delete
	active := true
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "To Be Deleted"
		i.Body.Active = &active
		return i
	}()

	created, createErr := repo.CreateOrganization(ctx, input)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	// Delete the organization
	deleted, deleteErr := repo.DeleteOrganization(ctx, created.ID)
	require.Nil(t, deleteErr)
	require.NotNil(t, deleted)
	assert.Equal(t, created.ID, deleted.ID)
	assert.Equal(t, "To Be Deleted", deleted.Name)

	// Verify it's gone
	_, getErr := repo.GetOrganizationByID(ctx, created.ID)
	assert.NotNil(t, getErr)
}

func TestDeleteOrganization_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Try to delete non-existent organization
	deleted, err := repo.DeleteOrganization(ctx, uuid.New())

	require.NotNil(t, err)
	assert.Nil(t, deleted)
}

func TestDeleteOrganization_AlreadyDeleted(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Create organization
	active := true
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Delete Twice"
		i.Body.Active = &active
		return i
	}()

	created, createErr := repo.CreateOrganization(ctx, input)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	// First delete should succeed
	deleted1, deleteErr1 := repo.DeleteOrganization(ctx, created.ID)
	require.Nil(t, deleteErr1)
	require.NotNil(t, deleted1)

	// Second delete should fail
	deleted2, deleteErr2 := repo.DeleteOrganization(ctx, created.ID)
	require.NotNil(t, deleteErr2)
	assert.Nil(t, deleted2)
}
