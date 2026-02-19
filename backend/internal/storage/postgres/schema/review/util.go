package review

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/schema/registration"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlReviewFiles embed.FS

func CreateTestReview(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Review {

	t.Helper()
	repo := NewReviewRepository(db)

	r := registration.CreateTestRegistration(t, ctx, db)
	g := guardian.CreateTestGuardian(t, ctx, db)

	input := &models.CreateReviewInput{}
	input.Body.RegistrationID = r.ID
	input.Body.GuardianID = g.ID
	input.Body.Description = "Test review, ten out of ten"
	input.Body.Categories = []string{"interesting", "informative"}

	review, err := repo.CreateReview(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, r)

	return review
}
