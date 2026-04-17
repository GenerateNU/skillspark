package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"skillspark/internal/config"
)

type forgotPasswordPayload struct {
	Email      string `json:"email"`
	RedirectTo string `json:"redirect_to"`
}

// SupabaseForgotPassword triggers a password reset email via Supabase.
// Always returns nil to avoid leaking whether an email is registered.
func SupabaseForgotPassword(cfg *config.Supabase, email string, redirectTo string) error {
	payload := forgotPasswordPayload{
		Email:      email,
		RedirectTo: redirectTo,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/v1/recover", cfg.URL), bytes.NewBuffer(payloadBytes))
	if err != nil {
		slog.Error("Error creating forgot password request", "err", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", cfg.AnonKey)

	res, err := Client.Do(req)
	if err != nil {
		slog.Error("Error executing forgot password request", "err", err)
		return err
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		slog.Error("Supabase forgot password error", "status", res.StatusCode, "body", string(body))
	}

	return nil
}
