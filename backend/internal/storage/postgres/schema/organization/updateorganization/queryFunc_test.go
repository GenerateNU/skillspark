package updateorganization

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization/createorganization"
	"skillspark/internal/storage/postgres/schema/organization/getorganizationbyid"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
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

	createErr := createorganization.Execute(ctx, testDB, org)
	assert.Nil(t, createErr)

	// Update it
	org.Name = "Updated Name"
	org.Active = false
	org.UpdatedAt = time.Now()

	updateErr := Execute(ctx, testDB, org)
	assert.Nil(t, updateErr)

	// Verify update
	updated, getErr := getorganizationbyid.Execute(ctx, testDB, orgID)
	assert.Nil(t, getErr)
	assert.Equal(t, "Updated Name", updated.Name)
	assert.Equal(t, false, updated.Active)
}

func TestExecute_UpdateExisting(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	// Get existing test organization
	orgID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	org, err := getorganizationbyid.Execute(ctx, testDB, orgID)
	assert.Nil(t, err)

	// Update it
	org.Name = "Babel Street Updated"
	org.UpdatedAt = time.Now()

	updateErr := Execute(ctx, testDB, org)
	assert.Nil(t, updateErr)

	// Verify
	updated, getErr := getorganizationbyid.Execute(ctx, testDB, orgID)
	assert.Nil(t, getErr)
	assert.Equal(t, "Babel Street Updated", updated.Name)
}

func TestExecute_NotFound(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	// Try to update non-existent organization
	org := &models.Organization{
		ID:        uuid.New(),
		Name:      "Does Not Exist",
		Active:    true,
		UpdatedAt: time.Now(),
	}

	err := Execute(ctx, testDB, org)

	assert.NotNil(t, err)
}