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
	sortBy string,
) ([]models.SimpleReviewAggregate, error) {

	baseQuery, err := schema.ReadSQLBaseScript("get_event_reviews_for_organization.sql", SqlReviewFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
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
		errr := errs.InternalServerError("Failed to get reviews: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	aggregates, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.SimpleReviewAggregate])
	if err != nil {
		errr := errs.InternalServerError("Failed to collect review rows: ", err.Error())
		return nil, &errr
	}

	return aggregates, nil
}
