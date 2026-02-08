package manager

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *ManagerRepository) CreateManager(ctx context.Context, manager *models.CreateManagerInput) (*models.Manager, error) {
	query, err := schema.ReadSQLBaseScript("manager/sql/create.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, manager.Body.Name, manager.Body.Email, manager.Body.Username, manager.Body.ProfilePictureS3Key, manager.Body.LanguagePreference, manager.Body.AuthID, manager.Body.OrganizationID, manager.Body.Role)

	var createdManager models.Manager

	err = row.Scan(&createdManager.ID, &createdManager.UserID, &createdManager.OrganizationID, &createdManager.Role, &createdManager.Name, &createdManager.Email, &createdManager.Username, &createdManager.ProfilePictureS3Key, &createdManager.LanguagePreference, &createdManager.AuthID, &createdManager.CreatedAt, &createdManager.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to create manager: ", err.Error())
		return nil, &err
	}

	return &createdManager, nil
}
