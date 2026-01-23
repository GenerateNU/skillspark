package models

import (
	"time"

	"github.com/google/uuid"
)


type Organization struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	Active     bool       `json:"active" db:"active"`
	PfpS3Key   *string    `json:"pfp_s3_key,omitempty" db:"pfp_s3_key"`
	LocationID *uuid.UUID `json:"location_id,omitempty" db:"location_id"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
}


type CreateOrganizationInput struct {
	Body struct {
		Name       string     `json:"name" minLength:"1" maxLength:"255" doc:"Organization name"`
		Active     *bool      `json:"active,omitempty" doc:"Active status (defaults to true)"`
		PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
		LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
	}
}

type CreateOrganizationOutput struct {
	Body Organization
}


type UpdateOrganizationInput struct {
	ID   uuid.UUID `path:"id" format:"uuid" doc:"Organization ID"`
	Body struct {
		Name       *string    `json:"name,omitempty" minLength:"1" maxLength:"255" doc:"Organization name"`
		Active     *bool      `json:"active,omitempty" doc:"Active status"`
		PfpS3Key   *string    `json:"pfp_s3_key,omitempty" maxLength:"500" doc:"S3 key for profile picture"`
		LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
	}
}

type UpdateOrganizationOutput struct {
	Body Organization
}


type GetOrganizationByIDInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Organization ID"`
}

type GetOrganizationByIDOutput struct {
	Body Organization
}


type GetAllOrganizationsInput struct {
	Page     int `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize int `query:"page_size" minimum:"1" maximum:"100" default:"10" doc:"Number of items per page"`
}

type GetAllOrganizationsOutput struct {
	Body [] Organization `json:"body" doc:"List of organizations"`
}

type DeleteOrganizationInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Organization ID"`
}

type DeleteOrganizationOutput struct {
	Body Organization `json:"body" doc:"The deleted organization"`
}