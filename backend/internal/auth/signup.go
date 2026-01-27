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

type SignUpPayload struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserSignupResponse struct {
	ID uuid.UUID `json:"id"`
}

type SignupResponse struct {
	AccessToken string             `json:"access_token"`
	User        UserSignupResponse `json:"user"`
}

func SupabaseSignup(cfg *config.Supabase, email string, password string) (SignupResponse, error) {
	if err := validatePasswordStrength(password); err != nil {
		return SignupResponse{}, errs.BadRequest(fmt.Sprintf("Weak Password: %v", err))
	}

	supabaseURL := cfg.URL
	apiKey := cfg.ServiceRoleKey

	payload := SignUpPayload{
		Email:    email,
		Password: password,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return SignupResponse{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/v1/signup", supabaseURL), bytes.NewBuffer(payloadBytes))
	if err != nil {
		slog.Error("Error in Request Creation: ", "err", err)
		return SignupResponse{}, errs.BadRequest(fmt.Sprintf("Failed to create request: %v", err))
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("apikey", apiKey)
	
	res, err := Client.Do(req)
	if err != nil {
		slog.Error("Error executing request: ", "err", err)
		return SignupResponse{}, errs.BadRequest(fmt.Sprintf("Failed to execute request: %v, %s", err, supabaseURL))
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Error reading response body: ", "body", body)
		return SignupResponse{}, errs.BadRequest(fmt.Sprintf("Failed to read response body: %s", body))
	}

	if res.StatusCode != http.StatusOK {
		slog.Error("Error Response: ", "res.StatusCode", res.StatusCode, "body", string(body))
		return SignupResponse{}, errs.BadRequest(fmt.Sprintf("Failed to login %d, %s", res.StatusCode, body))
	}

	var response SignupResponse
	fmt.Print(response)
	if err := json.Unmarshal(body, &response); err != nil {
		slog.Error("Error parsing response: ", "err", err)
		return SignupResponse{}, errs.BadRequest("Failed to parse request")
	}

	return response, nil
}
