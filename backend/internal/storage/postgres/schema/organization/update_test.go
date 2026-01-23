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

func TestUpdate(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Create an organization first
	orgID := uuid.New()
	org := &models.Organization{
		ID:        orgID,
		Name:      "Original Name",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createErr := repo.CreateOrganization(ctx, org)
	assert.Nil(t, createErr)

	// Update it
	org.Name = "Updated Name"
	org.Active = false
	org.UpdatedAt = time.Now()

	updateErr := repo.UpdateOrganization(ctx, org)
	assert.Nil(t, updateErr)

	// Verify update
	updated, getErr := repo.GetOrganizationByID(ctx, orgID)
	assert.Nil(t, getErr)
	assert.Equal(t, "Updated Name", updated.Name)
	assert.Equal(t, false, updated.Active)
}

func TestExecute_UpdateExisting(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Get existing test organization
	orgID := uuid.MustParse("40000000-0000-0000-0000-000000000004")
	org, err := repo.GetOrganizationByID(ctx, orgID)
	assert.Nil(t, err)

	// Update it
	org.Name = "Babel Street Updated"
	org.UpdatedAt = time.Now()

	updateErr := repo.UpdateOrganization(ctx, org)
	assert.Nil(t, updateErr)

	// Verify
	updated, getErr := repo.GetOrganizationByID(ctx, orgID)
	assert.Nil(t, getErr)
	assert.Equal(t, "Babel Street Updated", updated.Name)
}

func TestUpdate_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	// Try to update non-existent organization
	org := &models.Organization{
		ID:        uuid.New(),
		Name:      "Does Not Exist",
		Active:    true,
		UpdatedAt: time.Now(),
	}

	err := repo.UpdateOrganization(ctx, org)

	assert.NotNil(t, err)
}