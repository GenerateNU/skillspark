package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"skillspark/internal/auth"
	"skillspark/internal/models"
)

func (h *Handler) GuardianLogin(ctx context.Context, input *models.LoginInput) (*models.GuardianLoginOutput, error) {

	res, err := auth.SupabaseLogin(&h.config, input.Body.Email, input.Body.Password)
	fmt.Println(json.Marshal(res))
	if err != nil {
		slog.Error(fmt.Sprintf("Login Request Failed: %v", err))
		return nil, err
	}

	guardian, err := h.guardianRepository.GetGuardianByAuthID(ctx, res.User.ID.String())
	if err != nil {
		slog.Error(fmt.Sprintf("Could not find associated guardian: %v", err))
		return nil, err
	}

	guardianOutput := &models.GuardianLoginOutput{}

	guardianOutput.Body.GuardianID = guardian.ID

	guardianOutput.AccessTokenCookie = http.Cookie{
		Name:     "jwt",
		Value:    res.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // HTTPS only (set false for localhost dev)
		SameSite: http.SameSiteStrictMode,
		MaxAge:   res.ExpiresIn, // short-lived (e.g., 1 hour)
	}

	return guardianOutput, nil
}

func (h *Handler) ManagerLogin(ctx context.Context, input *models.LoginInput) (*models.ManagerLoginOutput, error) {

	res, err := auth.SupabaseLogin(&h.config, input.Body.Email, input.Body.Password)

	if err != nil {
		slog.Error(fmt.Sprintf("Login Request Failed: %v", err))
		return nil, err
	}

	manager, err := h.managerRepository.GetManagerByAuthID(ctx, res.User.ID.String())
	if err != nil {
		slog.Error(fmt.Sprintf("Could not find associated guardian: %v", err))
		return nil, err
	}

	managerOutput := &models.ManagerLoginOutput{}

	managerOutput.Body.ManagerID = manager.ID

	managerOutput.AccessTokenCookie = http.Cookie{
		Name:     "jwt",
		Value:    res.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // HTTPS only (set false for localhost dev)
		SameSite: http.SameSiteStrictMode,
		MaxAge:   res.ExpiresIn, // short-lived (e.g., 1 hour)
	}

	return managerOutput, nil

}
