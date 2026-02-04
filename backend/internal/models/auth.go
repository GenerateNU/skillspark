package models

import (
	"net/http"

	"github.com/google/uuid"
)

type GuardianSignUpInput struct {
	Body struct {
		Name                string  `json:"name" db:"name"`
		Email               string  `json:"email" db:"email"`
		Username            string  `json:"username" db:"username"`
		Password            string  `json:"password" db:"password"`
		ProfilePictureS3Key *string `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
		LanguagePreference  string  `json:"language_preference" db:"language_preference"`
	}
}

type GuardianSignUpOutput struct {
	Body struct {
		Token      string    `json:"token" db:"token"`
		GuardianID uuid.UUID `json:"guardian_id" db:"guardian_id"`
	} `json:"body"`
}

type ManagerSignUpInput struct {
	Body struct {
		Name                string     `json:"name" db:"name" doc:"name of the manager" required:"true"`
		Email               string     `json:"email" db:"email" doc:"email of the manager" required:"true"`
		Username            string     `json:"username" db:"username" doc:"username of the manager" required:"true"`
		Password            string     `json:"password" db:"password" doc:"password of the manager" required:"true"`
		ProfilePictureS3Key *string    `json:"profile_picture_s3_key,omitempty" db:"profile_picture_s3_key" doc:"profile picture s3 key of the manager" required:"false"`
		LanguagePreference  string     `json:"language_preference" db:"language_preference" doc:"language preference of the manager" required:"true"`
		OrganizationID      uuid.UUID `json:"organization_id" db:"organization_id" doc:"organization id of the organization the manager is associated with" required:"true"`
		Role                string    `json:"role" db:"role" doc:"role of the manager being created" required:"true"`
		AuthID              *uuid.UUID `json:"auth_id,omitempty" db:"auth_id" doc:"auth id of the manager being created" required:"false"`
	}
}

type ManagerSignUpOutput struct {
	Body struct {
		Token     string    `json:"token" db:"token"`
		ManagerID uuid.UUID `json:"manager_id" db:"manager_id"`
	}
}

type LoginInput struct {
	Body struct {
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}
}

type UserResponse struct {
	ID uuid.UUID `json:"id"`
}

type GuardianLoginOutput struct {
	AccessTokenCookie http.Cookie `header:"Set-Cookie"`
	Body              struct {
		GuardianID uuid.UUID `json:"guardian_id" db:"guardian_id"`
	} `json:"body"`
}

type ManagerLoginOutput struct {
	AccessTokenCookie http.Cookie `header:"Set-Cookie"`
	Body              struct {
		ManagerID uuid.UUID `json:"manager_id" db:"manager_id"`
	} `json:"body"`
}

// Login response from Supabase API
type LoginResponse struct {
	AccessToken  string       `json:"access_token"`
	TokenType    string       `json:"token_type"`
	ExpiresIn    int          `json:"expires_in"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
	Error        interface{}  `json:"error"`
}

type Payload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpPayload represents the payload for Supabase signup
type SignUpPayload struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

// UserSignupResponse represents the user data returned from Supabase signup
type UserSignupResponse struct {
	ID uuid.UUID `json:"id"`
}

// SignupResponse represents the complete response from Supabase signup
type SignupResponse struct {
	AccessToken string             `json:"access_token"`
	User        UserSignupResponse `json:"user"`
}

// Error response from Supabase API
type SupabaseError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Message   string `json:"msg"`
}

func (e *SupabaseError) Error() string {
	return e.Message
}
