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

func (h *Handler) GuardianLogin(ctx context.Context, input *models.LoginInput) (*models.LoginOutput, error) {

	res, err := auth.SupabaseLogin(&h.config, input.Body.Email, input.Body.Password)
	fmt.Println(json.Marshal(res))
	if err != nil {
		slog.Error(fmt.Sprintf("Login Request Failed: %v", err))
		return nil, errs.InternalServerError(fmt.Sprintf("Login Request Failed: %v", err))
	}

	guardianOutput := &models.LoginOutput{}

	mockToken := "mock-token"

	guardianOutput.Body.Token = mockToken
	// guardianOutput.Body.GuardianID = signInResponse.User.ID

	return guardianOutput, nil
}