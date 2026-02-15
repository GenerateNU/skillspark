package models

import (
	"time"

	"github.com/google/uuid"
)

type Guardian struct {
	ID                  uuid.UUID `json:"id" db:"id"`
	UserID              uuid.UUID `json:"user_id" db:"user_id"`
	Name                string    `json:"name" db:"name"`
	Email               string    `json:"email" db:"email"`
	Username            string    `json:"username" db:"username"`
	ProfilePictureS3Key *string   `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
	LanguagePreference  string    `json:"language_preference" db:"language_preference"`
	AuthID              uuid.UUID `json:"auth_id" db:"auth_id"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}

type GetGuardianByIDInput struct {
	ID uuid.UUID `path:"id"`
}

type DeleteGuardianInput struct {
	ID uuid.UUID `path:"id"`
}

type DeleteGuardianOutput struct {
	Body *Guardian `json:"body"`
}

type GetGuardianByChildIDInput struct {
	ChildID uuid.UUID `path:"child_id"`
}

type CreateGuardianInput struct {
	Body struct {
		Name                string  `json:"name" doc:"Name of the guardian"`
		Email               string  `json:"email" doc:"Email of the guardian"`
		Username            string  `json:"username" doc:"Username of the guardian"`
		ProfilePictureS3Key *string `json:"profile_picture_s3_key,omitempty" doc:"S3 key for profile picture" required:"false"`
		LanguagePreference  string  `json:"language_preference" doc:"Language preference"`
		AuthID              uuid.UUID `json:"auth_id" db:"auth_id" doc:"auth id of the guardian being created"`
	}
}

type UpdateGuardianInput struct {
	ID   uuid.UUID `path:"id"`
	Body struct {
		Name                string  `json:"name" doc:"Name of the guardian"`
		Email               string  `json:"email" doc:"Email of the guardian"`
		Username            string  `json:"username" doc:"Username of the guardian"`
		ProfilePictureS3Key *string `json:"profile_picture_s3_key,omitempty" doc:"S3 key for profile picture"`
		LanguagePreference  string  `json:"language_preference" doc:"Language preference"`
	}
}

type UpdateGuardianOutput struct {
	Body *Guardian `json:"body"`
}

type CreateGuardianOutput struct {
	Body *Guardian `json:"body"`
}

type GetGuardianByChildIDOutput struct {
	Body *Guardian `json:"body"`
}

type GetGuardianByIDOutput struct {
	Body *Guardian `json:"body"`
}
