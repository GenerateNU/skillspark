package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var language string

func (r *ReviewRepository) GetReviewsByGuardianID(ctx context.Context, id uuid.UUID, AcceptLanguage string, pagination utils.Pagination) ([]models.Review, error) {

	language = AcceptLanguage
	baseQuery, err := schema.ReadSQLBaseScript("review/sql/get_by_guardian_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, baseQuery, id, pagination.Limit, pagination.GetOffset())
	if err != nil {
		errr := errs.InternalServerError("Failed to get reviews: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	reviews, err := pgx.CollectRows(rows, ScanReviews)
	if err != nil {
		errr := errs.InternalServerError("Failed to collect reviews: ", err.Error())
		return nil, &errr
	}

	return reviews, nil
}

func ScanReviews(row pgx.CollectableRow) (models.Review, error) {
	var review models.GetReviewInput
	var output models.Review
	var desc *string

	err := row.Scan(&review.ID, &review.RegistrationID, &review.GuardianID, &review.Description_EN, &review.Description_TH, &review.Categories, &review.CreatedAt, &review.UpdatedAt)

	if language == "th" {
		desc = review.Description_TH
	} else {
		desc = &review.Description_EN
	}

	output = models.Review{
		ID:             review.ID,
		RegistrationID: review.RegistrationID,
		GuardianID:     review.GuardianID,
		Description:    *desc,
		Categories:     review.Categories,
		CreatedAt:      review.CreatedAt,
		UpdatedAt:      review.UpdatedAt,
	}

	return output, err
}
