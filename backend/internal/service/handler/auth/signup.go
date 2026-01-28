package auth

import (
	"context"
	"fmt"
	"log/slog"
	"skillspark/internal/auth"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) GuardianSignUp(ctx context.Context, input *models.GuardianSignUpInput) (*models.GuardianSignUpOutput, error) {

	// create user in supabase
	res, err := auth.SupabaseSignup(&h.config, input.Body.Email, input.Body.Password)
	fmt.Println(res)
	if err != nil {
		slog.Error(fmt.Sprintf("Signup Request Failed: %v", err))
		return nil, errs.InternalServerError(fmt.Sprintf("Signup Request Failed: %v", err))
	}

	// create user in our database
	user, err := h.userRepository.CreateUser(ctx, func() *models.CreateUserInput {
		userToCreate := &models.CreateUserInput{}
		userToCreate.Body.AuthID = &res.User.ID
		userToCreate.Body.Email = input.Body.Email
		userToCreate.Body.Name = input.Body.Name
		userToCreate.Body.LanguagePreference = input.Body.LanguagePreference
		userToCreate.Body.Username = input.Body.Username
		userToCreate.Body.ProfilePictureS3Key = input.Body.ProfilePictureS3Key
		return userToCreate
	}())
	if err != nil {
		// Cleanup: delete the Supabase auth user since database user creation failed
		if deleteErr := auth.SupabaseDeleteUser(&h.config, res.User.ID); deleteErr != nil {
			slog.Error(fmt.Sprintf("Failed to cleanup Supabase user after database user creation failure: %v", deleteErr))
			return nil, errs.InternalServerError(fmt.Sprintf("Creating User failed: %v. WARNING: Supabase auth user (ID: %s) is stranded and needs manual cleanup", err, res.User.ID))
		}
		return nil, errs.BadRequest(fmt.Sprintf("Creating User failed: %v", err))
	}

	// attach it to guardian entity
	createGuardian := func() *models.CreateGuardianInput {
		guardian := &models.CreateGuardianInput{}
		guardian.Body.UserID = user.ID

		return guardian
	}
	guardian, err := h.guardianRepository.CreateGuardian(ctx, createGuardian())
	if err != nil {
		// Cleanup: delete both the database user and Supabase auth user
		var cleanupErrors []string

		// Delete user from database
		if _, deleteErr := h.userRepository.DeleteUser(ctx, user.ID); deleteErr != nil {
			slog.Error(fmt.Sprintf("Failed to cleanup database user after guardian creation failure: %v", deleteErr))
			cleanupErrors = append(cleanupErrors, fmt.Sprintf("User (ID: %s)", user.ID))
		}

		// Delete Supabase auth user
		if deleteErr := auth.SupabaseDeleteUser(&h.config, res.User.ID); deleteErr != nil {
			slog.Error(fmt.Sprintf("Failed to cleanup Supabase user after guardian creation failure: %v", deleteErr))
			cleanupErrors = append(cleanupErrors, fmt.Sprintf("Supabase auth user (ID: %s)", res.User.ID))
		}

		if len(cleanupErrors) > 0 {
			return nil, errs.InternalServerError(fmt.Sprintf("Creating Guardian failed: %v. WARNING: The following types are stranded and need manual cleanup: %v", err, cleanupErrors))
		}

		return nil, errs.BadRequest(fmt.Sprintf("Creating Guardian/User failed: %v", err))
	}

	// expiration := time.Now().Add(30 * 24 * time.Hour)

	// ctx.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    "TODO:set-me",
	// 	Expires:  expiration,
	// 	Secure:   true,
	// 	SameSite: "Lax",
	// })

	// ctx.Cookie(&fiber.Cookie{
	// 	Name:     "userID",
	// 	Value:    "TODO:set-me",
	// 	Expires:  expiration,
	// 	Secure:   true,
	// 	SameSite: "Lax",
	// })

	guardianOutput := &models.GuardianSignUpOutput{}

	guardianOutput.Body.Token = res.AccessToken
	guardianOutput.Body.GuardianID = guardian.ID

	return guardianOutput, nil
}

