package review

import (
	"context"
	"fmt"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/google/uuid"
)

func (r *ReviewRepository) GetEventReviewsForOrganization(
	ctx context.Context,
	id uuid.UUID,
	pagination utils.Pagination,
	sortBy string,
) ([]models.SimpleReviewAggregate, error) {

	baseQuery, err := schema.ReadSQLBaseScript(
		"get_event_reviews_for_organization.sql",
		SqlReviewFiles,
	)
	if err != nil {
		e := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &e
	}

	var orderBy string
	switch sortBy {
	case "most_rated":
		orderBy = "total_reviews DESC"
	case "highest":
		orderBy = "average_rating DESC"
	case "lowest":
		orderBy = "average_rating ASC"
	default:
		orderBy = "total_reviews DESC"
	}

	query := fmt.Sprintf(baseQuery, orderBy)

	rows, err := r.db.Query(ctx, query, id, pagination.Limit, pagination.GetOffset())
	if err != nil {
		e := errs.InternalServerError("Failed to get reviews: ", err.Error())
		return nil, &e
	}
	defer rows.Close()

	var results []models.SimpleReviewAggregate

	for rows.Next() {
		var agg models.SimpleReviewAggregate

		var event models.Event
		var titleEN, descriptionEN string
		var titleTH, descriptionTH *string

		err := rows.Scan(
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
		if err != nil {
			e := errs.InternalServerError("Failed to scan review aggregate row: ", err.Error())
			return nil, &e
		}

		event.Title = titleEN
		event.Description = descriptionEN

		agg.Event = event
		results = append(results, agg)
	}

	if rows.Err() != nil {
		e := errs.InternalServerError("Row iteration error: ", rows.Err().Error())
		return nil, &e
	}

	return results, nil
}
