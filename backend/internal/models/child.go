package models

import (
	"time"

	"github.com/google/uuid"
)

type Child struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	AgeYears   int       `json:"age_years" db:"age_years"`
	CustomerID uuid.UUID `json:"customer_id" db:"customer_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}