package emergencycontact

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEmergencyContactByGuardianID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewEmergencyContactRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	ec := CreateTestEmergencyContact(t, ctx, testDB)

	result, err := repo.GetEmergencyContactByGuardianID(ctx, ec.GuardianID)

	require.Nil(t, err)
	require.NotNil(t, result)

	assert.NotEmpty(t, result)
	assert.Equal(t, ec.GuardianID, result[0].GuardianID)
	assert.Equal(t, ec.ID, result[0].ID)
}

func TestGetEmergencyContactByGuardianID_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewEmergencyContactRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	result, err := repo.GetEmergencyContactByGuardianID(ctx, uuid.New())

	require.Nil(t, err)
	assert.Empty(t, result)
}
