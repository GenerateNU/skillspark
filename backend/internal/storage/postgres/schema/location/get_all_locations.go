package location

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/jackc/pgx/v5"
)

func (r *LocationRepository) GetAllLocations(ctx context.Context, pagination utils.Pagination) ([]models.Location, error) {
	//Get base query
	query, err := schema.ReadSQLBaseScript("location/sql/get_all_locations.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	//Execute query
	rows, err := r.db.Query(ctx, query, pagination.Limit, pagination.GetOffset())
	if err != nil {
		errr := errs.InternalServerError("Failed to fetch all locations: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	//Collect the query results
	locations, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Location])
	if err != nil {
		errr := errs.InternalServerError("Failed to scan location: ", err.Error())
		return nil, &errr
	}
	return locations, nil
}
