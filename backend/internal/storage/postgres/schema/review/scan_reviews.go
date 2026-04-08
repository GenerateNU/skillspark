package review

import (
	"skillspark/internal/models"

	"github.com/jackc/pgx/v5"
)

func (r *ReviewRepository) ScanReviews(row pgx.CollectableRow) (models.Review, error) {
	var review models.GetReviewInput
	var output models.Review
	var desc *string

	err := row.Scan(&review.ID, &review.RegistrationID, &review.GuardianID, &review.EventID, &review.Rating, &review.Description_EN, &review.Description_TH, &review.Categories, &review.CreatedAt, &review.UpdatedAt)

	if language == "th" {
		desc = review.Description_TH
	} else {
		desc = &review.Description_EN
	}

	output = models.Review{
		ID:             review.ID,
		RegistrationID: review.RegistrationID,
		GuardianID:     review.GuardianID, // *uuid.UUID — nil for anonymous reviews
		EventID:        review.EventID,
		Rating:         review.Rating,
		Description:    *desc,
		Categories:     review.Categories,
		CreatedAt:      review.CreatedAt,
		UpdatedAt:      review.UpdatedAt,
	}

	return output, err
}
