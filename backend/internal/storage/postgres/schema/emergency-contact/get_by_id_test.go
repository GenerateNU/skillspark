package emergencycontact

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEmergencyContactByID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewEmergencyContactRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	ec := CreateTestEmergencyContact(t, ctx, testDB)

	result, err := repo.GetEmergencyContactByID(ctx, ec.ID)

	require.Nil(t, err)
	require.NotNil(t, result)

	assert.Equal(t, ec.ID, result.ID)
	assert.Equal(t, ec.GuardianID, result.GuardianID)
	assert.Equal(t, ec.Name, result.Name)
	assert.Equal(t, ec.PhoneNumber, result.PhoneNumber)
}

func TestGetEmergencyContactByID_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewEmergencyContactRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	result, err := repo.GetEmergencyContactByID(ctx, uuid.New())

	require.NotNil(t, err)
	assert.Nil(t, result)
}
