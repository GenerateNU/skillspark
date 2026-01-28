package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"skillspark/internal/auth"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) GuardianLogin(ctx context.Context, input *models.LoginInput) (*models.GuardianLoginOutput, error) {

	res, err := auth.SupabaseLogin(&h.config, input.Body.Email, input.Body.Password)
	fmt.Println(json.Marshal(res))
	if err != nil {
		slog.Error(fmt.Sprintf("Login Request Failed: %v", err))
		return nil, errs.InternalServerError(fmt.Sprintf("Login Request Failed: %v", err))
	}

	guardian, err := h.guardianRepository.GetGuardianByAuthID(ctx, res.User.ID.String())
	if err != nil {
		slog.Error(fmt.Sprintf("Could not find associated guardian: %v", err))
		return nil, errs.InternalServerError(fmt.Sprintf("Could not find associated guardian: %v", err))
	}

	

	guardianOutput := &models.GuardianLoginOutput{}

	guardianOutput.ExpiresIn = res.ExpiresIn
	guardianOutput.RefreshToken = res.RefreshToken
	guardianOutput.TokenType = res.TokenType
	guardianOutput.AccessToken = res.AccessToken
	guardianOutput.GuardianID = guardian.ID

	return guardianOutput, nil
}

func (h *Handler) ManagerLogin(ctx context.Context, input *models.LoginInput) (*models.ManagerLoginOutput, error) {

	res, err := auth.SupabaseLogin(&h.config, input.Body.Email, input.Body.Password)
	fmt.Println(json.Marshal(res))
	if err != nil {
		slog.Error(fmt.Sprintf("Login Request Failed: %v", err))
		return nil, errs.InternalServerError(fmt.Sprintf("Login Request Failed: %v", err))
	}

	manager, err := h.managerRepository.GetManagerByAuthID(ctx, res.User.ID.String())
	if err != nil {
		slog.Error(fmt.Sprintf("Could not find associated guardian: %v", err))
		return nil, errs.InternalServerError(fmt.Sprintf("Could not find associated guardian: %v", err))
	}

	managerOutput := &models.ManagerLoginOutput{}

	managerOutput.ExpiresIn = res.ExpiresIn
	managerOutput.RefreshToken = res.RefreshToken
	managerOutput.TokenType = res.TokenType
	managerOutput.AccessToken = res.AccessToken
	managerOutput.ManagerID = manager.ID

	return managerOutput, nil
}