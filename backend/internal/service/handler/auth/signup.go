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
	if err != nil {
		slog.Error(fmt.Sprintf("Signup Request Failed: %v", err))
		return nil, err
	}

	// attach it to guardian entity
	createGuardian := func() *models.CreateGuardianInput {
		guardian := &models.CreateGuardianInput{}
		guardian.Body.Email = input.Body.Email
		guardian.Body.Name = input.Body.Name
		guardian.Body.LanguagePreference = input.Body.LanguagePreference
		guardian.Body.Username = input.Body.Username
		guardian.Body.ProfilePictureS3Key = input.Body.ProfilePictureS3Key

		return guardian
	}
	guardian, err := h.guardianRepository.CreateGuardian(ctx, createGuardian())
	if err != nil {

		// Delete Supabase auth user as cleanup
		if deleteErr := auth.SupabaseDeleteUser(&h.config, res.User.ID); deleteErr != nil {
			slog.Error(fmt.Sprintf("Failed to cleanup Supabase user after guardian creation failure: %v", deleteErr))
			return nil, errs.InternalServerError(fmt.Sprintf("Creating Guardian failed: %v. WARNING: The following types are stranded and need manual cleanup: %v", err, deleteErr))
		}

		return nil, errs.BadRequest(fmt.Sprintf("Creating Guardian/User failed: %v", err))
	}

	_, err = h.userRepository.UpdateUser(ctx, func() *models.UpdateUserInput {
		updateUserInput := &models.UpdateUserInput{}
		updateUserInput.ID = guardian.UserID
		updateUserInput.Body.AuthID = &res.User.ID
		return updateUserInput
	}())
	if err != nil {
		// Cleanup: the Supabase auth user
		if deleteErr := auth.SupabaseDeleteUser(&h.config, res.User.ID); deleteErr != nil {
			slog.Error(fmt.Sprintf("Failed to cleanup Supabase user after guardian creation failure: %v", deleteErr))
			return nil, errs.InternalServerError(fmt.Sprintf("Creating Guardian failed: %v. WARNING: The following types are stranded and need manual cleanup: %v", err, deleteErr))
		}
		return nil, errs.BadRequest(fmt.Sprintf("Creating Guardian/User failed: %v", err))
	}

	guardianOutput := &models.GuardianSignUpOutput{}

	guardianOutput.Body.Token = res.AccessToken
	guardianOutput.Body.GuardianID = guardian.ID

	return guardianOutput, nil
}

func (h *Handler) ManagerSignUp(ctx context.Context, input *models.ManagerSignUpInput) (*models.ManagerSignUpOutput, error) {

	res, err := auth.SupabaseSignup(&h.config, input.Body.Email, input.Body.Password)
	if err != nil {
		slog.Error(fmt.Sprintf("Signup Request Failed: %v", err))
		return nil, err
	}

	createManager := func() *models.CreateManagerInput {
		manager := &models.CreateManagerInput{}
		manager.Body.Name = input.Body.Name
		manager.Body.Email = input.Body.Email
		manager.Body.Username = input.Body.Username
		manager.Body.ProfilePictureS3Key = input.Body.ProfilePictureS3Key
		manager.Body.LanguagePreference = input.Body.LanguagePreference
		manager.Body.OrganizationID = input.Body.OrganizationID
		manager.Body.Role = input.Body.Role
		return manager
	}
	manager, err := h.managerRepository.CreateManager(ctx, createManager())
	if err != nil {
		// Cleanup the Supabase auth user
		if deleteErr := auth.SupabaseDeleteUser(&h.config, res.User.ID); deleteErr != nil {
			slog.Error(fmt.Sprintf("Failed to cleanup Supabase user after manager creation failure: %v", deleteErr))
			return nil, errs.InternalServerError(fmt.Sprintf("Creating Manager failed: %v. WARNING: The following types are stranded and need manual cleanup: %v", err, deleteErr))
		}

		return nil, errs.BadRequest(fmt.Sprintf("Creating Manager/User failed: %v", err))
	}

	// attach auth user to user entity
	_, err = h.userRepository.UpdateUser(ctx, func() *models.UpdateUserInput {
		updateUserInput := &models.UpdateUserInput{}
		updateUserInput.ID = manager.UserID
		updateUserInput.Body.AuthID = &res.User.ID
		return updateUserInput
	}())
	if err != nil {
		// Cleanup the Supabase auth user
		if deleteErr := auth.SupabaseDeleteUser(&h.config, res.User.ID); deleteErr != nil {
			slog.Error(fmt.Sprintf("Failed to cleanup Supabase user after guardian creation failure: %v", deleteErr))
			return nil, errs.InternalServerError(fmt.Sprintf("Creating Guardian failed: %v. WARNING: The following types are stranded and need manual cleanup: %v", err, deleteErr))
		}
		return nil, errs.BadRequest(fmt.Sprintf("Creating Guardian/User failed: %v", err))
	}

	managerOutput := &models.ManagerSignUpOutput{}

	managerOutput.Body.Token = res.AccessToken
	managerOutput.Body.ManagerID = manager.ID

	return managerOutput, nil
}
