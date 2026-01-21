package deleteorganization

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

	// Create an organization to delete
	orgID := uuid.New()
	org := &models.Organization{
		ID:        orgID,
		Name:      "To Be Deleted",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createErr := createorganization.Execute(ctx, testDB, org)
	assert.Nil(t, createErr)

	// Delete it
	deleteErr := Execute(ctx, testDB, orgID)
	assert.Nil(t, deleteErr)

	// Verify it's gone
	deleted, getErr := getorganizationbyid.Execute(ctx, testDB, orgID)
	assert.NotNil(t, getErr)
	assert.Nil(t, deleted)
}

func TestExecute_NotFound(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	// Try to delete non-existent organization
	err := Execute(ctx, testDB, uuid.New())

	assert.NotNil(t, err)
}

func TestExecute_AlreadyDeleted(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	// Create and delete
	orgID := uuid.New()
	org := &models.Organization{
		ID:        orgID,
		Name:      "Delete Twice",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createErr := createorganization.Execute(ctx, testDB, org)
	assert.Nil(t, createErr)

	deleteErr1 := Execute(ctx, testDB, orgID)
	assert.Nil(t, deleteErr1)

	// Try to delete again
	deleteErr2 := Execute(ctx, testDB, orgID)
	assert.NotNil(t, deleteErr2)
}