package review

import (
	"context"
	"fmt"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *ReviewRepository) GetEventReviewsForOrganization(
	ctx context.Context,
	id uuid.UUID,
	pagination utils.Pagination,
	AcceptLanguage string,
	sortBy string,
) ([]models.SimpleReviewAggregate, error) {
	baseQuery, err := schema.ReadSQLBaseScript("get_event_reviews_for_organization.sql", SqlReviewFiles)
	if err != nil {
		e := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &e
	}

	orderBy := map[string]string{
		"most_rated": "total_reviews DESC",
		"highest":    "average_rating DESC",
		"lowest":     "average_rating ASC",
	}[sortBy]
	if orderBy == "" {
		orderBy = "total_reviews DESC"
	}

	rows, err := r.db.Query(ctx, fmt.Sprintf(baseQuery, orderBy), id, pagination.Limit, pagination.GetOffset())
	if err != nil {
		e := errs.InternalServerError("Failed to get reviews: ", err.Error())
		return nil, &e
	}

	results, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.SimpleReviewAggregate, error) {
		var agg models.SimpleReviewAggregate
		var event models.Event
		var titleEN, descriptionEN string
		var titleTH, descriptionTH *string

		err := row.Scan(
			&agg.EventID,
			&agg.TotalReviews,
			&agg.AverageRating,
			&event.ID,
			&titleEN,
			&titleTH,
			&descriptionEN,
			&descriptionTH,
			&event.OrganizationID,
			&event.AgeRangeMin,
			&event.AgeRangeMax,
			&event.Category,
			&event.HeaderImageS3Key,
			&event.CreatedAt,
			&event.UpdatedAt,
		)

		switch AcceptLanguage {
		case "th-TH":
			event.Title = *titleTH
			event.Description = *descriptionTH
		case "en-US":
			event.Title = titleEN
			event.Description = descriptionEN
		}

		agg.Event = event
		return agg, err
	})
	if err != nil {
		e := errs.InternalServerError("Failed to scan review aggregate rows: ", err.Error())
		return nil, &e
	}

	return results, nil
}
