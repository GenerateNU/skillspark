package manager

import (
	"context"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *ManagerRepository) PatchManager(ctx context.Context, manager *models.PatchManagerInput) (*models.Manager, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	var updatedManager models.Manager
	updatedManager.ID = manager.Body.ID

	managerQuery, err := schema.ReadSQLBaseScript("manager/sql/update_manager.sql")
	if err != nil {
		_ = tx.Rollback(ctx)
		err := errs.InternalServerError("Failed to read manager update query: ", err.Error())
		return nil, &err
	}

	err = tx.QueryRow(ctx, managerQuery,
		manager.Body.ID,
		manager.Body.OrganizationID,
		manager.Body.Role,
	).Scan(
		&updatedManager.UserID,
		&updatedManager.OrganizationID,
		&updatedManager.Role,
		&updatedManager.CreatedAt,
		&updatedManager.UpdatedAt,
	)

	if err != nil {
		_ = tx.Rollback(ctx)
		err := errs.InternalServerError("Failed to update manager table: ", err.Error())
		return nil, &err
	}

	userQuery, err := schema.ReadSQLBaseScript("manager/sql/update_user.sql")
	if err != nil {
		_ = tx.Rollback(ctx)
		err := errs.InternalServerError("Failed to read user update query: ", err.Error())
		return nil, &err
	}

	err = tx.QueryRow(ctx, userQuery,
		updatedManager.UserID,
		manager.Body.Name,
		manager.Body.Email,
		manager.Body.Username,
		manager.Body.ProfilePictureS3Key,
		manager.Body.LanguagePreference,
	).Scan(
		&updatedManager.Name,
		&updatedManager.Email,
		&updatedManager.Username,
		&updatedManager.ProfilePictureS3Key,
		&updatedManager.LanguagePreference,
	)

	if err != nil {
		_ = tx.Rollback(ctx)
		err := errs.InternalServerError("Failed to update user table: ", err.Error())
		return nil, &err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &updatedManager, nil
}
