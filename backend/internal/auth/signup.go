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

// type GuardianSignUpRequest struct {
// 	Name                string `json:"name" db:"name"`
// 	Email               string `json:"email" db:"email"`
// 	Username            string `json:"username" db:"username"`
// 	Password            string `json:"password" db:"password"`
// 	ProfilePictureS3Key string `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
// 	LanguagePreference  string `json:"language_preference" db:"language_preference"`
// }

// type ManagerSignUpRequest struct {
// 	Name                string `json:"name" db:"name"`
// 	Email               string `json:"email" db:"email"`
// 	Username            string `json:"username" db:"username"`
// 	Password            string `json:"password" db:"password"`
// 	ProfilePictureS3Key string `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
// 	OrganizationID      string `json:"organization_id" db:"organization_id"`
// 	Role                string `json:"role" db:"role"`
// }

// type GuardianSignUpResponse struct {
// 	ID                  string `json:"id" db:"id"`
// 	Name                string `json:"name" db:"name"`
// 	Email               string `json:"email" db:"email"`
// 	Username            string `json:"username" db:"username"`
// 	ProfilePictureS3Key string `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
// 	LanguagePreference  string `json:"language_preference" db:"language_preference"`
// }

// type ManagerSignUpResponse struct {
// 	ID                  string `json:"id" db:"id"`
// 	Name                string `json:"name" db:"name"`
// 	Email               string `json:"email" db:"email"`
// 	Username            string `json:"username" db:"username"`
// 	ProfilePictureS3Key string `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
// 	OrganizationID      string `json:"organization_id" db:"organization_id"`
// 	Role                string `json:"role" db:"role"`
// }

func validatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("Password must be atleast 8 characters long")
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#~$%^&*()+|_.,;<>?/{}\-]`).MatchString(password)

	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return errors.New("Password must include uppercase, lowercase, digit and special characters")
	}

	return nil
}

// type GuardianSignUpPayload struct {
// 	Email               string `json:"email" db:"email"`
// 	Password            string `json:"password" db:"password"`
// 	Name                string `json:"name" db:"name"`
// 	Username            string `json:"username" db:"username"`
// 	ProfilePictureS3Key string `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
// 	LanguagePreference  string `json:"language_preference" db:"language_preference"`
// }

func SupabaseSignup(cfg *config.Supabase, email string, password string) (models.SignupResponse, error) {
	if err := validatePasswordStrength(password); err != nil {
		return models.SignupResponse{}, errs.BadRequest(fmt.Sprintf("Weak Password: %v", err))
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
		return models.SignupResponse{}, errs.BadRequest(fmt.Sprintf("Failed to create request: %v", err))
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("apikey", apiKey)

	res, err := Client.Do(req)
	if err != nil {
		slog.Error("Error executing request: ", "err", err)
		return models.SignupResponse{}, errs.BadRequest(fmt.Sprintf("Failed to execute request: %v, %s", err, supabaseURL))
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Error reading response body: ", "body", body)
		return models.SignupResponse{}, errs.BadRequest(fmt.Sprintf("Failed to read response body: %s", body))
	}

	if res.StatusCode != http.StatusOK {
		slog.Error("Error Response: ", "res.StatusCode", res.StatusCode, "body", string(body))
		return models.SignupResponse{}, errs.BadRequest(fmt.Sprintf("Failed to login %d, %s", res.StatusCode, body))
	}

	var response models.SignupResponse
	fmt.Print(response)
	if err := json.Unmarshal(body, &response); err != nil {
		slog.Error("Error parsing response: ", "err", err)
		return models.SignupResponse{}, errs.BadRequest("Failed to parse request")
	}

	return response, nil
}

// SupabaseDeleteUser deletes a user from Supabase auth by their user ID
func SupabaseDeleteUser(cfg *config.Supabase, userID uuid.UUID) error {
	supabaseURL := cfg.URL
	apiKey := cfg.ServiceRoleKey

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/auth/v1/admin/users/%s", supabaseURL, userID.String()), nil)
	if err != nil {
		slog.Error("Error in Request Creation: ", "err", err)
		return errs.InternalServerError(fmt.Sprintf("Failed to create delete request: %v", err))
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("apikey", apiKey)

	res, err := Client.Do(req)
	if err != nil {
		slog.Error("Error executing delete request: ", "err", err)
		return errs.InternalServerError(fmt.Sprintf("Failed to execute delete request: %v", err))
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(res.Body)
		slog.Error("Error deleting user from Supabase: ", "res.StatusCode", res.StatusCode, "body", string(body))
		return errs.InternalServerError(fmt.Sprintf("Failed to delete user from Supabase: %d, %s", res.StatusCode, body))
	}

	return nil
}
