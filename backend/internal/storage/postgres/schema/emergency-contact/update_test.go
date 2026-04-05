package emergencycontact

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateEmergencyContact(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewEmergencyContactRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	ec := CreateTestEmergencyContact(t, ctx, testDB)

	input := &models.UpdateEmergencyContactInput{
		ID: ec.ID,
	}
	input.Body.Name = "Updated Name"
	input.Body.GuardianID = ec.GuardianID
	input.Body.PhoneNumber = "+19999999999"

	updated, err := repo.UpdateEmergencyContact(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, updated)

	assert.Equal(t, ec.ID, updated.Body.ID)
	assert.Equal(t, "Updated Name", updated.Body.Name)
	assert.Equal(t, "+19999999999", updated.Body.PhoneNumber)
	assert.Equal(t, ec.GuardianID, updated.Body.GuardianID)
}
