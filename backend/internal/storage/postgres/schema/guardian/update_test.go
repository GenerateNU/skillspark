package guardian

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGuardianRepository_Update_David_Kim(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianInput := func() *models.UpdateGuardianInput {
		input := &models.UpdateGuardianInput{}
		input.ID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
		input.Body.Name = "Updated David"
		input.Body.Email = "updated.david@example.com"
		input.Body.Username = "udavid"
		input.Body.LanguagePreference = "en"
		return input
	}()

	guardian, err := repo.UpdateGuardian(ctx, guardianInput)
	if err != nil {
		t.Fatalf("Failed to update guardian: %v", err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, guardian)
	assert.Equal(t, guardianInput.Body.Name, guardian.Name)
	assert.NotNil(t, guardian.CreatedAt)
	assert.NotNil(t, guardian.UpdatedAt)
	assert.Equal(t, guardianInput.ID, guardian.ID)

	retrievedGuardian, err := repo.GetGuardianByID(ctx, guardianInput.ID)
	if err != nil {
		t.Fatalf("Failed to retrieve guardian: %v", err)
	}

	assert.NotNil(t, retrievedGuardian)
	assert.Equal(t, guardianInput.Body.Name, retrievedGuardian.Name)
	assert.Equal(t, guardianInput.ID, retrievedGuardian.ID)
}

func TestGuardianRepository_Update_Errors(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()

	t.Run("Update Non-Existent Guardian", func(t *testing.T) {
		input := &models.UpdateGuardianInput{
			ID: uuid.New(),
		}
		input.Body.Name = "Ghost"

		guardian, err := repo.UpdateGuardian(ctx, input)
		assert.Error(t, err)
		assert.Nil(t, guardian)
	})
}

func TestGuardianRepository_Update_DoesNotModifyStripeCustomerID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardian := CreateTestGuardian(t, ctx, testDB)
	stripeCustomerID := "cus_test123"
	
	repo.SetStripeCustomerID(ctx, guardian.ID, stripeCustomerID)

	updateInput := &models.UpdateGuardianInput{}
	updateInput.ID = guardian.ID
	updateInput.Body.Name = "Updated Name"

	updated, err := repo.UpdateGuardian(ctx, updateInput)

	assert.Nil(t, err)
	assert.Equal(t, "Updated Name", updated.Name)
	assert.Equal(t, stripeCustomerID, *updated.StripeCustomerID)
}