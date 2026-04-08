package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *ReviewRepository) GetAggregateReviewsForOrganization(ctx context.Context, id uuid.UUID) (*models.ReviewAggregate, error) {
	baseQuery, err := schema.ReadSQLBaseScript("get_aggregate_reviews_for_organization.sql", SqlReviewFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, baseQuery, id)
	if err != nil {
		errr := errs.InternalServerError("Failed to get reviews: ", err.Error())
		return nil, &errr
	}

	ratingCounts, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.ReviewRatingCount])
	if err != nil {
		errr := errs.InternalServerError("Failed to collect review rows: ", err.Error())
		return nil, &errr
	}

	aggregate := &models.ReviewAggregate{EventID: id}
	var totalWeighted int

	for _, rc := range ratingCounts {
		aggregate.Breakdown = append(aggregate.Breakdown, rc)
		aggregate.TotalReviews += rc.ReviewCount
		totalWeighted += int(rc.Rating) * rc.ReviewCount
	}

	if aggregate.TotalReviews > 0 {
		aggregate.AverageRating = float64(totalWeighted) / float64(aggregate.TotalReviews)
	}

	return aggregate, nil
}
