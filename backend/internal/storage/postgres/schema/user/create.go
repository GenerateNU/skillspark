package user

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *UserRepository) CreateUser(ctx context.Context, user *models.CreateUserInput) (*models.User, error) {
	query, err := schema.ReadSQLBaseScript("create.sql", SqlUserFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, user.Body.Name, user.Body.Email, user.Body.Username, user.Body.ProfilePictureS3Key, user.Body.LanguagePreference, user.Body.AuthID)

	var createdUser models.User

	err = row.Scan(&createdUser.ID, &createdUser.Name, &createdUser.Email, &createdUser.Username, &createdUser.ProfilePictureS3Key, &createdUser.LanguagePreference, &createdUser.AuthID, &createdUser.CreatedAt, &createdUser.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to create user: ", err.Error())
		return nil, &err
	}

	return &createdUser, nil
}
