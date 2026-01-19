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

type GetManagerByIDInput struct {
	ID uuid.UUID `path:"id"`
}
type GetManagerByIDOutput struct {
	Body Manager `json:"body"`
}

type GetManagerByOrgIDInput struct {
	OrganizationID uuid.UUID `path:"organization_id"`
}

type GetManagerByOrgIDOutput struct {
	Body Manager `json:"body"`
}
