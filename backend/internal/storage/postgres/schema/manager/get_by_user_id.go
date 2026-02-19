package manager

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *ManagerRepository) GetManagerByUserID(ctx context.Context, userID uuid.UUID) (*models.Manager, error) {
	query, err := schema.ReadSQLBaseScript("get_by_user_id.sql", SqlManagerFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, userID)
	var manager models.Manager
	err = row.Scan(&manager.ID, &manager.UserID, &manager.OrganizationID, &manager.Role, &manager.Name, &manager.Email, &manager.Username, &manager.ProfilePictureS3Key, &manager.LanguagePreference,
		&manager.CreatedAt, &manager.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Manager", "user_id", userID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch manager by user id: ", err.Error())
		return nil, &err
	}

	return &manager, nil
}
