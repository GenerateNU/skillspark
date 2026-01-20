package models

import (
	"time"

	"github.com/google/uuid"
)

type Manager struct {
	ID             uuid.UUID `json:"id" db:"id"`
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	OrganizationID uuid.UUID `json:"organization_id" db:"organization_id"`
	Role           string    `json:"role" db:"role"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
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
		UserID         uuid.UUID `json:"user_id" db:"user_id" doc:"user id of the manager"`
		OrganizationID uuid.UUID `json:"organization_id" db:"organization_id" doc:"organization id of the organization the manager is associated with"`
		Role           string    `json:"role" db:"role" doc:"role of the manager being created"`
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
		ID             uuid.UUID `json:"id" db:"id" doc:"id of the manager"`
		UserID         uuid.UUID `json:"user_id" db:"user_id" doc:"user id of the manager"`
		OrganizationID uuid.UUID `json:"organization_id" db:"organization_id" doc:"organization id of the organization the manager is associated with"`
		Role           string    `json:"role" db:"role" doc:"role of the manager being created"`
	}
}

type PatchManagerOutput struct {
	Body *Manager `json:"body"`
}
