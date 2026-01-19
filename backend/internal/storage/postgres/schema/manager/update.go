package manager

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *ManagerRepository) PatchManager(ctx context.Context, manager *models.PatchManagerInput) (*models.Manager, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("location/sql/update.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, manager.Body.UserID, manager.Body.OrganizationID, manager.Body.Role)

	var createdManager models.Manager

	err = row.Scan(&createdManager.ID, &createdManager.UserID, &createdManager.OrganizationID, &createdManager.Role, &createdManager.CreatedAt, &createdManager.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to update location: ", err.Error())
		return nil, &err
	}

	return &createdManager, nil
}
