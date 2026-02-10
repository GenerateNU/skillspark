package guardianpaymentmethod

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateGuardianPaymentMethod_SetAsDefault(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianPaymentMethodRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianID := uuid.MustParse("88888888-8888-8888-8888-888888888888")

	input1 := &models.CreateGuardianPaymentMethodInput{}
	input1.Body.GuardianID = guardianID
	input1.Body.StripePaymentMethodID = "pm_test_old_default"
	input1.Body.IsDefault = true
	oldDefault, err := repo.CreateGuardianPaymentMethod(ctx, input1)
	require.Nil(t, err)
	require.True(t, oldDefault.IsDefault)

	input2 := &models.CreateGuardianPaymentMethodInput{}
	input2.Body.GuardianID = guardianID
	input2.Body.StripePaymentMethodID = "pm_test_new_default"
	input2.Body.IsDefault = false
	newCard, err := repo.CreateGuardianPaymentMethod(ctx, input2)
	require.Nil(t, err)
	require.False(t, newCard.IsDefault)

	updated, updateErr := repo.UpdateGuardianPaymentMethod(ctx, newCard.ID, true)

	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.True(t, updated.IsDefault)
	assert.Equal(t, newCard.ID, updated.ID)

	oldDefaultCheck, err := repo.GetPaymentMethodsByGuardianID(ctx, guardianID)
	require.Nil(t, err)
	
	defaultCount := 0
	for _, pm := range oldDefaultCheck {
		if pm.IsDefault {
			defaultCount++
			assert.Equal(t, newCard.ID, pm.ID)
		}
	}
	assert.Equal(t, 1, defaultCount)
}

func TestUpdateGuardianPaymentMethod_UnsetDefault(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianPaymentMethodRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianID := uuid.MustParse("88888888-8888-8888-8888-888888888888")

	input := &models.CreateGuardianPaymentMethodInput{}
	input.Body.GuardianID = guardianID
	input.Body.StripePaymentMethodID = "pm_test_unset"
	input.Body.IsDefault = true

	created, createErr := repo.CreateGuardianPaymentMethod(ctx, input)
	require.Nil(t, createErr)
	require.True(t, created.IsDefault)

	updated, updateErr := repo.UpdateGuardianPaymentMethod(ctx, created.ID, false)

	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.False(t, updated.IsDefault)
}

func TestUpdateGuardianPaymentMethod_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianPaymentMethodRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	updated, err := repo.UpdateGuardianPaymentMethod(ctx, uuid.New(), true)

	require.NotNil(t, err)
	assert.Nil(t, updated)
}

func TestUpdateGuardianPaymentMethod_OnlyOneDefault(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianPaymentMethodRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianID := uuid.MustParse("88888888-8888-8888-8888-888888888888")

	input1 := &models.CreateGuardianPaymentMethodInput{}
	input1.Body.GuardianID = guardianID
	input1.Body.StripePaymentMethodID = "pm_test_card1"
	input1.Body.IsDefault = true
	repo.CreateGuardianPaymentMethod(ctx, input1)

	input2 := &models.CreateGuardianPaymentMethodInput{}
	input2.Body.GuardianID = guardianID
	input2.Body.StripePaymentMethodID = "pm_test_card2"
	input2.Body.IsDefault = false
	card2, _ := repo.CreateGuardianPaymentMethod(ctx, input2)

	input3 := &models.CreateGuardianPaymentMethodInput{}
	input3.Body.GuardianID = guardianID
	input3.Body.StripePaymentMethodID = "pm_test_card3"
	input3.Body.IsDefault = false
	card3, _ := repo.CreateGuardianPaymentMethod(ctx, input3)

	repo.UpdateGuardianPaymentMethod(ctx, card2.ID, true)

	repo.UpdateGuardianPaymentMethod(ctx, card3.ID, true)

	paymentMethods, err := repo.GetPaymentMethodsByGuardianID(ctx, guardianID)
	require.Nil(t, err)

	defaultCount := 0
	for _, pm := range paymentMethods {
		if pm.IsDefault {
			defaultCount++
			assert.Equal(t, card3.ID, pm.ID)
		}
	}
	assert.Equal(t, 1, defaultCount)
}