package child

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *ChildRepository) CreateChild(ctx context.Context, child *models.CreateChildInput) (*models.Child, error) {

	query, err := schema.ReadSQLBaseScript("create.sql", SqlChildFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		child.Body.Name,
		child.Body.SchoolID,
		child.Body.BirthMonth,
		child.Body.BirthYear,
		child.Body.Interests,
		child.Body.GuardianID,
	)

	var createdChild models.Child

	err = row.Scan(&createdChild.ID, &createdChild.Name, &createdChild.SchoolID, &createdChild.SchoolName, &createdChild.BirthMonth, &createdChild.BirthYear, &createdChild.Interests, &createdChild.GuardianID, &createdChild.CreatedAt, &createdChild.UpdatedAt)

	if err != nil {
		err := errs.InternalServerError("Failed to create child: ", err.Error())
		return nil, &err
	}

	return &createdChild, nil

}
