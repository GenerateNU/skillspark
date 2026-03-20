package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *ReviewRepository) GetAggregateReviews(ctx context.Context, id uuid.UUID) (*models.ReviewAggregate, error) {

	baseQuery, err := schema.ReadSQLBaseScript("get_aggregate_reviews.sql", SqlReviewFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, baseQuery, id)
	if err != nil {
		errr := errs.InternalServerError("Failed to get reviews: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	aggregate := &models.ReviewAggregate{EventID: id}
	var totalWeighted int

	for rows.Next() {
		var rc models.ReviewRatingCount
		if err := rows.Scan(&rc.Rating, &rc.ReviewCount); err != nil {
			return nil, err
		}
		aggregate.Breakdown = append(aggregate.Breakdown, rc)
		aggregate.TotalReviews += rc.ReviewCount
		totalWeighted += int(rc.Rating) * rc.ReviewCount
	}

	if aggregate.TotalReviews > 0 {
		aggregate.AverageRating = float64(totalWeighted) / float64(aggregate.TotalReviews)
	}

	return aggregate, nil
}
