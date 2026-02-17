package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                  uuid.UUID  `json:"id" db:"id"`
	Name                string     `json:"name" db:"name"`
	Email               string     `json:"email" db:"email"`
	Username            string     `json:"username" db:"username"`
	ProfilePictureS3Key *string    `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
	LanguagePreference  string     `json:"language_preference" db:"language_preference"`
	AuthID              uuid.UUID  `json:"auth_id" db:"auth_id"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at"`
}

// Create
type CreateUserInput struct {
	Body struct {
		Name                string     `json:"name" db:"name" doc:"name of the user"`
		Email               string     `json:"email" db:"email" doc:"email of the user"`
		Username            string     `json:"username" db:"username" doc:"username of the user"`
		ProfilePictureS3Key *string    `json:"profile_picture_s3_key,omitempty" db:"profile_picture_s3_key" doc:"s3 key of the user's profile picture"`
		LanguagePreference  string     `json:"language_preference" db:"language_preference" doc:"language preference of the user"`
		AuthID              uuid.UUID  `json:"auth_id" db:"auth_id" doc:"supabase auth id of the user"`
	}
}

type CreateUserOutput struct {
	Body *User `json:"body"`
}

// Update
type UpdateUserInput struct {
	ID   uuid.UUID `path:"id"`
	Body struct {
		Name                *string    `json:"name,omitempty" db:"name" doc:"name of the user"`
		Email               *string    `json:"email,omitempty" db:"email" doc:"email of the user"`
		Username            *string    `json:"username,omitempty" db:"username" doc:"username of the user"`
		ProfilePictureS3Key *string    `json:"profile_picture_s3_key,omitempty" db:"profile_picture_s3_key" doc:"s3 key of the user's profile picture"`
		LanguagePreference  *string    `json:"language_preference,omitempty" db:"language_preference" doc:"language preference of the user"`
		AuthID              *uuid.UUID `json:"auth_id,omitempty" db:"auth_id" doc:"supabase auth id of the user"`
	}
}

type UpdateUserOutput struct {
	Body *User `json:"body"`
}

// Get By ID
type GetUserByIDInput struct {
	ID uuid.UUID `path:"id"`
}

type GetUserByIDOutput struct {
	Body *User `json:"body"`
}

// Delete
type DeleteUserInput struct {
	ID uuid.UUID `path:"id"`
}

type DeleteUserOutput struct {
	Body *User `json:"body"`
}
