package organization

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Create an organization to delete
	orgID := uuid.New()
	org := &models.Organization{
		ID:        orgID,
		Name:      "To Be Deleted",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createErr := repo.CreateOrganization(ctx, org)
	assert.Nil(t, createErr)

	deleteErr := repo.DeleteOrganization(ctx, orgID)
	assert.Nil(t, deleteErr)

	// Verify it's gone
	deleted, getErr := repo.GetOrganizationByID(ctx, orgID)
	assert.NotNil(t, getErr)
	assert.Nil(t, deleted)
}

func TestDelete_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Try to delete non-existent organization
	err := repo.DeleteOrganization(ctx, uuid.New())

	assert.NotNil(t, err)
}

func TestExecute_AlreadyDeleted(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	orgID := uuid.New()
	org := &models.Organization{
		ID:        orgID,
		Name:      "Delete Twice",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createErr := repo.CreateOrganization(ctx, org)
	assert.Nil(t, createErr)

	deleteErr1 := repo.DeleteOrganization(ctx, orgID)
	assert.Nil(t, deleteErr1)

	deleteErr2 := repo.DeleteOrganization(ctx, orgID)
	assert.NotNil(t, deleteErr2)
}