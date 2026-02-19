package manager

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *ManagerRepository) DeleteManager(ctx context.Context, id uuid.UUID, tx pgx.Tx) (*models.Manager, error) {
	query, err := schema.ReadSQLBaseScript("delete.sql", SqlManagerFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, query, id)
	} else {
		row = r.db.QueryRow(ctx, query, id)
	}

	var deletedManager models.Manager

	err = row.Scan(&deletedManager.ID, &deletedManager.UserID, &deletedManager.OrganizationID, &deletedManager.Role, &deletedManager.Name, &deletedManager.Email, &deletedManager.Username, &deletedManager.ProfilePictureS3Key, &deletedManager.LanguagePreference, &deletedManager.AuthID, &deletedManager.CreatedAt, &deletedManager.UpdatedAt)

	if err != nil {
		err := errs.InternalServerError("Failed to delete manager: ", err.Error())
		return nil, &err
	}

	return &deletedManager, nil
}
