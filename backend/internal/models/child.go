package models

import (
	"time"

	"github.com/google/uuid"
)

type Interest string

const (
	InterestFootball      Interest = "football"
	InterestSoccer        Interest = "soccer"
	InterestDragonDancing Interest = "dragon_dancing"
	InterestBedRotting    Interest = "bed_rotting"
	InterestClimbing      Interest = "climbing"
	InterestWalking       Interest = "walking"
	InterestBunnySpotting Interest = "bunny_spotting"
)

type CreateChildInput struct {
	Body struct {
		Name       string     `json:"name" db:"name" doc:"Name of the child" minLength:"1" maxLength:"200"`
		SchoolID   uuid.UUID  `json:"school_id" db:"school_id" doc:"ID of the school the child goes to"`
		BirthMonth int        `json:"birth_month" db:"birth_month" doc:"Birth month of the child" minimum:"1" maximum:"12"`
		BirthYear  int        `json:"birth_year" db:"birth_year" doc:"Birth year of the child" minimum:"2000" maximum:"2026"` // NOTE: uses current year, but could also be in the future
		Interests  []Interest `json:"interests" db:"interests" doc:"Interests of the child"`
		GuardianID uuid.UUID  `json:"guardian_id" db:"guardian_id" doc:"ID of the child's guardian"`
	}
}

type UpdateChildInput struct {
	Body struct {
		ID         uuid.UUID   `json:"id" path:"id"`
		Name       *string     `json:"name" db:"name" doc:"Name of the child" minLength:"1" maxLength:"200"`
		SchoolID   *uuid.UUID  `json:"school_id" db:"school_id" doc:"ID of the school the child goes to"`
		BirthMonth *int        `json:"birth_month" db:"birth_month" doc:"Birth month of the child" minimum:"1" maximum:"12"`
		BirthYear  *int        `json:"birth_year" db:"birth_year" doc:"Birth year of the child" minimum:"2000" maximum:"2026"` // NOTE: uses current year, but could also be in the future
		Interests  *[]Interest `json:"interests" db:"interests" doc:"Interests of the child"`
		GuardianID *uuid.UUID  `json:"guardian_id" db:"guardian_id" doc:"ID of the child's guardian"`
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
	SchoolID uuid.UUID `json:"school_id" db:"school_id"`
	// ignore school name in db
	SchoolName string     `json:"school_name" db:"-"`
	BirthMonth int        `json:"birth_month" db:"birth_month"`
	BirthYear  int        `json:"birth_year" db:"birth_year"`
	Interests  []Interest `json:"interests" db:"interests"`
	GuardianID uuid.UUID  `json:"guardian_id" db:"guardian_id"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
}
