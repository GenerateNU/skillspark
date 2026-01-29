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
	t.Parallel()

	// Create an organization first
	active := true
	createInput := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Original Name"
		i.Body.Active = &active
		return i
	}()

	created, createErr := repo.CreateOrganization(ctx, createInput)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	// Update it
	newName := "Updated Name"
	newActive := false
	updateInput := &models.UpdateOrganizationInput{
		ID: created.ID,
		Body: struct {
			Name       *string    `json:"name,omitempty" minLength:"1" maxLength:"255" doc:"Organization name"`
			Active     *bool      `json:"active,omitempty" doc:"Active status"`
			PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
			LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
		}{
			Name:   &newName,
			Active: &newActive,
		},
	}

	updated, updateErr := repo.UpdateOrganization(ctx, updateInput)
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
	t.Parallel()

	// Create organization
	active := true
	createInput := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Test Org"
		i.Body.Active = &active
		return i
	}()

	created, createErr := repo.CreateOrganization(ctx, createInput)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	// Update with location
	locationID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	newName := "Test Org with Location"
	updateInput := &models.UpdateOrganizationInput{
		ID: created.ID,
		Body: struct {
			Name       *string    `json:"name,omitempty" minLength:"1" maxLength:"255" doc:"Organization name"`
			Active     *bool      `json:"active,omitempty" doc:"Active status"`
			PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
			LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
		}{
			Name:       &newName,
			LocationID: &locationID,
		},
	}

	updated, updateErr := repo.UpdateOrganization(ctx, updateInput)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, "Test Org with Location", updated.Name)
	assert.Equal(t, &locationID, updated.LocationID)
}

func TestUpdateOrganization_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	// Try to update non-existent organization
	nonExistentID := uuid.New()
	newName := "Does Not Exist"
	updateInput := &models.UpdateOrganizationInput{
		ID: nonExistentID,
		Body: struct {
			Name       *string    `json:"name,omitempty" minLength:"1" maxLength:"255" doc:"Organization name"`
			Active     *bool      `json:"active,omitempty" doc:"Active status"`
			PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
			LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
		}{
			Name: &newName,
		},
	}

	updated, err := repo.UpdateOrganization(ctx, updateInput)

	require.NotNil(t, err)
	assert.Nil(t, updated)
}
