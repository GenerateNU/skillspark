package user

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query, err := schema.ReadSQLBaseScript("user/sql/delete.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, id)

	var deletedUser models.User

	err = row.Scan(&deletedUser.ID, &deletedUser.Name, &deletedUser.Email, &deletedUser.Username, &deletedUser.ProfilePictureS3Key, &deletedUser.LanguagePreference, &deletedUser.AuthID, &deletedUser.CreatedAt, &deletedUser.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to delete user: ", err.Error())
		return nil, &err
	}

	return &deletedUser, nil
}