// todo: fix the guardian signup to return the correct value

func (h *Handler) ManagerSignUp(ctx context.Context, input *models.ManagerSignUpInput) (*models.ManagerSignUpOutput, error) {

	res, err := auth.SupabaseSignup(&h.config, input.Body.Email, input.Body.Password)
	fmt.Println(res)
	if err != nil {
		slog.Error(fmt.Sprintf("Signup Request Failed: %v", err))
		return nil, errs.InternalServerError(fmt.Sprintf("Signup Request Failed: %v", err))
	}

	user, err := h.userRepository.CreateUser(ctx, func() *models.CreateUserInput {
		userToCreate := &models.CreateUserInput{}
		userToCreate.Body.AuthID = &res.User.ID
		userToCreate.Body.Email = input.Body.Email
		userToCreate.Body.Name = input.Body.Name
		userToCreate.Body.LanguagePreference = input.Body.LanguagePreference
		userToCreate.Body.Username = input.Body.Username
		userToCreate.Body.ProfilePictureS3Key = input.Body.ProfilePictureS3Key
		return userToCreate
	}())
	if err != nil {
		// Cleanup: delete the Supabase auth user since database user creation failed
		if deleteErr := auth.SupabaseDeleteUser(&h.config, res.User.ID); deleteErr != nil {
			slog.Error(fmt.Sprintf("Failed to cleanup Supabase user after database user creation failure: %v", deleteErr))
			return nil, errs.InternalServerError(fmt.Sprintf("Creating User failed: %v. WARNING: Supabase auth user (ID: %s) is stranded and needs manual cleanup", err, res.User.ID))
		}
		return nil, errs.BadRequest(fmt.Sprintf("Creating User failed: %v", err))
	}

	createManager := func() *models.CreateManagerInput {
		manager := &models.CreateManagerInput{}
		manager.Body.UserID = user.ID
		manager.Body.OrganizationID = input.Body.OrganizationID
		manager.Body.Role = input.Body.Role
		return manager
	}
	manager, err := h.managerRepository.CreateManager(ctx, createManager())
	if err != nil {
		// Cleanup: delete both the database user and Supabase auth user
		var cleanupErrors []string

		// Delete user from database
		if _, deleteErr := h.userRepository.DeleteUser(ctx, user.ID); deleteErr != nil {
			slog.Error(fmt.Sprintf("Failed to cleanup database user after manager creation failure: %v", deleteErr))
			cleanupErrors = append(cleanupErrors, fmt.Sprintf("User (ID: %s)", user.ID))
		}

		// Delete Supabase auth user
		if deleteErr := auth.SupabaseDeleteUser(&h.config, res.User.ID); deleteErr != nil {
			slog.Error(fmt.Sprintf("Failed to cleanup Supabase user after manager creation failure: %v", deleteErr))
			cleanupErrors = append(cleanupErrors, fmt.Sprintf("Supabase auth user (ID: %s)", res.User.ID))
		}

		if len(cleanupErrors) > 0 {
			return nil, errs.InternalServerError(fmt.Sprintf("Creating Manager failed: %v. WARNING: The following types are stranded and need manual cleanup: %v", err, cleanupErrors))
		}

		return nil, errs.BadRequest(fmt.Sprintf("Creating Manager/User failed: %v", err))
	}

	// expiration := time.Now().Add(30 * 24 * time.Hour)

	// ctx.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    "TODO:set-me",
	// 	Expires:  expiration,
	// 	Secure:   true,
	// 	SameSite: "Lax",
	// })

	// ctx.Cookie(&fiber.Cookie{
	// 	Name:     "userID",
	// 	Value:    "TODO:set-me",
	// 	Expires:  expiration,
	// 	Secure:   true,
	// 	SameSite: "Lax",
	// })

	managerOutput := &models.ManagerSignUpOutput{}

	managerOutput.Body.Token = res.AccessToken
	managerOutput.Body.ManagerID = manager.ID

	return managerOutput, nil
}
