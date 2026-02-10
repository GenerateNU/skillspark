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

func TestCreateGuardianPaymentMethod(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianPaymentMethodRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianID := uuid.MustParse("88888888-8888-8888-8888-888888888888")
	cardBrand := "visa"
	cardLast4 := "4242"
	expMonth := 12
	expYear := 2027

	input := &models.CreateGuardianPaymentMethodInput{}
	input.Body.GuardianID = guardianID
	input.Body.StripePaymentMethodID = "pm_test_visa_4242"
	input.Body.CardBrand = &cardBrand
	input.Body.CardLast4 = &cardLast4
	input.Body.CardExpMonth = &expMonth
	input.Body.CardExpYear = &expYear
	input.Body.IsDefault = true

	created, err := repo.CreateGuardianPaymentMethod(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.NotEqual(t, uuid.Nil, created.ID)
	assert.Equal(t, guardianID, created.GuardianID)
	assert.Equal(t, "pm_test_visa_4242", created.StripePaymentMethodID)
	assert.Equal(t, "visa", *created.CardBrand)
	assert.Equal(t, "4242", *created.CardLast4)
	assert.Equal(t, 12, *created.CardExpMonth)
	assert.Equal(t, 2027, *created.CardExpYear)
	assert.True(t, created.IsDefault)
	assert.NotNil(t, created.CreatedAt)
	assert.NotNil(t, created.UpdatedAt)
}

func TestCreateGuardianPaymentMethod_NonDefault(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewGuardianPaymentMethodRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	guardianID := uuid.MustParse("88888888-8888-8888-8888-888888888888")
	cardBrand := "mastercard"
	cardLast4 := "5555"
	expMonth := 6
	expYear := 2028

	input := &models.CreateGuardianPaymentMethodInput{}
	input.Body.GuardianID = guardianID
	input.Body.StripePaymentMethodID = "pm_test_mastercard_5555"
	input.Body.CardBrand = &cardBrand
	input.Body.CardLast4 = &cardLast4
	input.Body.CardExpMonth = &expMonth
	input.Body.CardExpYear = &expYear
	input.Body.IsDefault = false

	created, err := repo.CreateGuardianPaymentMethod(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.False(t, created.IsDefault)
}

func TestCreateGuardianPaymentMethod_MinimalInfo(t *testing.T) {
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
	input.Body.StripePaymentMethodID = "pm_test_minimal"
	input.Body.IsDefault = false

	created, err := repo.CreateGuardianPaymentMethod(ctx, input)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, "pm_test_minimal", created.StripePaymentMethodID)
	assert.Nil(t, created.CardBrand)
	assert.Nil(t, created.CardLast4)
	assert.Nil(t, created.CardExpMonth)
	assert.Nil(t, created.CardExpYear)
}

func TestCreateGuardianPaymentMethod_MultipleCards(t *testing.T) {
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

	created1, err := repo.CreateGuardianPaymentMethod(ctx, input1)
	require.Nil(t, err)
	assert.True(t, created1.IsDefault)

	input2 := &models.CreateGuardianPaymentMethodInput{}
	input2.Body.GuardianID = guardianID
	input2.Body.StripePaymentMethodID = "pm_test_card2"
	input2.Body.IsDefault = false

	created2, err := repo.CreateGuardianPaymentMethod(ctx, input2)
	require.Nil(t, err)
	assert.False(t, created2.IsDefault)

	paymentMethods, err := repo.GetPaymentMethodsByGuardianID(ctx, guardianID)
	require.Nil(t, err)
	assert.GreaterOrEqual(t, len(paymentMethods), 2)
}