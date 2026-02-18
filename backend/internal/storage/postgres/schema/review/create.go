package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *ReviewRepository) CreateReview(ctx context.Context, input *models.CreateReviewDBInput) (*models.Review, error) {
	query, err := schema.ReadSQLBaseScript("review/sql/create.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		input.Body.RegistrationID,
		input.Body.GuardianID,
		input.Body.Description_EN,
		input.Body.Description_TH,
		input.Body.Categories,
	)

	var createdReview models.Review

	err = row.Scan(&createdReview.ID, &createdReview.RegistrationID, &createdReview.GuardianID, &createdReview.Description, &createdReview.Categories, &createdReview.CreatedAt, &createdReview.UpdatedAt)

	if err != nil {
		err := errs.InternalServerError("Failed to create child: ", err.Error())
		return nil, &err
	}

	return &createdReview, nil

}
