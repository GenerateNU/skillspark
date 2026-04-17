package models

import (
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

type OrgLink struct {
	Href  string `json:"href" doc:"URL to the organization resource"`
	Label string `json:"label" doc:"Human-readable label for the organization link"`
}

type Organization struct {
	ID                     uuid.UUID         `json:"id" db:"id"`
	Name                   string            `json:"name" db:"name"`
	About                  *string           `json:"about,omitempty" db:"about"`
	Active                 bool              `json:"active" db:"active"`
	Links                  []OrgLink         `json:"links" db:"links"`
	PfpS3Key               *string           `json:"pfp_s3_key,omitempty" db:"pfp_s3_key"`
	PresignedURL           *string           `json:"presigned_url,omitempty"`
	LocationID             *uuid.UUID        `json:"location_id,omitempty" db:"location_id"`
	Latitude               *float64          `json:"latitude,omitempty"`
	Longitude              *float64          `json:"longitude,omitempty"`
	District               *string           `json:"district,omitempty"`
	CreatedAt              time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt              time.Time         `json:"updated_at" db:"updated_at"`
	StripeAccountID        *string           `json:"stripe_account_id,omitempty" db:"stripe_account_id"`
	StripeAccountActivated bool              `json:"stripe_account_activated" db:"stripe_account_activated" default:"false"`
	ReviewSummary          *OrgReviewSummary `json:"review_summary,omitempty"`
}

// CreateOrganizationRouteInput is the multipart form input for creating an organization with an image
type CreateOrganizationRouteInput struct {
	AcceptLanguage string `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	RawBody        huma.MultipartFormFiles[CreateOrganizationFormData]
}

// CreateOrganizationFormData holds the parsed form data for creating an organization
type UpdateOrganizationFormData struct {
	Name         string        `form:"name" required:"true" minLength:"1" maxLength:"255"`
	About        string        `form:"about"`
	Active       bool          `form:"active"`
	LocationID   uuid.UUID     `form:"location_id"`
	ProfileImage huma.FormFile `form:"profile_image" contentType:"image/png,image/jpeg"`
	Links        string        `form:"links"`
}

type CreateOrganizationFormData struct {
	Name         string        `form:"name" required:"true" minLength:"1" maxLength:"255"`
	About        string        `form:"about"`
	Active       bool          `form:"active"`
	LocationID   uuid.UUID     `form:"location_id"`
	Links        string        `form:"links"`
	ProfileImage huma.FormFile `form:"profile_image" contentType:"image/png,image/jpeg"`
}

type UpdateOrganizationRouteInput struct {
	AcceptLanguage string    `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	ID             uuid.UUID `path:"id"`
	RawBody        huma.MultipartFormFiles[UpdateOrganizationFormData]
}

type CreateOrganizationBody struct {
	Name       string     `json:"name" minLength:"1" maxLength:"255" doc:"Organization name"`
	About      *string    `json:"about,omitempty" doc:"Short description of the organization"`
	Active     *bool      `json:"active,omitempty" doc:"Active status (defaults to true)"`
	LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
	Links      []OrgLink  `json:"links,omitempty" doc:"List of links associated with the organization"`
}

type UpdateOrganizationBody struct {
	Name       *string    `json:"name" minLength:"1" maxLength:"255" doc:"Organization name"`
	About      *string    `json:"about,omitempty" doc:"Short description of the organization"`
	Active     *bool      `json:"active,omitempty" doc:"Active status (defaults to true)"`
	LocationID *uuid.UUID `json:"location_id,omitempty" format:"uuid" doc:"Associated location ID"`
	Links      *[]OrgLink `json:"links,omitempty" doc:"List of links associated with the organization"`
}

type CreateOrganizationInput struct {
	AcceptLanguage string `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	Body           CreateOrganizationBody
}

type CreateOrganizationOutput struct {
	Body Organization
}

type UpdateOrganizationInput struct {
	AcceptLanguage string    `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	ID             uuid.UUID `path:"id" format:"uuid" doc:"Organization ID"`
	Body           UpdateOrganizationBody
}

type UpdateOrganizationOutput struct {
	Body Organization `json:"body"`
}

type GetOrganizationByIDInput struct {
	AcceptLanguage string    `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	ID             uuid.UUID `path:"id" format:"uuid" doc:"Organization ID"`
}

type GetOrganizationByIDOutput struct {
	Body Organization `json:"body"`
}

type GetAllOrganizationsInput struct {
	AcceptLanguage string `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	Page           int    `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize       int    `query:"page_size" minimum:"1" maximum:"100" default:"10" doc:"Number of items per page"`
}

type GetAllOrganizationsOutput struct {
	Body []Organization `json:"body" doc:"List of organizations"`
}

type DeleteOrganizationInput struct {
	AcceptLanguage string    `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	ID             uuid.UUID `path:"id" format:"uuid" doc:"Organization ID"`
}

type DeleteOrganizationOutput struct {
	Body Organization `json:"body" doc:"The deleted organization"`
}

type GetEventOccurrencesByOrganizationIDInput struct {
	AcceptLanguage string    `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	ID             uuid.UUID `path:"organization_id" doc:"ID of an organization"`
}

type GetEventOccurrencesByOrganizationIDOutput struct {
	Body []EventOccurrence `json:"body" doc:"List of event occurrences in the database that match the organization ID"`
}

// CreateOrgDBBody / UpdateOrgDBBody hold the bilingual columns written to the DB.
type CreateOrgDBBody struct {
	Name       string
	AboutEN    *string
	AboutTH    *string
	Active     *bool
	LocationID *uuid.UUID
	Links      []OrgLink
}

type UpdateOrgDBBody struct {
	Name       *string
	AboutEN    *string
	AboutTH    *string
	Active     *bool
	LocationID *uuid.UUID
	Links      *[]OrgLink
}

type CreateOrganizationDBInput struct {
	AcceptLanguage string
	Body           CreateOrgDBBody
}

type UpdateOrganizationDBInput struct {
	AcceptLanguage string
	ID             uuid.UUID
	Body           UpdateOrgDBBody
}
