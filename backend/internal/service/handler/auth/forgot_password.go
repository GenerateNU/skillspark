package auth

import (
	"context"
	"fmt"
	"skillspark/internal/auth"
	"skillspark/internal/models"
)

func (h *Handler) ForgotPassword(_ context.Context, input *models.ForgotPasswordInput) (*models.ForgotPasswordOutput, error) {
	redirectTo := fmt.Sprintf("%s/reset-password", h.appConfig.FrontendURL)

	// Ignore the error intentionally to avoid leaking whether an email is registered.
	_ = auth.SupabaseForgotPassword(&h.config, input.Body.Email, redirectTo)

	return &models.ForgotPasswordOutput{
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "If this email is registered, you will receive password reset instructions.",
		},
	}, nil
}

func (h *Handler) ResetPassword(_ context.Context, input *models.ResetPasswordInput) (*models.ResetPasswordOutput, error) {
	if err := auth.SupabaseResetPassword(&h.config, input.Body.AccessToken, input.Body.NewPassword); err != nil {
		return nil, err
	}

	return &models.ResetPasswordOutput{
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "Password updated successfully.",
		},
	}, nil
}
