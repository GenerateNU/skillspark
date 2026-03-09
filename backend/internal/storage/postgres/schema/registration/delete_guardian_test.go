package registration

import (
	"context"
	"testing"

	guardian "skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/testutil"

	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGuardianRepository_DeleteGuardian_SetFieldsNull(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewRegistrationRepository(testDB)
	guardianRepo := guardian.NewGuardianRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	regID := uuid.MustParse("80000000-0000-0000-0000-000000000001")
	childID := uuid.MustParse("30000000-0000-0000-0000-000000000001")
	
	input1 := &models.GetRegistrationsByGuardianIDInput{
		GuardianID: guardianID,
	}
	input2 := &models.GetRegistrationByIDInput{
		ID: regID,
	}

	registrations, err := repo.GetRegistrationsByGuardianID(ctx, input1)
	require.NoError(t, err)
	require.NotEmpty(t, registrations.Body.Registrations)
	
	reg1 := registrations.Body.Registrations[0]
	assert.Equal(t, regID, reg1.ID)
	assert.Equal(t, childID, reg1.ChildID)

	deletedGuardian, err := guardianRepo.DeleteGuardian(ctx, guardianID, nil)
	require.NoError(t, err)
	require.NotNil(t, deletedGuardian)

	fetchedGuardian, err := guardianRepo.GetGuardianByID(ctx, deletedGuardian.ID)
	assert.Error(t, err)
	assert.Nil(t, fetchedGuardian)

	registration, err := repo.GetRegistrationByID(ctx, input2, nil)
	require.NoError(t, err)
	require.NotNil(t, registration)
	
	assert.Equal(t, models.RegistrationStatusCancelled, registration.Body.Status)
	assert.NotEqual(t, uuid.Nil, registration.Body.ID)
}