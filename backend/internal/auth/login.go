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



func SupabaseLogin(cfg *config.Supabase, email string, password string) (models.LoginResponse, error) {
	supabaseURL := cfg.URL
	serviceroleKey := cfg.ServiceRoleKey

	payload := models.Payload{
		Email:    email,
		Password: password,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return models.LoginResponse{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/v1/token?grant_type=password", supabaseURL), bytes.NewBuffer(payloadBytes))
	if err != nil {
		return models.LoginResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", serviceroleKey))
	req.Header.Set("apikey", serviceroleKey)
	

	res, err := Client.Do(req)
	if err != nil {
		slog.Error("Failed to execute Request", "err", err)
		return models.LoginResponse{}, errs.BadRequest("Failed to execute Request")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Failed to read response body", "err", err)
		return models.LoginResponse{}, errs.BadRequest("Failed to read response body")
	}

	if res.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("Failed to login %d, %s", res.StatusCode, body)
		fmt.Print(errorMsg)
		return models.LoginResponse{}, errs.BadRequest(errorMsg)
	}

	var signInResponse models.LoginResponse
	err = json.Unmarshal(body, &signInResponse)
	if err != nil {
		slog.Error("Failed to parse response body", "body", err)
		return models.LoginResponse{}, errs.BadRequest("Failed to parse response body")
	}

	if signInResponse.Error != nil {
		return models.LoginResponse{}, errs.BadRequest(fmt.Sprintf("Sign In Response Error %v", signInResponse.Error))
	}

	return signInResponse, nil
}