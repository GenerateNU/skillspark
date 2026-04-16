package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateChildInput struct {
	Body struct {
		Name             string    `json:"name" db:"name" doc:"Name of the child" minLength:"1" maxLength:"200"`
		SchoolID         uuid.UUID `json:"school_id" db:"school_id" doc:"ID of the school the child goes to"`
		BirthMonth       int       `json:"birth_month" db:"birth_month" doc:"Birth month of the child" minimum:"1" maximum:"12"`
		BirthYear        int       `json:"birth_year" db:"birth_year" doc:"Birth year of the child" minimum:"1950" maximum:"2100"` // NOTE: uses current year, but could also be in the future
		Interests        []string  `json:"interests" db:"interests" doc:"Interests of the child"`
		GuardianID       uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the child's guardian"`
		AvatarFace       *string   `json:"avatar_face,omitempty" db:"avatar_face" doc:"Avatar face identifier for the child's profile picture"`
		AvatarBackground *string   `json:"avatar_background,omitempty" db:"avatar_background" doc:"Background color hex for the child's avatar"`
	}
}

type UpdateChildInput struct {
	ID uuid.UUID `json:"-" path:"id"`

	Body struct {
		Name             *string    `json:"name,omitempty" db:"name" doc:"Name of the child" minLength:"1" maxLength:"200"`
		SchoolID         *uuid.UUID `json:"school_id,omitempty" db:"school_id" doc:"ID of the school the child goes to"`
		BirthMonth       *int       `json:"birth_month,omitempty" db:"birth_month" doc:"Birth month of the child" minimum:"1" maximum:"12"`
		BirthYear        *int       `json:"birth_year,omitempty" db:"birth_year" doc:"Birth year of the child" minimum:"2000" maximum:"2026"`
		Interests        *[]string  `json:"interests,omitempty" db:"interests" doc:"Interests of the child"`
		GuardianID       *uuid.UUID `json:"guardian_id,omitempty" db:"guardian_id" doc:"ID of the child's guardian"`
		AvatarFace       *string    `json:"avatar_face,omitempty" db:"avatar_face" doc:"Avatar face identifier for the child's profile picture"`
		AvatarBackground *string    `json:"avatar_background,omitempty" db:"avatar_background" doc:"Background color hex for the child's avatar"`
	}
}

// unify the input type that queries a child by the ID
type ChildIDInput struct {
	ID uuid.UUID `path:"id"`
}

type GuardianIDInput struct {
	ID uuid.UUID `path:"id"`
}

// unify the input type that outputs a child object
type ChildOutput struct {
	Body *Child `json:"body"`
}

type ChildrenOutput struct {
	Body []Child `json:"body"`
}

type Child struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name"`
	SchoolID uuid.UUID `json:"school_id" db:"-"`
	// ignore school name in db
	SchoolName       string    `json:"school_name" db:"school_name"`
	BirthMonth       int       `json:"birth_month" db:"birth_month"`
	BirthYear        int       `json:"birth_year" db:"birth_year"`
	Interests        []string  `json:"interests" db:"interests"`
	GuardianID       uuid.UUID `json:"guardian_id" db:"guardian_id"`
	AvatarFace       *string   `json:"avatar_face" db:"avatar_face"`
	AvatarBackground *string   `json:"avatar_background" db:"avatar_background"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}
