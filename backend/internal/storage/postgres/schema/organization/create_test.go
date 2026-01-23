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

func TestCreate(t *testing.T) {

	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	org := &models.Organization{
		ID:        uuid.New(),
		Name:      "Test Corp",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.CreateOrganization(ctx, org)

	assert.Nil(t, err)
}

func TestExecute_WithLocation(t *testing.T) {

	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	locationID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	org := &models.Organization{
		ID:         uuid.New(),
		Name:       "Test Corp with Location",
		Active:     true,
		LocationID: &locationID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := repo.CreateOrganization(ctx, org)

	assert.Nil(t, err)
}

func TestExecute_DuplicateID(t *testing.T) {

	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()

	orgID := uuid.New()
	org1 := &models.Organization{
		ID:        orgID,
		Name:      "First Org",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.CreateOrganization(ctx, org1)
	assert.Nil(t, err)

	// Try to create another with same ID
	org2 := &models.Organization{
		ID:        orgID,
		Name:      "Second Org",
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err2 := repo.CreateOrganization(ctx, org2)
	assert.NotNil(t, err2)
}
