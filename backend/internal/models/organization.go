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
	Page     int    `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize int    `query:"page_size" minimum:"1" maximum:"100" default:"20" doc:"Number of items per page"`
}

type GetAllOrganizationsOutput struct {
	Body struct {
		Organizations []Organization `json:"organizations" doc:"List of organizations"`
		Page          int            `json:"page" doc:"Current page number"`
		PageSize      int            `json:"page_size" doc:"Items per page"`
		TotalCount    int            `json:"total_count" doc:"Total number of organizations"`
		TotalPages    int            `json:"total_pages" doc:"Total number of pages"`
	}
}

type DeleteOrganizationInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Organization ID"`
}

type DeleteOrganizationOutput struct {
	Body struct {
		Message string `json:"message" doc:"Confirmation message"`
		ID      string `json:"id" doc:"Deleted organization ID"`
	}
}