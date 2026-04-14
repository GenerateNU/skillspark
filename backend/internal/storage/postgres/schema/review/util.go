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
	desc_ptr := "รีวิวการทดสอบ ให้คะแนนเต็มสิบ"
	input := &models.CreateReviewDBInput{}
	input.Body.RegistrationID = r.ID
	input.Body.GuardianID = &g.ID
	input.Body.Rating = 3
	input.Body.Description_EN = "Test review, ten out of ten"
	input.Body.Description_TH = &desc_ptr
	input.Body.Categories = []string{"interesting", "informative"}

	review, err := repo.CreateReview(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, r)

	return review
}

func buildReviewOrderBy(sortBy string) string {
	switch sortBy {
	case "highest_rating":
		return "r.rating DESC, r.created_at DESC"
	case "lowest_rating":
		return "r.rating ASC, r.created_at DESC"
	default:
		return "r.created_at DESC"
	}
}
