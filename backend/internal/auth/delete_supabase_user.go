package auth

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"skillspark/internal/config"
	"skillspark/internal/errs"

	"github.com/google/uuid"
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
		return errs.InternalServerError("Error deleting user from Supabase: " + string(body))
	}

	return nil
}
