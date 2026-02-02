package models

import (
	"time"

	"github.com/danielgtaylor/huma/v2"
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

// CreateOrganizationRouteInput is the multipart form input for creating an organization with an image
type CreateOrganizationRouteInput struct {
	RawBody huma.MultipartFormFiles[CreateOrganizationFormData]
}

// CreateOrganizationFormData holds the parsed form data for creating an organization
type UpdateOrganizationFormData struct {
	Name         string        `form:"name" required:"true" minLength:"1" maxLength:"255"`
	Active       bool          `form:"active"`
	LocationID   uuid.UUID     `form:"location_id"`
	ProfileImage huma.FormFile `form:"profile_image" contentType:"image/png,image/jpeg"`
}

type CreateOrganizationFormData struct {
	Name         string        `form:"name" required:"true" minLength:"1" maxLength:"255"`
	Active       bool          `form:"active"`
	LocationID   uuid.UUID     `form:"location_id"`
	ProfileImage huma.FormFile `form:"profile_image" contentType:"image/png,image/jpeg"`
}

type UpdateOrganizationRouteInput struct {
	ID      uuid.UUID `path:"id"`
	RawBody huma.MultipartFormFiles[UpdateOrganizationFormData]
}

type CreateOrganizationBody struct {
	Name       string     `json:"name" minLength:"1" maxLength:"255" doc:"Organization name"`
	Active     *bool      `json:"active,omitempty" doc:"Active status (defaults to true)"`
	LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
}

type UpdateOrganizationBody struct {
	Name       *string    `json:"name" minLength:"1" maxLength:"255" doc:"Organization name"`
	Active     *bool      `json:"active,omitempty" doc:"Active status (defaults to true)"`
	LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
}

type CreateOrganizationInput struct {
	Body CreateOrganizationBody
}

type CreateOrganizationOutput struct {
	Body         Organization
	PresignedURL *string `json:"presigned_url"`
}

type UpdateOrganizationInput struct {
	ID   uuid.UUID `path:"id" format:"uuid" doc:"Organization ID"`
	Body UpdateOrganizationBody
}

type UpdateOrganizationOutput struct {
	Body         Organization `json:"body"`
	PresignedURL *string      `json:"presigned_url"`
}

type GetOrganizationByIDInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Organization ID"`
}

type GetOrganizationByIDOutput struct {
	Body         Organization `json:"body"`
	PresignedURL *string      `json:"presigned_url"`
}

type GetAllOrganizationsInput struct {
	Page     int `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize int `query:"page_size" minimum:"1" maximum:"100" default:"10" doc:"Number of items per page"`
}

type GetAllOrganizationsOutput struct {
	Body          []Organization `json:"body" doc:"List of organizations"`
	PresignedURLS []string       `json:"presigned_urls" doc:"List of Presigned Urls"`
}

type DeleteOrganizationInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Organization ID"`
}

type DeleteOrganizationOutput struct {
	Body Organization `json:"body" doc:"The deleted organization"`
}
