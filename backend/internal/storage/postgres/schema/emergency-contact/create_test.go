package emergencycontact

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateEmergencyContact(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewEmergencyContactRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	g := guardian.CreateTestGuardian(t, ctx, testDB)

	input := &models.CreateEmergencyContactInput{}
	input.Body.Name = "Jane Doe"
	input.Body.GuardianID = g.ID
	input.Body.PhoneNumber = "+16462996961"

	created, err := repo.CreateEmergencyContact(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)

	assert.NotEqual(t, uuid.Nil, created.Body.ID)
	assert.Equal(t, "Jane Doe", created.Body.Name)
	assert.Equal(t, g.ID, created.Body.GuardianID)
	assert.Equal(t, "+16462996961", created.Body.PhoneNumber)
	assert.NotZero(t, created.Body.CreatedAt)
	assert.NotZero(t, created.Body.UpdatedAt)
}

func TestCreateEmergencyContact_FailsInvalidGuardian(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewEmergencyContactRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	input := &models.CreateEmergencyContactInput{}
	input.Body.Name = "Jane Doe"
	input.Body.GuardianID = uuid.New()
	input.Body.PhoneNumber = "+16462996961"

	created, err := repo.CreateEmergencyContact(ctx, input)

	require.NotNil(t, err)
	assert.Nil(t, created)
}
