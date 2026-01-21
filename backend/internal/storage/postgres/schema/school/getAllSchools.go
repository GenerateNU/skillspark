package school

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/jackc/pgx/v5"
)

func (r *SchoolRepository) GetAllSchools(ctx context.Context, pagination utils.Pagination) ([]models.School, error) {
	//Get base query
	query, err := schema.ReadSQLBaseScript("school/sql/get_all_schools.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	//Execute query
	rows, err := r.db.Query(ctx, query, pagination.Limit, pagination.GetOffset())
	if err != nil {
		errr := errs.InternalServerError("Failed to fetch all schools: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	//Collect the query results
	schools, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.School])
	if err != nil {
		errr := errs.InternalServerError("Failed to scan school: ", err.Error())
		return nil, &errr
	}
	return schools, nil
}
