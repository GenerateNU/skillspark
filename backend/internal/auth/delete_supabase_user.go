package auth

import (
	"io"
	"net/http"
	"skillspark/internal/config"
	"log/slog"
	"github.com/google/uuid"
	"fmt"
)

// deletes a user from Supabase auth by their ID
func SupabaseDeleteUser(cfg *config.Supabase, userID uuid.UUID) error {
	supabaseURL := cfg.URL
	apiKey := cfg.ServiceRoleKey

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/auth/v1/admin/users/%s", supabaseURL, userID.String()), nil)
	if err != nil {
		slog.Error("Error in Request Creation: ", "err", err)
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("apikey", apiKey)

	res, err := Client.Do(req)
	if err != nil {
		slog.Error("Error executing delete request: ", "err", err)
		return err
	}

	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(res.Body)
		slog.Error("Error deleting user from Supabase: ", "res.StatusCode", res.StatusCode, "body", string(body))
		return err
	}

	return nil
}
