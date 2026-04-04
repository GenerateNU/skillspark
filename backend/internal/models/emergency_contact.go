package models

import (
	"time"

	"github.com/google/uuid"
)

type EmergencyContact struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	GuardianID  uuid.UUID `json:"guardian_id" db:"guardian_id"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type GetEmergencyContactByGuardianIDInput struct {
	GuardianID uuid.UUID `path:"guardian_id" required:"true"`
}

type GetEmergencyContactByGuardianIDOutput struct {
	Body []*EmergencyContact `json:"body" doc:"List of all emergency contacts related to the given guardian in the database"`
}

type CreateEmergencyContactInput struct {
	Body struct {
		Name        string    `json:"name" db:"name"`
		GuardianID  uuid.UUID `json:"guardian_id" db:"guardian_id"`
		PhoneNumber string    `json:"phone_number" db:"phone_number"`
	} `json:"body" doc:"New emergency contact to add"`
}

type CreateEmergencyContactOutput struct {
	Body *EmergencyContact `json:"body" doc:"Created Emergency Contact"`
}

type UpdateEmergencyContactInput struct {
	ID   uuid.UUID `path:"id" format:"uuid" doc:"Emergency Contact ID to update" required:"true"`
	Body struct {
		GuardianID  uuid.UUID `json:"guardian_id" db:"guardian_id"`
		PhoneNumber string    `json:"phone_number" db:"phone_number"`
		Name        string    `json:"name" db:"name"`
	} `json:"body" doc:"Emergency contact to update"`
}

type UpdateEmergencyContactOutput struct {
	Body *EmergencyContact `json:"body" doc:"Updated Emergency Contact"`
}

type DeleteEmergencyContactInput struct {
	ID uuid.UUID `path:"id"`
}

type DeleteEmergencyContactBody struct {
	SuccessMessage string `json:"success_message"`
}

type DeleteEmergencyContactOutput struct {
	Body *DeleteEmergencyContactBody
}
