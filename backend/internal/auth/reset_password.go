package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"skillspark/internal/config"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

type updatePasswordPayload struct {
	Password string `json:"password"`
}

// SupabaseResetPassword updates the user's password using the access token
// from the password reset email redirect.
func SupabaseResetPassword(cfg *config.Supabase, accessToken string, newPassword string) error {
	if err := validatePasswordStrength(newPassword); err != nil {
		return err
	}

	payload := updatePasswordPayload{Password: newPassword}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/auth/v1/user", cfg.URL), bytes.NewBuffer(payloadBytes))
	if err != nil {
		slog.Error("Error creating reset password request", "err", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", cfg.AnonKey)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	res, err := Client.Do(req)
	if err != nil {
		slog.Error("Error executing reset password request", "err", err)
		return err
	}
	defer func() { _ = res.Body.Close() }()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		supabaseError := &models.SupabaseError{}
		if jsonErr := json.Unmarshal(body, supabaseError); jsonErr != nil {
			slog.Error("Error parsing reset password response", "err", jsonErr)
			return errs.BadRequest("Failed to reset password")
		}
		slog.Error("Supabase reset password error", "status", res.StatusCode, "body", string(body))
		return errs.NewHTTPError(res.StatusCode, supabaseError)
	}

	return nil
}
