package emergencycontact

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteEmergencyContact(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewEmergencyContactRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	ec := CreateTestEmergencyContact(t, ctx, testDB)

	deleted, err := repo.DeleteEmergencyContact(ctx, ec.ID)

	require.Nil(t, err)
	require.NotNil(t, deleted)

	assert.Equal(t, ec.ID, deleted.Body.ID)
	assert.Equal(t, ec.GuardianID, deleted.Body.GuardianID)
}

func TestDeleteEmergencyContact_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewEmergencyContactRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	deleted, err := repo.DeleteEmergencyContact(ctx, uuid.New())

	require.NotNil(t, err)
	assert.Nil(t, deleted)
}
