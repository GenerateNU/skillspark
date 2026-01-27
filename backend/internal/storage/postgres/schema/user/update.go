package user

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.UpdateUserInput) (*models.User, error) {
	query, err := schema.ReadSQLBaseScript("user/sql/update.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	// Get existing user first
	existing, err := r.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	// Merge optional fields
	if user.Body.Name != nil {
		existing.Name = *user.Body.Name
	}
	if user.Body.Email != nil {
		existing.Email = *user.Body.Email
	}
	if user.Body.Username != nil {
		existing.Username = *user.Body.Username
	}
	if user.Body.ProfilePictureS3Key != nil {
		existing.ProfilePictureS3Key = user.Body.ProfilePictureS3Key
	}
	if user.Body.LanguagePreference != nil {
		existing.LanguagePreference = *user.Body.LanguagePreference
	}
	if user.Body.AuthID != nil {
		existing.AuthID = user.Body.AuthID
	}

	row := r.db.QueryRow(ctx, query, existing.Name, existing.Email, existing.Username, existing.ProfilePictureS3Key, existing.LanguagePreference, existing.AuthID, existing.ID)

	var updatedUser models.User

	err = row.Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email, &updatedUser.Username, &updatedUser.ProfilePictureS3Key, &updatedUser.LanguagePreference, &updatedUser.AuthID, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to update user: ", err.Error())
		return nil, &err
	}

	return &updatedUser, nil
}
