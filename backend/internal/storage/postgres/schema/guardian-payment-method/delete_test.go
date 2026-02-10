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

func TestDeleteGuardianPaymentMethod(t *testing.T) {
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
	input.Body.StripePaymentMethodID = "pm_test_delete"
	input.Body.IsDefault = false

	created, createErr := repo.CreateGuardianPaymentMethod(ctx, input)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	deleted, deleteErr := repo.DeleteGuardianPaymentMethod(ctx, created.ID)

	require.Nil(t, deleteErr)
	require.NotNil(t, deleted)
	assert.Equal(t, created.ID, deleted.ID)
	assert.Equal(t, "pm_test_delete", deleted.StripePaymentMethodID)
	assert.Equal(t, guardianID, deleted.GuardianID)

	paymentMethods, err := repo.GetPaymentMethodsByGuardianID(ctx, guardianID)
	require.Nil(t, err)

	for _, pm := range paymentMethods {
		assert.NotEqual(t, created.ID, pm.ID)
	}
}

func TestDeleteGuardianPaymentMethod_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianPaymentMethodRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	deleted, err := repo.DeleteGuardianPaymentMethod(ctx, uuid.New())

	require.NotNil(t, err)
	assert.Nil(t, deleted)
}

func TestDeleteGuardianPaymentMethod_AlreadyDeleted(t *testing.T) {
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
	input.Body.StripePaymentMethodID = "pm_test_double_delete"
	input.Body.IsDefault = false

	created, createErr := repo.CreateGuardianPaymentMethod(ctx, input)
	require.Nil(t, createErr)

	deleted1, deleteErr1 := repo.DeleteGuardianPaymentMethod(ctx, created.ID)
	require.Nil(t, deleteErr1)
	require.NotNil(t, deleted1)

	deleted2, deleteErr2 := repo.DeleteGuardianPaymentMethod(ctx, created.ID)
	require.NotNil(t, deleteErr2)
	assert.Nil(t, deleted2)
}

func TestDeleteGuardianPaymentMethod_DefaultCard(t *testing.T) {
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
	input.Body.StripePaymentMethodID = "pm_test_delete_default"
	input.Body.IsDefault = true

	created, createErr := repo.CreateGuardianPaymentMethod(ctx, input)
	require.Nil(t, createErr)
	require.True(t, created.IsDefault)

	deleted, deleteErr := repo.DeleteGuardianPaymentMethod(ctx, created.ID)

	require.Nil(t, deleteErr)
	require.NotNil(t, deleted)
	assert.True(t, deleted.IsDefault)
}