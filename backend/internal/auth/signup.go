package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"regexp"
	"skillspark/internal/config"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

// checks if password is strong enough
func validatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be atleast 8 characters long")
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#~$%^&*()+|_.,;<>?/{}\-]`).MatchString(password)

	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return errors.New("password must include uppercase, lowercase, digit and special characters")
	}

	return nil
}

type SupabaseError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"msg"`
}

func (e *SupabaseError) Error() string {
	return e.Message
}

// Creates a new user in Supabase auth
func SupabaseSignup(cfg *config.Supabase, email string, password string) (models.SignupResponse, error) {
	if err := validatePasswordStrength(password); err != nil {
		return models.SignupResponse{}, err
	}

	supabaseURL := cfg.URL
	apiKey := cfg.ServiceRoleKey

	payload := models.SignUpPayload{
		Email:    email,
		Password: password,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return models.SignupResponse{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/v1/signup", supabaseURL), bytes.NewBuffer(payloadBytes))
	if err != nil {
		slog.Error("Error in Request Creation: ", "err", err)
		return models.SignupResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("apikey", apiKey)

	res, err := Client.Do(req)
	if err != nil {
		slog.Error("Error executing request: ", "err", err)
		return models.SignupResponse{}, err
	}
	err = res.Body.Close()
	if err != nil {
		slog.Error("Error closing response body: ", "err", err)
		return models.SignupResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Error reading response body: ", "body", body)
		return models.SignupResponse{}, err
	}

	if res.StatusCode != http.StatusOK {
		supabaseError := &SupabaseError{}
		if err := json.Unmarshal(body, supabaseError); err != nil {
			slog.Error("Error parsing response: ", "err", err)
			return models.SignupResponse{}, err
		}
		slog.Error("Error Response: ", "res.StatusCode", res.StatusCode, "body", string(body))
		return models.SignupResponse{}, errs.NewHTTPError(res.StatusCode, supabaseError)
	}

	var response models.SignupResponse
	if err := json.Unmarshal(body, &response); err != nil {
		slog.Error("Error parsing response: ", "err", err)
		return models.SignupResponse{}, err
	}

	return response, nil
}

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
	err = res.Body.Close()
	if err != nil {
		slog.Error("Error closing response body: ", "err", err)
		return err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(res.Body)
		slog.Error("Error deleting user from Supabase: ", "res.StatusCode", res.StatusCode, "body", string(body))
		return err
	}

	return nil
}
