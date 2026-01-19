package models

import (
	"time"

	"github.com/google/uuid"
)

type Guardian struct {
	ID uuid.UUID `json:"id" db:"id"`
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type GetGuardianByIDInput struct {
	ID uuid.UUID `path:"id"`
}

type GetGuardianByChildIDInput struct {
	ChildID uuid.UUID `path:"child_id"`
}

type CreateGuardianInput struct {
	Body struct {
		UserID uuid.UUID `json:"user_id" doc:"The associated user ID of the guardian"`
	}
}

type UpdateGuardianInput struct {
	ID uuid.UUID `path:"id"`
	Body struct {
		UserID uuid.UUID `json:"user_id" doc:"The associated user ID of the guardian"`
	}
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