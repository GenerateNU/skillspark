package school

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *SchoolRepository) CreateSchool(ctx context.Context, school *models.CreateSchoolInput) (*models.School, error) {

	query, err := schema.ReadSQLBaseScript("school/sql/create.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		school.Body.Name,
		school.Body.LocationID,
	)

	var createdSchool models.School

	err = row.Scan(&createdSchool.ID, &createdSchool.Name, &createdSchool.LocationID, &createdSchool.CreatedAt, &createdSchool.UpdatedAt)

	if err != nil {
		err := errs.InternalServerError("Failed to create school: ", err.Error())
		return nil, &err
	}

	return &createdSchool, nil

}
