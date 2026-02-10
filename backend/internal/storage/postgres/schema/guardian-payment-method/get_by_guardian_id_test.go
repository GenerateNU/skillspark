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

func TestGetPaymentMethodsByGuardianID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianPaymentMethodRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianID := uuid.MustParse("88888888-8888-8888-8888-888888888888")

	cardBrand1 := "visa"
	cardLast1 := "4242"
	input1 := &models.CreateGuardianPaymentMethodInput{}
	input1.Body.GuardianID = guardianID
	input1.Body.StripePaymentMethodID = "pm_test_visa"
	input1.Body.CardBrand = &cardBrand1
	input1.Body.CardLast4 = &cardLast1
	input1.Body.IsDefault = true

	created1, err := repo.CreateGuardianPaymentMethod(ctx, input1)
	require.Nil(t, err)

	cardBrand2 := "mastercard"
	cardLast2 := "5555"
	input2 := &models.CreateGuardianPaymentMethodInput{}
	input2.Body.GuardianID = guardianID
	input2.Body.StripePaymentMethodID = "pm_test_mastercard"
	input2.Body.CardBrand = &cardBrand2
	input2.Body.CardLast4 = &cardLast2
	input2.Body.IsDefault = false

	created2, err := repo.CreateGuardianPaymentMethod(ctx, input2)
	require.Nil(t, err)

	paymentMethods, err := repo.GetPaymentMethodsByGuardianID(ctx, guardianID)

	require.Nil(t, err)
	require.NotNil(t, paymentMethods)
	assert.GreaterOrEqual(t, len(paymentMethods), 2)

	foundVisa := false
	foundMastercard := false
	for _, pm := range paymentMethods {
		if pm.ID == created1.ID {
			foundVisa = true
			assert.Equal(t, "pm_test_visa", pm.StripePaymentMethodID)
			assert.Equal(t, "visa", *pm.CardBrand)
			assert.True(t, pm.IsDefault)
		}
		if pm.ID == created2.ID {
			foundMastercard = true
			assert.Equal(t, "pm_test_mastercard", pm.StripePaymentMethodID)
			assert.Equal(t, "mastercard", *pm.CardBrand)
			assert.False(t, pm.IsDefault)
		}
	}
	assert.True(t, foundVisa)
	assert.True(t, foundMastercard)
}

func TestGetPaymentMethodsByGuardianID_Empty(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianPaymentMethodRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianID := uuid.MustParse("88888888-8888-8888-8888-888888888889")

	paymentMethods, err := repo.GetPaymentMethodsByGuardianID(ctx, guardianID)

	require.Nil(t, err)
	require.NotNil(t, paymentMethods)
	assert.Equal(t, 0, len(paymentMethods))
}

func TestGetPaymentMethodsByGuardianID_DefaultFirst(t *testing.T) {
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
	input1.Body.StripePaymentMethodID = "pm_test_non_default"
	input1.Body.IsDefault = false
	repo.CreateGuardianPaymentMethod(ctx, input1)

	input2 := &models.CreateGuardianPaymentMethodInput{}
	input2.Body.GuardianID = guardianID
	input2.Body.StripePaymentMethodID = "pm_test_default"
	input2.Body.IsDefault = true
	repo.CreateGuardianPaymentMethod(ctx, input2)

	paymentMethods, err := repo.GetPaymentMethodsByGuardianID(ctx, guardianID)

	require.Nil(t, err)
	assert.GreaterOrEqual(t, len(paymentMethods), 2)

	if len(paymentMethods) >= 2 {
		assert.True(t, paymentMethods[0].IsDefault)
	}
}

func TestGetPaymentMethodsByGuardianID_OnlyReturnsOwnMethods(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianPaymentMethodRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardian1ID := uuid.MustParse("88888888-8888-8888-8888-888888888888")
	guardian2ID := uuid.MustParse("88888888-8888-8888-8888-888888888889")

	input1 := &models.CreateGuardianPaymentMethodInput{}
	input1.Body.GuardianID = guardian1ID
	input1.Body.StripePaymentMethodID = "pm_guardian1_card"
	input1.Body.IsDefault = true
	repo.CreateGuardianPaymentMethod(ctx, input1)

	input2 := &models.CreateGuardianPaymentMethodInput{}
	input2.Body.GuardianID = guardian2ID
	input2.Body.StripePaymentMethodID = "pm_guardian2_card"
	input2.Body.IsDefault = true
	repo.CreateGuardianPaymentMethod(ctx, input2)

	paymentMethods, err := repo.GetPaymentMethodsByGuardianID(ctx, guardian1ID)

	require.Nil(t, err)
	for _, pm := range paymentMethods {
		assert.Equal(t, guardian1ID, pm.GuardianID)
		assert.NotEqual(t, "pm_guardian2_card", pm.StripePaymentMethodID)
	}
}