package user

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query, err := schema.ReadSQLBaseScript("get_by_username.sql", SqlUserFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, username)

	var user models.User

	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.ProfilePictureS3Key, &user.LanguagePreference, &user.AuthID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			err := errs.NotFound("User", "username", username)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to get user: ", err.Error())
		return nil, &err
	}

	return &user, nil
}
