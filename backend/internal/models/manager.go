package models

import (
	"time"

	"github.com/google/uuid"
)

type Manager struct {
	ID                  uuid.UUID `json:"id" db:"id"`
	UserID              uuid.UUID `json:"user_id" db:"user_id"`
	OrganizationID      uuid.UUID `json:"organization_id" db:"organization_id"`
	Role                string    `json:"role" db:"role"`
	Name                string    `json:"name" db:"name"`
	Email               string    `json:"email" db:"email"`
	Username            string    `json:"username" db:"username"`
	ProfilePictureS3Key *string   `json:"profile_picture_s3_key" db:"profile_picture_s3_key"`
	LanguagePreference  string    `json:"language_preference" db:"language_preference"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}

// get by id
type GetManagerByIDInput struct {
	ID uuid.UUID `path:"id"`
}
type GetManagerByIDOutput struct {
	Body *Manager `json:"body"`
}

// get by org id
type GetManagerByOrgIDInput struct {
	OrganizationID uuid.UUID `path:"organization_id"`
}

type GetManagerByOrgIDOutput struct {
	Body *Manager `json:"body"`
}

// create
type CreateManagerInput struct {
	Body struct {
		Name                string     `json:"name" doc:"Name of the guardian"`
		Email               string     `json:"email" doc:"Email of the guardian"`
		Username            string     `json:"username" doc:"Username of the guardian"`
		ProfilePictureS3Key *string    `json:"profile_picture_s3_key,omitempty" doc:"S3 key for profile picture" required:"false"`
		LanguagePreference  string     `json:"language_preference" doc:"Language preference"`
		OrganizationID      *uuid.UUID `json:"organization_id,omitempty" db:"organization_id" doc:"organization id of the organization the manager is associated with"`
		Role                string     `json:"role" db:"role" doc:"role of the manager being created"`
		AuthID              *uuid.UUID `json:"auth_id,omitempty" db:"auth_id" doc:"auth id of the manager being created" required:"false"`
	}
}

type CreateManagerOutput struct {
	Body *Manager `json:"body"`
}

// delete

type DeleteManagerInput struct {
	ID uuid.UUID `path:"id"`
}

type DeleteManagerOutput struct {
	Body *Manager `json:"body"`
}

//patch/update

type PatchManagerInput struct {
	Body struct {
		ID                  uuid.UUID  `json:"id" db:"id" doc:"id of the manager"`
		Name                *string    `json:"name,omitempty" doc:"Name of the guardian"`
		Email               *string    `json:"email,omitempty" doc:"Email of the guardian"`
		Username            *string    `json:"username,omitempty" doc:"Username of the guardian"`
		ProfilePictureS3Key *string    `json:"profile_picture_s3_key,omitempty" doc:"S3 key for profile picture"`
		LanguagePreference  *string    `json:"language_preference,omitempty" doc:"Language preference"`
		OrganizationID      *uuid.UUID `json:"organization_id,omitempty" db:"organization_id" doc:"organization id"`
		Role                *string    `json:"role,omitempty" db:"role" doc:"role of the manager"`
	}
}

type PatchManagerOutput struct {
	Body *Manager `json:"body"`
}
