package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *OrganizationRepository) GetAllOrganizations(ctx context.Context, offset, pageSize int) ([]models.Organization, int, *errs.HTTPError) {
	
	baseQuery, err := schema.ReadSQLBaseScript("organization/sql/get_all.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, 0, &errr
	}

	countQuery, err := schema.ReadSQLBaseScript("organization/sql/count_all.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read count query: ", err.Error())
		return nil, 0, &errr
	}

	// Get total count
	var totalCount int
	err = r.db.QueryRow(ctx, countQuery).Scan(&totalCount)
	if err != nil {
		errr := errs.InternalServerError("Failed to count organizations: ", err.Error())
		return nil, 0, &errr
	}

	// Execute query with pagination
	rows, err := r.db.Query(ctx, baseQuery, pageSize, offset)
	if err != nil {
		errr := errs.InternalServerError("Failed to get organizations: ", err.Error())
		return nil, 0, &errr
	}
	defer rows.Close()

	orgs := []models.Organization{}
	
	for rows.Next() {
		var org models.Organization
		err := rows.Scan(
			&org.ID,
			&org.Name,
			&org.Active,
			&org.PfpS3Key,
			&org.LocationID,
			&org.CreatedAt,
			&org.UpdatedAt,
		)
		if err != nil {
			errr := errs.InternalServerError("Failed to scan organization: ", err.Error())
			return nil, 0, &errr
		}
		orgs = append(orgs, org)
	}

	return orgs, totalCount, nil
}