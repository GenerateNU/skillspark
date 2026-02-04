package registration

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"
	guardian "skillspark/internal/storage/postgres/schema/guardian"

	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

	// delete a guardian that has children, should delete guardian and children fks from registration
	input1 := &models.GetRegistrationsByGuardianIDInput{
		GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
	}
	guardianId := uuid.MustParse("80000000-0000-0000-0000-000000000001")
	childId := uuid.MustParse("30000000-0000-0000-0000-000000000001")

	input2 := &models.GetRegistrationByIDInput{
		ID: uuid.MustParse("80000000-0000-0000-0000-000000000001"),
	}

	registrations, err := repo.GetRegistrationsByGuardianID(ctx, input1)
	reg1 := registrations.Body.Registrations[0]
	assert.Nil(t, err)
	assert.NotNil(t, reg1)
	assert.Equal(t, guardianId, reg1.ID)
	assert.Equal(t, &childId, reg1.ChildID)

	guardian, guardianErr := guardianRepo.DeleteGuardian(ctx, uuid.MustParse("11111111-1111-1111-1111-111111111111"))
	assert.Nil(t, guardianErr)
	assert.NotNil(t, guardian)

	guardian, guardianErr = guardianRepo.GetGuardianByID(ctx, guardian.ID) // guardian deleted
	assert.Nil(t, guardian)
	assert.NotNil(t, guardianErr)

	registration, getErr := repo.GetRegistrationByID(ctx, input2)
	assert.Nil(t, getErr)
	assert.NotNil(t, registration)
	assert.Nil(t, registration.Body.GuardianID)
	assert.Nil(t, registration.Body.ChildID)
}