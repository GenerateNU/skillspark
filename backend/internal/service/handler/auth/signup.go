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
