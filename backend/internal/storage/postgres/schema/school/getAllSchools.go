package school

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *SchoolRepository) GetAllSchools(ctx context.Context) ([]models.School, *errs.HTTPError) {
	//Get base query
	query, err := schema.ReadSQLBaseScript("school/sql/get_all_schools.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	//Execute query
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		errr := errs.InternalServerError("Failed to fetch all schools: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	//Scan rows into schools
	var schools []models.School
	for rows.Next() {
		var school models.School
		err = rows.Scan(&school.ID, &school.Name, &school.Location.ID, &school.CreatedAt, &school.UpdatedAt)
		if err != nil {
			errr := errs.InternalServerError("Failed to scan school: ", err.Error())
			return nil, &errr
		}
		schools = append(schools, school)
	}
	return schools, nil
}
